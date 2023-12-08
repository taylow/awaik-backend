package message

import (
	"github.com/taylow/awaik-backend/internal/config"
	"github.com/taylow/awaik-backend/services/monitor/command/domain"
)

// messageBroker is a monitor repository
type messageBroker struct {
	domain.MessageBroker
}

// BrokerFromConfig creates a message broker from the configuration
func BrokerFromConfig(cfg *config.MessageBroker) (*messageBroker, error) {
	switch cfg.Type {
	case config.BrokerTypeNoop:
		return &messageBroker{NewNoopMessageBroker()}, nil
	// case config.BrokerTypeNATS:
	// 	return &messageBroker{NewNATSMessageBroker()}, nil
	// case config.BrokerTypeKafka:
	// 	return &messageBroker{NewKafkaMessageBroker()}, nil
	default:
		return nil, domain.ErrUnknownMessageBrokerType
	}
}
