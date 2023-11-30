package scheduling

import (
	"github.com/taylow/awaik-backend/services"
)

// ServiceName holds the unique name of the service.
const ServiceName = "TaskSchedulingService"

// init inisialises the service and registers it with the service registry.
func init() {
	services.Register(ServiceName, TaskSchedulingService{})
}

// TaskSchedulingService represents the service that schedules tasks.
type TaskSchedulingService struct{}

// Name returns the name of the service.
func (s TaskSchedulingService) Name() string {
	return ServiceName
}

// Start starts the TaskSchedulingService.
func (s TaskSchedulingService) Start() error {
	return nil
}

// Stop stops the TaskSchedulingService.
func (s TaskSchedulingService) Stop() error {
	return nil
}
