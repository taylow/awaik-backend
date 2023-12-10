package message

import (
	"context"

	"github.com/taylow/awaik-backend/services/monitor/domain"
)

var _ domain.MessageBroker = (*noop)(nil)

// NewNoopMessageBroker creates a new noop message broker
func NewNoopMessageBroker() domain.MessageBroker {
	return &noop{}
}

// noop is a no-op implementation of the message broker
type noop struct{}

// Created implements domain.MessageBroker
func (*noop) Created(ctx context.Context, monitor *domain.Monitor) error {
	return nil
}

// Deleted implements domain.MessageBroker
func (*noop) Deleted(ctx context.Context, id string, isSoft bool) error {
	return nil
}

// MonitorRescheduled implements domain.MessageBroker
func (*noop) MonitorRescheduled(ctx context.Context, monitor *domain.Monitor) error {
	return nil
}

// MonitorScheduled implements domain.MessageBroker
func (*noop) MonitorScheduled(ctx context.Context, monitor *domain.Monitor) error {
	return nil
}

// Paused implements domain.MessageBroker
func (*noop) Paused(ctx context.Context, id string) error {
	return nil
}

// Resumed implements domain.MessageBroker
func (*noop) Resumed(ctx context.Context, id string) error {
	return nil
}

// Updated implements domain.MessageBroker
func (*noop) Updated(ctx context.Context, monitor *domain.Monitor) error {
	return nil
}
