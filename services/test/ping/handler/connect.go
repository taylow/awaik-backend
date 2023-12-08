package handler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"connectrpc.com/connect"
	testv1 "github.com/taylow/awaik-backend/internal/gen/proto/test/v1"
	"github.com/taylow/awaik-backend/internal/gen/proto/test/v1/testv1connect"
)

// StartServer starts a test server for the PingService
func StartServer(handlerOpts []connect.HandlerOption, clientOpts []connect.ClientOption) testv1connect.PingServiceClient {
	svc := &pingServiceHandler{}
	mux := http.NewServeMux()
	mux.Handle(testv1connect.NewPingServiceHandler(svc, handlerOpts...))
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	pingClient := testv1connect.NewPingServiceClient(server.Client(), server.URL, clientOpts...)
	return pingClient
}

// NewConnectHandler returns a new ConnectHandler for the PingService
func NewConnectHandler() *pingServiceHandler {
	return &pingServiceHandler{}
}

// pingServiceHandler implements the PingService service
type pingServiceHandler struct {
	testv1connect.UnimplementedPingServiceHandler
}

// Ping implements testv1connect.PingServiceHandler
func (p *pingServiceHandler) Ping(
	ctx context.Context,
	req *connect.Request[testv1.PingRequest],
) (*connect.Response[testv1.PingResponse], error) {
	return connect.NewResponse(&testv1.PingResponse{
		Message: req.Msg.Message + " pong",
	}), nil

}

// PingStream implements testv1connect.PingServiceHandler
func (p *pingServiceHandler) PingStream(
	ctx context.Context,
	stream *connect.BidiStream[testv1.PingStreamRequest, testv1.PingStreamResponse],
) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		msg, err := stream.Receive()
		if err != nil && errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return fmt.Errorf("receive request: %w", err)
		}
		if err := stream.Send(&testv1.PingStreamResponse{
			Message: msg.Message + " pong",
		}); err != nil {
			return fmt.Errorf("send response: %w", err)
		}
	}
}
