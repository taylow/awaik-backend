package query

import (
	"context"
	"net/http"
	"time"

	"connectrpc.com/connect"
	sharedconfig "github.com/taylow/awaik-backend/internal/config"
	"github.com/taylow/awaik-backend/internal/cors"
	"github.com/taylow/awaik-backend/internal/gen/proto/monitor/v1/monitorv1connect"
	"github.com/taylow/awaik-backend/internal/interceptor"
	"github.com/taylow/awaik-backend/services"
	"github.com/taylow/awaik-backend/services/monitor/query/config"
	"github.com/taylow/awaik-backend/services/monitor/query/handler"
	"github.com/taylow/awaik-backend/services/monitor/query/infra/monitor"
)

// ServiceName holds the unique name of the service
const ServiceName = "MonitorQueryService"
const Emoji = "👀 🔎"

// init inisialises the service and registers it with the service registry
func init() {
	// TODO load config from file/env
	cfg := &config.MonitorQueryServiceConfig{
		Service: &sharedconfig.Service{
			Name:  ServiceName,
			Emoji: Emoji,
		},
		Connect: &sharedconfig.Connect{
			Host: "",
			Port: 7781,
		},
		MonitorRepository: &sharedconfig.MonitorRepository{
			Type: "memory",
		},
	}

	services.Register(ServiceName, &MonitorQueryService{
		cfg: cfg,
	})
}

// MonitorQueryService represents the service that edits monitors
type MonitorQueryService struct {
	cfg *config.MonitorQueryServiceConfig

	service handler.QueryService
	server  *http.Server
}

// Name returns the name of the service
func (s *MonitorQueryService) Name() string {
	return ServiceName
}

// Emoji returns the emoji of the service
func (s *MonitorQueryService) Emoji() string {
	return Emoji
}

// Address returns the address of the service
func (s *MonitorQueryService) Address() string {
	return s.cfg.Connect.Address()
}

// Start starts the MonitorQueryService
func (s *MonitorQueryService) Start() error {
	monitorRepo, err := monitor.RepositoryFromConfig(s.cfg.MonitorRepository)
	if err != nil {
		return err
	}

	s.service = NewQueryService(monitorRepo)

	serviceHandler := handler.NewConnectHandler(s.service)

	opts := []connect.HandlerOption{
		connect.WithInterceptors(interceptor.WithTimeout(5 * time.Second)),
		connect.WithInterceptors(interceptor.NewLogInterceptor()),
		connect.WithInterceptors(interceptor.NewValidateInterceptor()),
	}

	mux := http.NewServeMux()
	mux.Handle(monitorv1connect.NewMonitorQueryServiceHandler(serviceHandler, opts...))
	handler := cors.DevelopmentHandler.Handler(mux)

	s.server = &http.Server{
		Addr:    s.cfg.Connect.Address(),
		Handler: handler,
	}

	errs := make(chan error, 1)

	go func() {
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			errs <- err
		}
	}()

	select {
	case err := <-errs:
		return err
	case <-time.After(1 * time.Second):
	}

	return nil
}

// Stop stops the MonitorQueryService
func (s *MonitorQueryService) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	// TODO close dependencies

	return nil
}
