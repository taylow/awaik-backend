// Config provides the configuration for the command service
package config

import "github.com/taylow/awaik-backend/internal/config"

// Service holds the configuration for the command service
type Service struct {
	// Service holds the configuration for the service
	*config.Service
	// Connect holds the configuration Connect handler
	*config.Connect
	// MonitorRepository holds the configuration for the monitor repository
	MonitorRepository *MonitorRepository `yaml:"monitor_store"`
	// MessageBroker
	MessageBroker *MessageBroker `yaml:"message_broker"`
}

// MonitorRepository holds the configuration for the monitor store
type MonitorRepository struct {
	// Type holds the type of the monitor store
	Type RepositoryType `yaml:"type"`
}

// RepositoryType holds the type of the repository
type RepositoryType string

const (
	// RepositoryTypeNoOp represents the noop repository
	RepositoryTypeNoOp RepositoryType = "noop"
	// RepositoryTypeMemory represents the memory store
	RepositoryTypeMemory RepositoryType = "memory"
)

// MessageBroker holds the configuration for the message broker
type MessageBroker struct {
	// Type holds the type of the message broker
	Type BrokerType `yaml:"type"`
	// NATS holds the configuration for the NATS broker
	NATS NATS `yaml:"nats"`
	// Kafka holds the configuration for the Kafka broker
	Kafka Kafka `yaml:"kafka"`
}

// BrokerType holds the type of the broker
type BrokerType string

const (
	// BrokerTypeNoop represents the noop broker
	BrokerTypeNoop BrokerType = "noop"
	// BrokerTypeNATS represents the NATS broker
	BrokerTypeNATS BrokerType = "nats"
	// BrokerTypeKafka represents the NATS broker
	BrokerTypeKafka BrokerType = "kafka"
)

// NATS holds the configuration for the NATS broker
type NATS struct {
	// TODO implement NATS configuration
}

// Kafka holds the configuration for the Kafka broker
type Kafka struct {
	// TODO implement Kafka configuration
}
