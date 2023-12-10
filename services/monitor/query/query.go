package query

import (
	"context"

	"github.com/taylow/awaik-backend/services/monitor/domain"
)

// queryService holds business logic for monitors
type queryService struct {
	repo domain.MonitorReadRepository
}

// NewQueryService initialises a service with the provided dependencies
func NewQueryService(repo domain.MonitorReadRepository) *queryService {
	return &queryService{
		repo: repo,
	}
}

// Read fetches a persisted monitor
func (s queryService) Read(ctx context.Context, id string) (*domain.Monitor, error) {
	monitor, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return monitor, nil
}

// ListByProject retrieves all monitors for a specific project
func (s queryService) ListByProject(ctx context.Context, projectID string) ([]*domain.Monitor, error) {
	monitors, err := s.repo.ListByProject(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return monitors, nil
}
