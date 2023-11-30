package services

import (
	"fmt"
	"strings"
)

// Services stores a map of service names to Service instances and is populated by each service `init()` function.
var Services = services{}

// services is an alias for storing a map of service names against Service instances.
type services map[string]Service

// Register makes a service available by the provided name. If Register is called twice with the same name, no name, or if service is nil, it panics.
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

// Filter filters the service list based on flags.
func (s services) Filter(allServices map[string]Service, servicesToKeep SliceFlag, keepAll bool) map[string]Service {
	if keepAll {
		return allServices
	}

	filteredServices := map[string]Service{}
	for _, service := range servicesToKeep {
		if s, ok := allServices[service]; ok {
			filteredServices[service] = s
		} else {
			fmt.Printf("could not find service %q\n", service)
		}
	}
	return filteredServices
}

// Service represents an isolated unit capable of hosing some service.
type Service interface {
	// Name returns the name of the service
	Name() string

	// Start starts the service
	Start() error

	// Stop stops the service
	Stop() error
}

// SliceFlag holds a slice of strings and allows multiple flag values to be set.
type SliceFlag []string

// String returns a comma separated list of values.
func (s SliceFlag) String() string {
	return strings.Join(s, ", ")
}

// Set appends a value onto a slide of strings, skipping over any duplicates.
func (s *SliceFlag) Set(value string) error {
	for _, v := range *s {
		if v == value {
			return nil
		}
	}
	*s = append(*s, value)
	return nil
}
