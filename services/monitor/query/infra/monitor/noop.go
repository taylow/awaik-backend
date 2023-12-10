package monitor

import (
	"context"

	"github.com/taylow/awaik-backend/services/monitor/domain"
)

var _ domain.MonitorReadRepository = (*noOpMonitorReadRepository)(nil)

// noOpMonitorReadRepository is a no-op implementation of the monitor repository
type noOpMonitorReadRepository struct{}

// NewNoOpMonitorRepository initialises a new in-memory monitor repository
func NewNoOpMonitorRepository() *noOpMonitorReadRepository {
	return &noOpMonitorReadRepository{}
}

// Get implements domain.MonitorRepository
func (r *noOpMonitorReadRepository) Get(ctx context.Context, id string) (*domain.Monitor, error) {
	return &domain.Monitor{}, nil
}

// ListByProject implements domain.MonitorRepository
func (r *noOpMonitorReadRepository) ListByProject(ctx context.Context, projectID string) ([]*domain.Monitor, error) {
	return []*domain.Monitor{}, nil
}
