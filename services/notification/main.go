package editing

import "github.com/taylow/awaik-backend/services"

// ServiceName holds the unique name of the service.
const ServiceName = "NotificationService"

// init inisialises the service and registers it with the service registry.
func init() {
	services.Register(ServiceName, NotificationService{})
}

// NotificationService represents the service that handles notifications.
type NotificationService struct{}

// Name returns the name of the service.
func (s NotificationService) Name() string {
	return ServiceName
}

// Start starts the NotificationService.
func (s NotificationService) Start() error {
	return nil
}

// Stop stops the NotificationService.
func (s NotificationService) Stop() error {
	return nil
}
