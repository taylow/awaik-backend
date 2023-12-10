package handler

import (
	"context"

	"github.com/taylow/awaik-backend/services/monitor/domain"
)

// CommandService is an abstraction on monitor command business logic
type CommandService interface {
	// Create creates a new monitor, persists its state and sends a monitor created event
	Create(
		ctx context.Context,
		projectID, name, description, interval string,
		regions []string,
		protocol domain.Protocol,
		config domain.ProtocolConfig,
	) (*domain.Monitor, error)
	// Update modifies the values of a persisted monitor and fires a monitor updated event
	Update(ctx context.Context, monitor *domain.Monitor) (*domain.Monitor, error)
	// Delete removes the persisted monitor and fires a monitor deleted event
	Delete(ctx context.Context, id string) error
	// SoftDelete marks the persisted monitor as deleted and fires a monitor deleted event
	SoftDelete(ctx context.Context, id string) error
	// Pause marks the persisted monitor as paused and fires a monitor paused event
	Pause(ctx context.Context, id string) error
	// Resume marks the persisted monitor as pending and fires a monitor resumed event
	Resume(ctx context.Context, id string) error
}
