package command

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/taylow/awaik-backend/services/monitor/domain"
)

// commandService is an application service that holds business logic for monitors
type commandService struct {
	repo  domain.MonitorWriteRepository
	event domain.MessageBroker
}

// NewCommandService initialises a service with the provided dependencies
func NewCommandService(repo domain.MonitorWriteRepository, event domain.MessageBroker) *commandService {
	return &commandService{
		repo:  repo,
		event: event,
	}
}

// Create creates a new monitor, persists its state and sends a monitor created event
func (s commandService) Create(
	ctx context.Context,
	projectID, name, description, interval string,
	regions []string,
	protocol domain.Protocol,
	config domain.ProtocolConfig,
) (*domain.Monitor, error) {
	monitor := &domain.Monitor{
		ID:          uuid.New().String(),
		ProjectID:   projectID,
		Name:        name,
		Description: description,
		Interval:    interval,
		Status:      domain.StatusPending,
		Regions:     regions,
		Protocol:    protocol,
		HTTP:        &domain.HTTP{},
		ICMP:        &domain.ICMP{},
		Port:        &domain.Port{},
		Browser:     &domain.BrowserAutomation{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Time{},
	}

	config.Apply(monitor)

	if err := s.repo.Set(ctx, monitor); err != nil {
		return nil, err
	}

	if err := s.event.Created(ctx, monitor); err != nil {
		return nil, err
	}

	return monitor, nil
}

// Update modifies the values of a persisted monitor and fires a monitor updated event
func (s commandService) Update(ctx context.Context, monitor *domain.Monitor) (*domain.Monitor, error) {
	err := s.repo.Set(ctx, monitor)
	if err != nil {
		return nil, err
	}

	if err := s.event.Updated(ctx, monitor); err != nil {
		return nil, err
	}

	return monitor, nil
}

// Delete removes the persisted monitor and fires a monitor deleted event
func (s commandService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id, false); err != nil {
		return err
	}

	if err := s.event.Deleted(ctx, id, false); err != nil {
		return err
	}

	return nil
}

// SoftDelete marks the persisted monitor as deleted and fires a monitor deleted event
func (s commandService) SoftDelete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id, true)
	if err != nil {
		return err
	}

	if err := s.event.Deleted(ctx, id, true); err != nil {
		return err
	}

	return nil
}

// Pause marks the persisted monitor as paused and fires a monitor paused event
func (s commandService) Pause(ctx context.Context, id string) error {
	err := s.repo.UpdateStatus(ctx, id, domain.StatusPaused)
	if err != nil {
		return err
	}

	if err := s.event.Paused(ctx, id); err != nil {
		return err
	}

	return nil
}

// Resume marks the persisted monitor as pending and fires a monitor resumed event
func (s commandService) Resume(ctx context.Context, id string) error {
	err := s.repo.UpdateStatus(ctx, id, domain.StatusPending)
	if err != nil {
		return err
	}

	if err := s.event.Resumed(ctx, id); err != nil {
		return err
	}

	return nil
}
