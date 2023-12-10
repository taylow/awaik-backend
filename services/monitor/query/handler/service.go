package handler

import (
	"context"

	"github.com/taylow/awaik-backend/services/monitor/domain"
)

// QueryService is an abstraction on monitor query business logic
type QueryService interface {
	// Read fetches a persisted monitor
	Read(ctx context.Context, id string) (*domain.Monitor, error)
	// ListByProject retrieves all monitors for a specific project
	ListByProject(ctx context.Context, projectID string) ([]*domain.Monitor, error)
}
