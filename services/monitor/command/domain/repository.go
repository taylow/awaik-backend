package domain

import "context"

// MonitorRepository is an abstraction on monitor persistance
type MonitorRepository interface {
	// Set persists a monitor
	Set(ctx context.Context, monitor *Monitor) error
	// Get retrieves a monitor
	Get(ctx context.Context, id string) (*Monitor, error)
	// ListByProject retrieves all monitors for a specific project
	ListByProject(ctx context.Context, projectID string) ([]*Monitor, error)
	// Delete removes a monitor
	Delete(ctx context.Context, id string, isSoft bool) error
	// UpdateStatus updates the status of a monitor
	UpdateStatus(ctx context.Context, id string, status Status) error
}
