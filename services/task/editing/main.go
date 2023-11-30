package editing

import "github.com/taylow/awaik-backend/services"

// ServiceName holds the unique name of the service
const ServiceName = "TaskEditingService"

// init inisialises the service and registers it with the service registry
func init() {
	services.Register(ServiceName, TaskEditingService{})
}

// TaskEditingService represents the service that edits tasks
type TaskEditingService struct{}

// Name returns the name of the service
func (s TaskEditingService) Name() string {
	return ServiceName
}

// Start starts the TaskEditingService
func (s TaskEditingService) Start() error {
	return nil
}

// Stop stops the TaskEditingService
func (s TaskEditingService) Stop() error {
	return nil
}
