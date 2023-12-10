package monitor

import (
	"context"

	"github.com/taylow/awaik-backend/services/monitor/domain"
)

var _ domain.MonitorWriteRepository = (*noOpMonitorRepository)(nil)

// noOpMonitorRepository is a no-op implementation of the monitor repository
type noOpMonitorRepository struct{}

// NewNoOpMonitorRepository initialises a new in-memory monitor repository
func NewNoOpMonitorRepository() *noOpMonitorRepository {
	return &noOpMonitorRepository{}
}

// Set implements domain.MonitorRepository
func (r *noOpMonitorRepository) Set(ctx context.Context, monitor *domain.Monitor) error {
	return nil
}

// Get implements domain.MonitorRepository
func (r *noOpMonitorRepository) Get(ctx context.Context, id string) (*domain.Monitor, error) {
	return &domain.Monitor{}, nil
}

// ListByProject implements domain.MonitorRepository
func (r *noOpMonitorRepository) ListByProject(ctx context.Context, projectID string) ([]*domain.Monitor, error) {
	return []*domain.Monitor{}, nil
}

// UpdateStatus implements domain.MonitorRepository
func (r *noOpMonitorRepository) UpdateStatus(ctx context.Context, id string, status domain.Status) error {
	return nil
}

// Delete implements domain.MonitorRepository
func (r *noOpMonitorRepository) Delete(ctx context.Context, id string, isSoft bool) error {
	return nil
}
