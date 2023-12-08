package domain

import "errors"

var (
	ErrMonitorNotFound              = errors.New("monitor not found")
	ErrUnknownMonitorRepositoryType = errors.New("unknown monitor repository type")
	ErrUnknownMessageBrokerType     = errors.New("unknown message broker type")
)
