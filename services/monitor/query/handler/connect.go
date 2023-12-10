// Handler holds all handler implementations based on code generated from Protobufs
package handler

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	monitorv1 "github.com/taylow/awaik-backend/internal/gen/proto/monitor/v1"
	"github.com/taylow/awaik-backend/internal/gen/proto/monitor/v1/monitorv1connect"
	"github.com/taylow/awaik-backend/services/monitor/transformer"
)

var _ monitorv1connect.MonitorQueryServiceHandler = (*connectHandler)(nil)

// New creates Connect handlers
func NewConnectHandler(service QueryService) *connectHandler {
	return &connectHandler{
		service: service,
	}
}

// connectHandler implements Connect handlers for each RPC endpoint
type connectHandler struct {
	monitorv1connect.UnimplementedMonitorQueryServiceHandler

	service QueryService
}

// Read implements monitorv1connect.MonitorQueryServiceHandler.
func (h *connectHandler) Read(ctx context.Context, req *connect.Request[monitorv1.ReadRequest]) (*connect.Response[monitorv1.ReadResponse], error) {
	monitor, err := h.service.Read(
		ctx,
		req.Msg.Id,
	)
	if err != nil {
		return nil, err
	}

	return &connect.Response[monitorv1.ReadResponse]{
		Msg: &monitorv1.ReadResponse{
			Message: "Successfully fetched monitor",
			Monitor: transformer.MonitorToProto(monitor),
		},
	}, nil
}

// ListByProject implements monitorv1connect.MonitorQueryServiceHandler.
func (h *connectHandler) ListByProject(ctx context.Context, req *connect.Request[monitorv1.ListByProjectRequest]) (*connect.Response[monitorv1.ListByProjectResponse], error) {
	monitors, err := h.service.ListByProject(
		ctx,
		req.Msg.ProjectId,
	)
	if err != nil {
		return nil, err
	}

	return &connect.Response[monitorv1.ListByProjectResponse]{
		Msg: &monitorv1.ListByProjectResponse{
			Message:  fmt.Sprintf("Found %d monitors for project", len(monitors)),
			Monitors: transformer.MonitorsToProto(monitors),
		},
	}, nil
}
