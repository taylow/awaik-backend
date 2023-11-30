package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Define a mock service for testing
type MockService struct {
	name string
}

func (s *MockService) Name() string {
	return s.name
}

func (s *MockService) Start() error {
	return nil
}

func (s *MockService) Stop() error {
	return nil
}

// TestRegister tests the Register function with testify/assert
func TestRegister(t *testing.T) {
	// Test valid registration
	t.Run("Valid Registration", func(t *testing.T) {
		resetServices()
		mockService := &MockService{name: "MockService"}
		Register(mockService.name, mockService)
		assert.NotNil(t, Services[mockService.name], "Service should be registered")
	})

	// Test registration with no name
	t.Run("Registration with No Name", func(t *testing.T) {
		resetServices()
		mockService := &MockService{name: "MockService"}
		assert.PanicsWithValue(t, "service registration called with no name", func() {
			Register("", mockService)
		})
	})

	// Test registration with nil service
	t.Run("Registration with Nil Service", func(t *testing.T) {
		resetServices()
		assert.PanicsWithValue(t, fmt.Sprintf("nil service registration for service %q", "MockService"), func() {
			Register("MockService", nil)
		})
	})

	// Test duplicate registration
	t.Run("Duplicate Registration", func(t *testing.T) {
		resetServices()
		mockService1 := &MockService{name: "MockService"}
		mockService2 := &MockService{name: "MockService"}
		Register(mockService1.name, mockService1)
		assert.PanicsWithValue(t, fmt.Sprintf("duplicate service registration for service %q", "MockService"), func() {
			Register(mockService2.name, mockService2)
		})
	})
}

// resetServices resets the Services map before each test
func resetServices() {
	Services = make(map[string]Service)
}

// assertPanic checks if the given function panics with the expected message
func assertPanic(t *testing.T, expectedMessage string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic with message '%s', but no panic occurred", expectedMessage)
		} else if r != expectedMessage {
			t.Errorf("Expected panic with message '%s', but got '%v'", expectedMessage, r)
		}
	}()

	f()
}

// TestFilter tests the Filter method of the services type
func TestFilter(t *testing.T) {
	// Test filtering with keepAll set to true
	t.Run("Filter with keepAll", func(t *testing.T) {
		s := services{
			"Service1": &MockService{name: "Service1"},
			"Service2": &MockService{name: "Service2"},
		}

		filtered := s.Filter(nil, true)
		assert.Equal(t, s, filtered)
	})

	// Test filtering with specific services to keep
	t.Run("Filter with specific services", func(t *testing.T) {
		s := services{
			"Service1": &MockService{name: "Service1"},
			"Service2": &MockService{name: "Service2"},
			"Service3": &MockService{name: "Service3"},
		}

		servicesToKeep := SliceFlag{"Service1", "Service3"}
		filtered := s.Filter(servicesToKeep, false)

		expectedFiltered := services{
			"Service1": s["Service1"],
			"Service3": s["Service3"],
		}

		assert.Equal(t, expectedFiltered, filtered)
	})

	// Test filtering with a service that does not exist
	t.Run("Filter with non-existent service", func(t *testing.T) {
		s := services{
			"Service1": &MockService{name: "Service1"},
			"Service2": &MockService{name: "Service2"},
		}

		servicesToKeep := SliceFlag{"Service1", "NonExistentService"}
		filtered := s.Filter(servicesToKeep, false)

		expectedFiltered := services{
			"Service1": s["Service1"],
		}

		assert.Equal(t, expectedFiltered, filtered)
	})

	// Test filtering with a nil services to keep
	t.Run("Filter with a nil services to keep", func(t *testing.T) {
		s := services{
			"Service1": &MockService{name: "Service1"},
			"Service2": &MockService{name: "Service2"},
		}

		filtered := s.Filter(nil, false)

		expectedFiltered := services{}

		assert.Equal(t, expectedFiltered, filtered)
	})
}

// TestSliceFlag tests the SliceFlag type
func TestSliceFlag(t *testing.T) {
	// Test String method
	t.Run("String method", func(t *testing.T) {
		flag := SliceFlag{"value1", "value2"}
		expectedString := "value1, value2"
		assert.Equal(t, expectedString, flag.String(), "String method result is not as expected")
	})

	// Test Set method with unique value
	t.Run("Set method with unique value", func(t *testing.T) {
		flag := SliceFlag{"value1", "value2"}
		err := flag.Set("value3")
		expectedFlag := SliceFlag{"value1", "value2", "value3"}

		assert.NoError(t, err, "Set method should not return an error")
		assert.Equal(t, expectedFlag, flag, "Set method did not append the unique value")
	})

	// Test Set method with duplicate value
	t.Run("Set method with duplicate value", func(t *testing.T) {
		flag := SliceFlag{"value1", "value2"}
		err := flag.Set("value2")

		assert.NoError(t, err, "Set method should not return an error for duplicate value")
		assert.Len(t, flag, 2, "Set method should not append the duplicate value")
	})
}
