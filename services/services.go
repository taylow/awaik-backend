package services

import (
	"fmt"
	"strings"
)

// Services stores a map of service names to Service instances and is populated by each service `init()` function
var Services = ServiceRegistry{}

// ServiceRegistry is an alias for storing a map of service names against Service instances
type ServiceRegistry map[string]Service

// Names returns a slice of service names
func (s ServiceRegistry) Names() []string {
	names := make([]string, 0, len(s))
	for name := range s {
		names = append(names, name)
	}
	return names
}

// NamesWithEmojis returns a slice of service names prefixed with their emoji
func (s ServiceRegistry) NamesWithEmojis() []string {
	names := make([]string, 0, len(s))
	for name, service := range s {
		names = append(names, fmt.Sprintf("%s  %s", service.Emoji(), name))
	}
	return names
}

// Filter filters the service list based on flags
func (s ServiceRegistry) Filter(servicesToKeep SliceFlag, keepAll bool) ServiceRegistry {
	if keepAll {
		return s
	}

	if servicesToKeep == nil {
		return ServiceRegistry{}
	}

	filteredServices := map[string]Service{}
	for _, service := range servicesToKeep {
		if s, ok := s[service]; ok {
			filteredServices[service] = s
		} else {
			fmt.Printf("could not find service %q\n", service)
		}
	}
	return filteredServices
}

// Register makes a service available by the provided name. If Register is called twice with the same name, no name, or if service is nil, it panics
func Register(name string, service Service) {
	if name == "" {
		panic("service registration called with no name")
	}
	if service == nil {
		panic(fmt.Sprintf("nil service registration for service %q", name))
	}
	if _, ok := Services[service.Name()]; ok {
		panic(fmt.Sprintf("duplicate service registration for service %q", name))
	}
	fmt.Printf("%s registered\n", name)
	Services[name] = service
}

// Service represents an isolated unit capable of hosing some service
type Service interface {
	// Name returns the name of the service
	Name() string

	// Emoji returns the emoji of the service
	Emoji() string

	// Address returns the address of the service
	Address() string

	// Start starts the service
	Start() error

	// Stop stops the service
	Stop() error
}

// SliceFlag holds a slice of strings and allows multiple flag values to be set
type SliceFlag []string

// String returns a comma separated list of values
func (s SliceFlag) String() string {
	return strings.Join(s, ", ")
}

// Set appends a value onto a slide of strings, skipping over any duplicates
func (s *SliceFlag) Set(value string) error {
	for _, v := range *s {
		if v == value {
			return nil
		}
	}
	*s = append(*s, value)
	return nil
}
