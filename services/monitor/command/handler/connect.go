// Handler holds all handler implementations based on code generated from Protobufs
package handler

import (
	"context"

	"connectrpc.com/connect"
	monitorv1 "github.com/taylow/awaik-backend/internal/gen/proto/monitor/v1"
	"github.com/taylow/awaik-backend/internal/gen/proto/monitor/v1/monitorv1connect"
	"github.com/taylow/awaik-backend/services/monitor/command/domain"
)

var _ monitorv1connect.MonitorCommandServiceHandler = (*connectHandler)(nil)

// New creates Connect handlers
func NewConnectHandler(service domain.Service) *connectHandler {
	return &connectHandler{
		service: service,
	}
}

// connectHandler implements Connect handlers for each RPC endpoint
type connectHandler struct {
	monitorv1connect.UnimplementedMonitorCommandServiceHandler

	service domain.Service
}

// Create implements monitorv1connect.MonitorCommandServiceHandler.
func (h *connectHandler) Create(ctx context.Context, req *connect.Request[monitorv1.CreateRequest]) (*connect.Response[monitorv1.CreateResponse], error) {
	// err := req.Msg.ValidateAll()
	// if err != nil {
	// 	return nil, err
	// }

	monitor, err := h.service.Create(
		ctx,
		req.Msg.Monitor.ProjectId,
		req.Msg.Monitor.Name,
		req.Msg.Monitor.Description,
		req.Msg.Monitor.Interval,
		req.Msg.Monitor.Regions,
		protocolFromProto(req.Msg.Monitor.Protocol),
		protocolConfigFromProto(req.Msg.Monitor.ProtocolConfig),
	)
	if err != nil {
		return nil, err
	}

	return &connect.Response[monitorv1.CreateResponse]{
		Msg: &monitorv1.CreateResponse{
			Message: "Successfully created monitor",
			Monitor: monitorToProto(monitor),
		},
	}, nil
}

// Delete implements monitorv1connect.MonitorCommandServiceHandler.
func (h *connectHandler) Delete(ctx context.Context, req *connect.Request[monitorv1.DeleteRequest]) (*connect.Response[monitorv1.DeleteResponse], error) {
	err := h.service.Delete(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return &connect.Response[monitorv1.DeleteResponse]{
		Msg: &monitorv1.DeleteResponse{
			Message: "Successfully deleted monitor",
		},
	}, nil
}

// Pause implements monitorv1connect.MonitorCommandServiceHandler.
func (h *connectHandler) Pause(ctx context.Context, req *connect.Request[monitorv1.PauseRequest]) (*connect.Response[monitorv1.PauseResponse], error) {
	err := h.service.Pause(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return &connect.Response[monitorv1.PauseResponse]{
		Msg: &monitorv1.PauseResponse{
			Message: "Successfully paused monitor",
		},
	}, nil
}

// Resume implements monitorv1connect.MonitorCommandServiceHandler.
func (h *connectHandler) Resume(ctx context.Context, req *connect.Request[monitorv1.ResumeRequest]) (*connect.Response[monitorv1.ResumeResponse], error) {
	err := h.service.Resume(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return &connect.Response[monitorv1.ResumeResponse]{
		Msg: &monitorv1.ResumeResponse{
			Message: "Successfully resumed monitor",
		},
	}, nil
}

// Update implements monitorv1connect.MonitorCommandServiceHandler.
func (h *connectHandler) Update(ctx context.Context, req *connect.Request[monitorv1.UpdateRequest]) (*connect.Response[monitorv1.UpdateResponse], error) {
	monitor, err := h.service.Update(
		ctx,
		monitorFromProto(req.Msg.Monitor),
	)
	if err != nil {
		return nil, err
	}

	return &connect.Response[monitorv1.UpdateResponse]{
		Msg: &monitorv1.UpdateResponse{
			Message: "Successfully updated monitor",
			Monitor: monitorToProto(monitor),
		},
	}, nil
}
