# MonitoringService

The `MonitoringService` is responsible for monitoring heartbeats from internal services and reacting to those that are missing.

## Events

### Produces

- [ServiceDownEvent](../../../docs/events.md#taskrecoveredevent)

### Consumes

- [HeartbeatEvent](../../../docs/events.md#heartbeatevent)
