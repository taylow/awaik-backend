package domain

import (
	"context"
	"log"
)

// ProfiledService wraps service with logging and tracing
type ProfiledService struct {
	service Service
}

// NewProfiledService initialises a profiled service with the provided dependencies
//
// TODO add logger dependency and replace `log` calls
func NewProfiledService(service Service) *ProfiledService {
	return &ProfiledService{
		service: service,
	}
}

// Create creates a new monitor, persists its state and sends a monitor created event
func (s ProfiledService) Create(
	ctx context.Context,
	projectID, name, description, interval string,
	regions []string,
	protocol Protocol,
	config ProtocolConfig,
) (*Monitor, error) {
	log.Println("creating monitor")

	monitor, err := s.service.Create(
		ctx,
		projectID,
		name,
		description,
		interval,
		regions,
		protocol,
		config,
	)
	if err != nil {
		log.Println("failed to create monitor", err)
		return nil, err
	}

	log.Println("successfully created monitor")

	return monitor, nil
}

// Read fetches a persisted monitor
func (s ProfiledService) Read(ctx context.Context, id string) (*Monitor, error) {
	log.Println("reading monitor")

	monitor, err := s.service.Read(ctx, id)
	if err != nil {
		log.Println("failed to read monitor", err)
		return nil, err
	}

	log.Println("successfully read monitor")

	return monitor, nil
}

// Update modifies the values of a persisted monitor and fires a monitor updated event
func (s ProfiledService) Update(ctx context.Context, monitor *Monitor) (*Monitor, error) {
	log.Println("updating monitor")

	monitor, err := s.service.Update(ctx, monitor)
	if err != nil {
		log.Println("failed to update monitor", err)
		return nil, err
	}

	log.Println("successfully updated monitor")

	return monitor, nil
}

// Delete removes the persisted monitor and fires a monitor deleted event
func (s ProfiledService) Delete(ctx context.Context, id string) error {
	log.Println("deleting monitor")

	err := s.service.Delete(ctx, id)
	if err != nil {
		log.Println("failed to delete monitor", err)
		return err
	}

	log.Println("successfully deleted monitor")

	return nil
}

// SoftDelete marks the persisted monitor as deleted and fires a monitor deleted event
func (s ProfiledService) SoftDelete(ctx context.Context, id string) error {
	log.Println("soft deleting monitor")

	err := s.service.SoftDelete(ctx, id)
	if err != nil {
		log.Println("failed to soft delete monitor", err)
		return err
	}

	log.Println("successfully soft deleted monitor")

	return nil
}

// Pause marks the persisted monitor as paused and fires a monitor paused event
func (s ProfiledService) Pause(ctx context.Context, id string) error {
	log.Println("pausing monitor")

	err := s.service.Pause(ctx, id)
	if err != nil {
		log.Println("failed to pause monitor", err)
		return err
	}

	log.Println("successfully paused monitor")

	return nil
}

// Resume marks the persisted monitor as pending and fires a monitor resumed event
func (s ProfiledService) Resume(ctx context.Context, id string) error {
	log.Println("resuming monitor")

	err := s.service.Resume(ctx, id)
	if err != nil {
		log.Println("failed to resume monitor", err)
		return err
	}

	log.Println("successfully resumed monitor")

	return nil
}
