package monitor

import (
	"github.com/taylow/awaik-backend/services/monitor/command/config"
	"github.com/taylow/awaik-backend/services/monitor/command/domain"

	"github.com/pkg/errors"
)

// monitorRepository is a monitor repository
type monitorRepository struct {
	domain.MonitorRepository
}

// RepositoryFromConfig creates a monitor repository from the configuration
func RepositoryFromConfig(cfg *config.MonitorRepository) (*monitorRepository, error) {
	switch cfg.Type {
	case config.RepositoryTypeNoOp:
		return &monitorRepository{NewNoOpMonitorRepository()}, nil
	case config.RepositoryTypeMemory:
		return &monitorRepository{NewInMemoryMonitorRepository()}, nil
	default:
		return nil, errors.Wrap(domain.ErrUnknownMonitorRepositoryType, string(cfg.Type))
	}
}
