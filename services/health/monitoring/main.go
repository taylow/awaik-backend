package editing

import "github.com/taylow/awaik-backend/services"

// ServiceName holds the unique name of the service.
const ServiceName = "HealthMonitoringService"

// init inisialises the service and registers it with the service registry.
func init() {
	services.Register(ServiceName, HealthMonitoringService{})
}

// HealthMonitoringService represents the service that monitors service health.
type HealthMonitoringService struct{}

// Name returns the name of the service.
func (s HealthMonitoringService) Name() string {
	return ServiceName
}

// Start starts the HealthMonitoringService.
func (s HealthMonitoringService) Start() error {
	return nil
}

// Stop stops the HealthMonitoringService.
func (s HealthMonitoringService) Stop() error {
	return nil
}
