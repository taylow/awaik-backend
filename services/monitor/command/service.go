package command

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
	"github.com/taylow/awaik-backend/services/monitor/command/config"
	"github.com/taylow/awaik-backend/services/monitor/command/handler"
	"github.com/taylow/awaik-backend/services/monitor/command/infra/message"
	"github.com/taylow/awaik-backend/services/monitor/command/infra/monitor"
)

// ServiceName holds the unique name of the service
const ServiceName = "MonitorCommandService"
const Emoji = "👀 📝"

// init inisialises the service and registers it with the service registry
func init() {
	// TODO load config from file/env
	cfg := &config.MonitorCommandServiceConfig{
		Service: &sharedconfig.Service{
			Name:  ServiceName,
			Emoji: Emoji,
		},
		Connect: &sharedconfig.Connect{
			Host: "",
			Port: 8080,
		},
		MonitorRepository: &sharedconfig.MonitorRepository{
			Type: "memory",
		},
		MessageBroker: &sharedconfig.MessageBroker{
			Type: "noop",
		},
	}

	services.Register(ServiceName, &MonitorCommandService{
		cfg: cfg,
	})
}

// MonitorCommandService represents the service that edits monitors
type MonitorCommandService struct {
	cfg *config.MonitorCommandServiceConfig

	service handler.CommandService
	server  *http.Server
}

// Name returns the name of the service
func (s *MonitorCommandService) Name() string {
	return ServiceName
}

// Emoji returns the emoji of the service
func (s *MonitorCommandService) Emoji() string {
	return Emoji
}

// Address returns the address of the service
func (s *MonitorCommandService) Address() string {
	return s.cfg.Connect.Address()
}

// Start starts the MonitorCommandService
func (s *MonitorCommandService) Start() error {
	monitorRepo, err := monitor.RepositoryFromConfig(s.cfg.MonitorRepository)
	if err != nil {
		return err
	}

	messageBroker, err := message.BrokerFromConfig(s.cfg.MessageBroker)
	if err != nil {
		return err
	}

	s.service = NewCommandService(monitorRepo, messageBroker)

	serviceHandler := handler.NewConnectHandler(s.service)

	opts := []connect.HandlerOption{
		connect.WithInterceptors(interceptor.WithTimeout(5 * time.Second)),
		connect.WithInterceptors(interceptor.NewLogInterceptor()),
		connect.WithInterceptors(interceptor.NewValidateInterceptor()),
	}

	mux := http.NewServeMux()
	mux.Handle(monitorv1connect.NewMonitorCommandServiceHandler(serviceHandler, opts...))
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

// Stop stops the MonitorCommandService
func (s *MonitorCommandService) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	// TODO close dependencies

	return nil
}
