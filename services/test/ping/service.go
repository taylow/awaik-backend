package ping

import (
	"context"
	"net/http"
	"time"

	"github.com/taylow/awaik-backend/internal/cors"
	"github.com/taylow/awaik-backend/internal/gen/proto/test/v1/testv1connect"
	"github.com/taylow/awaik-backend/services"
	"github.com/taylow/awaik-backend/services/test/ping/handler"
)

// ServiceName holds the unique name of the service
const (
	ServiceName = "TestPingService"
	Emoji       = "🧪 📡"
	Address     = ":8008"
)

// init inisialises the service and registers it with the service registry
func init() {
	services.Register(ServiceName, &TestPingService{})
}

// TestPingService represents the service that edits monitors
type TestPingService struct {
	server *http.Server
}

// Name returns the name of the service
func (s *TestPingService) Name() string {
	return ServiceName
}

// Emoji returns the emoji of the service
func (s *TestPingService) Emoji() string {
	return Emoji
}

// Address returns the address of the service
func (s *TestPingService) Address() string {
	return Address
}

// Start starts the TestPingService
func (s *TestPingService) Start() error {
	mux := http.NewServeMux()
	serviceHandler := handler.NewConnectHandler()

	mux.Handle(testv1connect.NewPingServiceHandler(serviceHandler))
	handler := cors.DevelopmentHandler.Handler(mux)

	s.server = &http.Server{
		Addr:    Address,
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

// Stop stops the TestPingService
func (s *TestPingService) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
