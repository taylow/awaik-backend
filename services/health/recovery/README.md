# RecoveryService

The `HealthRecoveryService` is responsible for recovering tasks that are lost in failed `TaskSchedulerService` instances.

## Events

### Produces

- [HeartbeatEvent](../../../docs/events.md#heartbeatevent)
- [TaskRecoveredEvent](../../../docs/events.md#taskrecoveredevent)

### Consumes

- [ServiceDownEvent](../../../docs/events.md#servicedownevent)
