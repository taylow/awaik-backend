package editing

import "github.com/taylow/awaik-backend/services"

// ServiceName holds the unique name of the service
const ServiceName = "HealthRecoveryService"

// init inisialises the service and registers it with the service registry
func init() {
	services.Register(ServiceName, HealthRecoveryService{})
}

// HealthRecoveryService represents the service that recovers tasks from down services
type HealthRecoveryService struct{}

// Name returns the name of the service
func (s HealthRecoveryService) Name() string {
	return ServiceName
}

// Start starts the HealthRecoveryService
func (s HealthRecoveryService) Start() error {
	return nil
}

// Stop stops the HealthRecoveryService
func (s HealthRecoveryService) Stop() error {
	return nil
}
