package editing

import "github.com/taylow/awaik-backend/services"

// ServiceName holds the unique name of the service
const ServiceName = "TaskStatusService"

// init inisialises the service and registers it with the service registry
func init() {
	services.Register(ServiceName, TaskStatusService{})
}

// TaskStatusService represents the service that edits tasks
type TaskStatusService struct{}

// Name returns the name of the service
func (s TaskStatusService) Name() string {
	return ServiceName
}

// Start starts the TaskEditingService
func (s TaskStatusService) Start() error {
	return nil
}

// Stop stops the TaskEditingService
func (s TaskStatusService) Stop() error {
	return nil
}
