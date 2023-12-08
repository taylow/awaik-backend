package handler

import (
	"time"

	monitorv1 "github.com/taylow/awaik-backend/internal/gen/proto/monitor/v1"
	"github.com/taylow/awaik-backend/services/monitor/command/domain"
)

// monitorToProto converts a domain.Monitor to a monitorv1.Monitor
func monitorToProto(monitor *domain.Monitor) *monitorv1.Monitor {
	return &monitorv1.Monitor{
		Id:             monitor.ID,
		ProjectId:      monitor.ProjectID,
		Name:           monitor.Name,
		Description:    monitor.Description,
		Status:         statusToProto(monitor.Status),
		Interval:       monitor.Interval,
		Regions:        monitor.Regions,
		Protocol:       protocolToProto(monitor.Protocol),
		ProtocolConfig: protocolConfigToProto(monitor.GetProtocolConfig()),
		CreatedAt:      monitor.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      monitor.UpdatedAt.Format(time.RFC3339),
		DeletedAt:      dateOrEmpty(monitor.DeletedAt),
	}
}

// statusToProto converts a domain.Status to a monitorv1.Status
func statusToProto(status domain.Status) monitorv1.Status {
	switch status {
	case domain.StatusPending:
		return monitorv1.Status_Pending
	case domain.StatusActive:
		return monitorv1.Status_Active
	case domain.StatusPaused:
		return monitorv1.Status_Paused
	case domain.StatusDeleted:
		return monitorv1.Status_Deleted
	default:
		return monitorv1.Status_UnknownStatus
	}
}

// protocolToProto converts a domain.Protocol to a monitorv1.Protocol
func protocolToProto(protocol domain.Protocol) monitorv1.Protocol {
	switch protocol {
	case domain.ProtocolHTTP:
		return monitorv1.Protocol_HTTP
	case domain.ProtocolICMP:
		return monitorv1.Protocol_ICMP
	case domain.ProtocolPort:
		return monitorv1.Protocol_Port
	case domain.ProtocolBrowserAutomation:
		return monitorv1.Protocol_BrowserAutomation
	default:
		return monitorv1.Protocol_UnknownProtocol
	}
}

// protocolConfigToProto converts a domain.Protocol to a monitorv1.Protocol
func protocolConfigToProto(protocol any) *monitorv1.ProtocolConfig {
	switch p := protocol.(type) {
	case *domain.HTTP:
		return &monitorv1.ProtocolConfig{
			Protocol: &monitorv1.ProtocolConfig_Http{
				Http: &monitorv1.HTTPConfig{
					Address:            p.Address,
					Method:             p.Method,
					Timeout:            p.Timeout,
					FollowRedirects:    p.FollowRedirects,
					Headers:            p.Headers,
					BodyFormat:         bodyFormatToProto(p.BodyFormat),
					Body:               p.Body,
					Proxy:              p.Proxy,
					ExpectedStatusCode: p.ExpectedStatusCode,
					BodyAssertion:      p.BodyAssertion,
				},
			},
		}
	case *domain.ICMP:
		return &monitorv1.ProtocolConfig{
			Protocol: &monitorv1.ProtocolConfig_Icmp{
				Icmp: &monitorv1.ICMPConfig{
					Address: p.Address,
				},
			},
		}
	case *domain.Port:
		return &monitorv1.ProtocolConfig{
			Protocol: &monitorv1.ProtocolConfig_Port{
				Port: &monitorv1.PortConfig{
					Address: p.Address,
					Port:    p.Port,
				},
			},
		}
	case *domain.BrowserAutomation:
		return &monitorv1.ProtocolConfig{
			Protocol: &monitorv1.ProtocolConfig_BrowserAutomation{
				BrowserAutomation: &monitorv1.BrowserAutomationConfig{
					Script: p.Script,
				},
			},
		}
	default:
		return nil
	}
}

// bodyFormatToProto converts a domain.BodyFormat to a monitorv1.BodyFormat
func bodyFormatToProto(bodyFormat domain.BodyFormat) monitorv1.BodyFormat {
	switch bodyFormat {
	case domain.JSON:
		return monitorv1.BodyFormat_JSON
	case domain.Raw:
		return monitorv1.BodyFormat_Raw
	case domain.Form:
		return monitorv1.BodyFormat_Form
	default:
		return monitorv1.BodyFormat_None
	}
}

// monitorFromProto converts a monitorv1.Monitor to a domain.Monitor
func monitorFromProto(monitor *monitorv1.Monitor) *domain.Monitor {
	m := &domain.Monitor{
		ID:          monitor.Id,
		ProjectID:   monitor.ProjectId,
		Name:        monitor.Name,
		Description: monitor.Description,
		Interval:    monitor.Interval,
		Status:      statusFromProto(monitor.Status),
		Regions:     monitor.Regions,
		Protocol:    protocolFromProto(monitor.Protocol),
		CreatedAt:   must(time.Parse(time.RFC3339, monitor.CreatedAt)),
		UpdatedAt:   must(time.Parse(time.RFC3339, monitor.UpdatedAt)),
		DeletedAt:   must(time.Parse(time.RFC3339, monitor.DeletedAt)),
	}

	config := protocolConfigFromProto(monitor.ProtocolConfig)
	config.Apply(m)

	return m
}

// statusFromProto converts a monitorv1.Status to a domain.Status
func statusFromProto(status monitorv1.Status) domain.Status {
	switch status {
	case monitorv1.Status_Pending:
		return domain.StatusPending
	case monitorv1.Status_Active:
		return domain.StatusActive
	case monitorv1.Status_Paused:
		return domain.StatusPaused
	case monitorv1.Status_Deleted:
		return domain.StatusDeleted
	default:
		return domain.StatusUnknown
	}
}

// protocolFromProto converts a monitorv1.Protocol to a domain.Protocol
func protocolFromProto(protocol monitorv1.Protocol) domain.Protocol {
	switch protocol {
	case monitorv1.Protocol_HTTP:
		return domain.ProtocolHTTP
	case monitorv1.Protocol_ICMP:
		return domain.ProtocolICMP
	case monitorv1.Protocol_Port:
		return domain.ProtocolPort
	case monitorv1.Protocol_BrowserAutomation:
		return domain.ProtocolBrowserAutomation
	default:
		return domain.ProtocolUnknown
	}
}

// protocolConfigFromProto converts a monitorv1.ProtocolConfig to a domain.ProtocolConfig
func protocolConfigFromProto(protocolConfig *monitorv1.ProtocolConfig) domain.ProtocolConfig {
	switch p := protocolConfig.Protocol.(type) {
	case *monitorv1.ProtocolConfig_Http:
		return &domain.HTTP{
			Address:            p.Http.Address,
			Method:             p.Http.Method,
			Timeout:            p.Http.Timeout,
			FollowRedirects:    p.Http.FollowRedirects,
			Headers:            p.Http.Headers,
			BodyFormat:         bodyFormatFromProto(p.Http.BodyFormat),
			Body:               p.Http.Body,
			Proxy:              p.Http.Proxy,
			ExpectedStatusCode: p.Http.ExpectedStatusCode,
			BodyAssertion:      p.Http.BodyAssertion,
		}
	case *monitorv1.ProtocolConfig_Icmp:
		return &domain.ICMP{
			Address: p.Icmp.Address,
		}
	case *monitorv1.ProtocolConfig_Port:
		return &domain.Port{
			Address: p.Port.Address,
			Port:    p.Port.Port,
		}
	case *monitorv1.ProtocolConfig_BrowserAutomation:
		return &domain.BrowserAutomation{
			Script: p.BrowserAutomation.Script,
		}
	default:
		return nil
	}
}

// bodyFormatFromProto converts a monitorv1.BodyFormat to a domain.BodyFormat
func bodyFormatFromProto(bodyFormat monitorv1.BodyFormat) domain.BodyFormat {
	switch bodyFormat {
	case monitorv1.BodyFormat_JSON:
		return domain.JSON
	case monitorv1.BodyFormat_Raw:
		return domain.Raw
	case monitorv1.BodyFormat_Form:
		return domain.Form
	default:
		return domain.None
	}
}

// must checks if err is nil, if not it panics
func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

// dateOrEmpty returns the date formatted as RFC3339 or an empty string
func dateOrEmpty(date time.Time) string {
	if date.IsZero() {
		return ""
	}

	return date.Format(time.RFC3339)
}
