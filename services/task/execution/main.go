package execution

import "github.com/taylow/awaik-backend/services"

// ServiceName holds the unique name of the service
const ServiceName = "TaskExecutionService"

// init inisialises the service and registers it with the service registry
func init() {
	services.Register(ServiceName, TaskExecutionService{})
}

// TaskExecutionService represents the service that executes tasks
type TaskExecutionService struct{}

// Name returns the name of the service
func (s TaskExecutionService) Name() string {
	return ServiceName
}

// Start starts the TaskExecutionService
func (s TaskExecutionService) Start() error {
	return nil
}

// Stop stops the TaskExecutionService
func (s TaskExecutionService) Stop() error {
	return nil
}
