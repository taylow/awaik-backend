package config

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
