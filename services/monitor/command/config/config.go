// Config provides the configuration for the command service
package config

import "github.com/taylow/awaik-backend/internal/config"

// MonitorCommandServiceConfig holds the configuration for the command service
type MonitorCommandServiceConfig struct {
	// Service holds the configuration for the service
	*config.Service
	// Connect holds the configuration Connect handler
	*config.Connect
	// MonitorRepository holds the configuration for the monitor repository
	MonitorRepository *config.MonitorRepository `yaml:"monitor_repository"`
	// MessageBroker
	MessageBroker *config.MessageBroker `yaml:"message_broker"`
}
