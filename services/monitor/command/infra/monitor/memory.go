package monitor

import (
	"context"
	"sync"
	"time"

	"github.com/taylow/awaik-backend/services/monitor/command/domain"
)

var _ domain.MonitorRepository = (*inMemoryMonitorRepository)(nil)

// inMemoryMonitorRepository is an in-memory implementation of the monitor repository
type inMemoryMonitorRepository struct {
	monitors map[string]*domain.Monitor
	mx       *sync.RWMutex
}

// NewInMemoryMonitorRepository initialises a new in-memory monitor repository
func NewInMemoryMonitorRepository() *inMemoryMonitorRepository {
	return &inMemoryMonitorRepository{
		monitors: make(map[string]*domain.Monitor),
		mx:       &sync.RWMutex{},
	}
}

// Set implements domain.MonitorRepository
func (r *inMemoryMonitorRepository) Set(ctx context.Context, monitor *domain.Monitor) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.monitors[monitor.ID] = monitor

	return nil
}

// Get implements domain.MonitorRepository
func (r *inMemoryMonitorRepository) Get(ctx context.Context, id string) (*domain.Monitor, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	monitor, ok := r.monitors[id]
	if !ok {
		return nil, domain.ErrMonitorNotFound
	}

	return monitor, nil
}

// ListByProject implements domain.MonitorRepository
func (r *inMemoryMonitorRepository) ListByProject(ctx context.Context, projectID string) ([]*domain.Monitor, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	var monitors []*domain.Monitor
	for _, monitor := range r.monitors {
		if monitor.ProjectID == projectID {
			monitors = append(monitors, monitor)
		}
	}

	return monitors, nil
}

// UpdateStatus implements domain.MonitorRepository
func (r *inMemoryMonitorRepository) UpdateStatus(ctx context.Context, id string, status domain.Status) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.monitors[id].Status = status
	r.monitors[id].UpdatedAt = time.Now()

	return nil
}

// Delete implements domain.MonitorRepository
func (r *inMemoryMonitorRepository) Delete(ctx context.Context, id string, isSoft bool) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	if isSoft {
		r.monitors[id].Status = domain.StatusDeleted
		r.monitors[id].UpdatedAt = time.Now()
		r.monitors[id].DeletedAt = time.Now()

		return nil
	}

	delete(r.monitors, id)

	return nil
}
