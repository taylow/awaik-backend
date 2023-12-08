package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Service is an abstraction on monitor business logic
type Service interface {
	// Create creates a new monitor, persists its state and sends a monitor created event
	Create(
		ctx context.Context,
		projectID, name, description, interval string,
		regions []string,
		protocol Protocol,
		config ProtocolConfig,
	) (*Monitor, error)
	// Read fetches a persisted monitor
	Read(ctx context.Context, id string) (*Monitor, error)
	// Update modifies the values of a persisted monitor and fires a monitor updated event
	Update(ctx context.Context, monitor *Monitor) (*Monitor, error)
	// Delete removes the persisted monitor and fires a monitor deleted event
	Delete(ctx context.Context, id string) error
	// SoftDelete marks the persisted monitor as deleted and fires a monitor deleted event
	SoftDelete(ctx context.Context, id string) error
	// Pause marks the persisted monitor as paused and fires a monitor paused event
	Pause(ctx context.Context, id string) error
	// Resume marks the persisted monitor as pending and fires a monitor resumed event
	Resume(ctx context.Context, id string) error
}

// service holds business logic for monitors
type service struct {
	repo  MonitorRepository
	event MessageBroker
}

// NewService initialises a service with the provided dependencies
func NewService(repo MonitorRepository, event MessageBroker) *service {
	return &service{
		repo:  repo,
		event: event,
	}
}

// Create creates a new monitor, persists its state and sends a monitor created event
func (s service) Create(
	ctx context.Context,
	projectID, name, description, interval string,
	regions []string,
	protocol Protocol,
	config ProtocolConfig,
) (*Monitor, error) {
	monitor := &Monitor{
		ID:          uuid.New().String(),
		ProjectID:   projectID,
		Name:        name,
		Description: description,
		Interval:    interval,
		Status:      StatusPending,
		Regions:     regions,
		Protocol:    protocol,
		HTTP:        &HTTP{},
		ICMP:        &ICMP{},
		Port:        &Port{},
		Browser:     &BrowserAutomation{},
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

// Read fetches a persisted monitor
func (s service) Read(ctx context.Context, id string) (*Monitor, error) {
	monitor, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return monitor, nil
}

// Update modifies the values of a persisted monitor and fires a monitor updated event
func (s service) Update(ctx context.Context, monitor *Monitor) (*Monitor, error) {
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
func (s service) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id, false); err != nil {
		return err
	}

	if err := s.event.Deleted(ctx, id, false); err != nil {
		return err
	}

	return nil
}

// SoftDelete marks the persisted monitor as deleted and fires a monitor deleted event
func (s service) SoftDelete(ctx context.Context, id string) error {
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
func (s service) Pause(ctx context.Context, id string) error {
	err := s.repo.UpdateStatus(ctx, id, StatusPaused)
	if err != nil {
		return err
	}

	if err := s.event.Paused(ctx, id); err != nil {
		return err
	}

	return nil
}

// Resume marks the persisted monitor as pending and fires a monitor resumed event
func (s service) Resume(ctx context.Context, id string) error {
	err := s.repo.UpdateStatus(ctx, id, StatusPending)
	if err != nil {
		return err
	}

	if err := s.event.Resumed(ctx, id); err != nil {
		return err
	}

	return nil
}
