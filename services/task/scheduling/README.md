# TaskSchedulingService

The Task Scheduling Service is responsible for scheduling tasks to a pool of workers ([TaskExecutionService](../execution/README.md))

## Events

### Produces

- [HeartbeatEvent](../../../docs/events.md#heartbeatevent)
- [TaskScheduledEvent](../../../docs/events.md#taskscheduledevent)
- [TaskRescheduledEvent](../../../docs/events.md#taskrescheduledevent)

### Consumes

- [TaskCreatedEvent](../../../docs/events.md#taskcreatedevent)
- [TaskModifiedEvent](../../../docs/events.md#taskmodifiedevent)
- [TaskDeletedEvent](../../../docs/events.md#taskdeletedevent)
- [TaskRecoveredEvent](../../../docs/events.md#taskrecoveredevent)
