package config

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
