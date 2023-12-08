package config

import "fmt"

// Service holds the configuration for the service
type Service struct {
	// Name holds the name of the service
	Name string `yaml:"name"`
	// Emoji holds the emoji of the service
	Emoji string `yaml:"emoji"`
}

// Connect holds connect configuration values
type Connect struct {
	// Host holds the host address of the service
	Host string `yaml:"host"`
	// Port holds the port of the service
	Port int `yaml:"port"`
}

// Address returns the address of the service
func (c *Connect) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
