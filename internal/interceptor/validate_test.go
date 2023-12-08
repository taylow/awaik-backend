package interceptor

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
	testv1 "github.com/taylow/awaik-backend/internal/gen/proto/test/v1"
	"github.com/taylow/awaik-backend/internal/gen/proto/test/v1/testv1connect"
	"github.com/taylow/awaik-backend/services/test/ping/handler"
)

func startServer() testv1connect.PingServiceClient {
	handlerOpts := []connect.HandlerOption{
		connect.WithInterceptors(NewValidateInterceptor()),
	}
	return handler.StartServer(handlerOpts, nil)
}

func TestValidateInterceptor_Unary(t *testing.T) {
	handlerOpts := []connect.HandlerOption{
		connect.WithInterceptors(NewValidateInterceptor()),
	}
	client := handler.StartServer(handlerOpts, nil)

	t.Run("valid unary request that has been validated", func(t *testing.T) {
		res, err := client.Ping(context.Background(), &connect.Request[testv1.PingRequest]{
			Msg: &testv1.PingRequest{
				Message: "ping",
			},
		})

		assert.NoError(t, err)
		assert.Equal(t, "ping pong", res.Msg.Message)
	})

	t.Run("invalid unary request that has been validated and fails", func(t *testing.T) {
		res, err := client.Ping(context.Background(), &connect.Request[testv1.PingRequest]{
			Msg: &testv1.PingRequest{
				Message: "",
			},
		})

		assert.Error(t, err)
		assert.Equal(t, err.(*connect.Error).Code(), connect.CodeInvalidArgument)
		assert.Nil(t, res)
	})
}

// TODO implement stream validation and test
// func TestValidateInterceptor_Stream(t *testing.T) {
// 	client := startServer()

// 	t.Run("valid stream request that has been validated", func(t *testing.T) {
// 		req := client.PingStream(context.Background())

// 		err := req.Send(&testv1.PingStreamRequest{
// 			Message: "ping",
// 		})

// 		res, err := req.Receive()
// 		assert.NoError(t, err)
// 		assert.Equal(t, "ping pong", res.Message)
// 	})

// 	t.Run("invalid stream request that has been validated and fails", func(t *testing.T) {
// 		req := client.PingStream(context.Background())

// 		err := req.Send(&testv1.PingStreamRequest{
// 			Message: "",
// 		})

// 		res, err := req.Receive()
// 		assert.Error(t, err)
// 		assert.Nil(t, res)
// 	})
// }
