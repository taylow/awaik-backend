package monitor

import (
	"github.com/taylow/awaik-backend/internal/config"
	"github.com/taylow/awaik-backend/services/monitor/domain"

	"github.com/pkg/errors"
)

// monitorReadRepository is a monitor repository
type monitorReadRepository struct {
	domain.MonitorReadRepository
}

// RepositoryFromConfig creates a monitor repository from the configuration
func RepositoryFromConfig(cfg *config.MonitorRepository) (*monitorReadRepository, error) {
	switch cfg.Type {
	case config.RepositoryTypeNoOp:
		return &monitorReadRepository{NewNoOpMonitorRepository()}, nil
	case config.RepositoryTypeMemory:
		return &monitorReadRepository{NewInMemoryMonitorRepository()}, nil
	default:
		return nil, errors.Wrap(domain.ErrUnknownMonitorRepositoryType, string(cfg.Type))
	}
}
