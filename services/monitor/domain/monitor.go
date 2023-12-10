package domain

import (
	"time"
)

// Monitor is a monitor
type Monitor struct {
	ID          string
	ProjectID   string
	Name        string
	Description string
	Interval    string
	Status      Status
	Regions     []string
	Protocol    Protocol
	HTTP        *HTTP
	ICMP        *ICMP
	Port        *Port
	Browser     *BrowserAutomation
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

// GetProtocol returns the protocol
func (m *Monitor) GetProtocolConfig() any {
	switch m.Protocol {
	case ProtocolHTTP:
		return m.HTTP
	case ProtocolICMP:
		return m.ICMP
	case ProtocolPort:
		return m.Port
	case ProtocolBrowserAutomation:
		return m.Browser
	}

	return nil
}

// ConfigApplier applies a config to a monitor
type ProtocolConfig interface {
	Apply(*Monitor)
}

// Status is a status
type Status int

const (
	// StatusUnknown is an unknown status
	StatusUnknown Status = iota
	// StatusPending is a pending status
	StatusPending
	// StatusActive is an active status
	StatusActive
	// StatusPaused is a paused status
	StatusPaused
	// StatusDeleted is a deleted status
	StatusDeleted
)

// Protocol is a protocol
type Protocol int

const (
	ProtocolUnknown Protocol = iota
	ProtocolHTTP
	ProtocolICMP
	ProtocolPort
	ProtocolBrowserAutomation
)

type Protocols interface {
	HTTP | ICMP | Port | BrowserAutomation
}

// HTTP is an HTTP protocol
type HTTP struct {
	Address            string
	Method             string
	Timeout            string
	FollowRedirects    bool
	Headers            map[string]string
	BodyFormat         BodyFormat
	Body               string
	Proxy              string
	ExpectedStatusCode string
	BodyAssertion      string
}

// Apply applies the HTTP config to a monitor
func (h *HTTP) Apply(m *Monitor) {
	m.Protocol = ProtocolHTTP
	m.HTTP = h
	m.ICMP = nil
	m.Port = nil
	m.Browser = nil
}

// BodyFormat is a body format
type BodyFormat int

const (
	// None is a no body format
	None BodyFormat = iota
	// Raw is a raw body format
	Raw
	// JSON is a JSON body format
	JSON
	// Form is a form body format
	Form
)

// ICMP is an ICMP protocol
type ICMP struct {
	Address string
}

// Apply applies the ICMP config to a monitor
func (i *ICMP) Apply(m *Monitor) {
	m.Protocol = ProtocolICMP
	m.HTTP = nil
	m.ICMP = i
	m.Port = nil
	m.Browser = nil
}

// Port is a port protocol
type Port struct {
	Address string
	Port    int32
}

// Apply applies the port config to a monitor
func (p *Port) Apply(m *Monitor) {
	m.Protocol = ProtocolPort
	m.HTTP = nil
	m.ICMP = nil
	m.Port = p
	m.Browser = nil
}

// BrowserAutomation is a browser automation protocol
type BrowserAutomation struct {
	Address string
	Script  string
}

// Apply applies the browser automation config to a monitor
func (b *BrowserAutomation) Apply(m *Monitor) {
	m.Protocol = ProtocolBrowserAutomation
	m.HTTP = nil
	m.ICMP = nil
	m.Port = nil
	m.Browser = b
}
