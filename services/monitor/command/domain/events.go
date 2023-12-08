package domain

import "context"

// MessageBroker is an abstraction on the monitor event message broker
type MessageBroker interface {
	// Created sends a monitor created event
	Created(ctx context.Context, monitor *Monitor) error
	// Updated sends a monitor updated event
	Updated(ctx context.Context, monitor *Monitor) error
	// Deleted sends a monitor deleted event
	Deleted(ctx context.Context, id string, isSoft bool) error
	// Paused sends a monitor paused event
	Paused(ctx context.Context, id string) error
	// Resumed sends a monitor resumed event
	Resumed(ctx context.Context, id string) error
	// MonitorScheduled sends a monitor scheduled event
	MonitorScheduled(ctx context.Context, monitor *Monitor) error
	// MonitorRescheduled sends a monitor rescheduled event
	MonitorRescheduled(ctx context.Context, monitor *Monitor) error
}
