# Events

Awaik is built on an event-driven architecture. Below is a list of all events, their contents, and how they are used within the system.

## Health

### HeartbeatEvent

Sent periodically by all services to indicate they are alive

Contains:

- `service_id` - UUID of the specific service sending the heartbeat
- `timestamp` - Timestamp of when the heartbeat was produced

### ServiceDownEvent

Sent when a service is detected as no longer alive

Contains:

- `service_id` - UUID of the specific service sending the heartbeat
- `timestamp` - Timestamp of when the service was detected as down

### TaskRecoveredEvent

Sent when a task has been recovered

Contains:

- `task_id` - UUID of the specific service sending the heartbeat
- `timestamp` - Timestamp of when the task was recovered

## Task

### TaskCreatedEvent

Sent when a new `task` is created

Contains:

- `task_id` - UUID of the new task
- `cron` - CRON expression for the scheduling of the task

### TaskModifiedEvent

Sent when a `task` is modified

Contains:

- `task_id` - UUID of the modified task
- `cron` - Updated CRON expression for the scheduling of the task

### TaskDeletedEvent

Sent when a `task` is deleted

Contains:

- `task_id` - UUID of the deleted task

### TaskPausedEvent

Sent when a running `task` is paused

Contains:

- `task_id` - UUID of the paused task

### TaskResumedEvent

Sent when a paused `task` is resumed

Contains:

- `task_id` - UUID of the resumed task
- `cron` - Updated CRON expression for the scheduling of the task

### TaskScheduledEvent

Sent when a `task` is scheduled

Contains:

- `task_id` - UUID of the scheduled task
- `scheduled_at` - Timestamp of when the task was scheduled
- `scheduled_for` - Timestamp of when the task was scheduled for

### TaskRescheduledEvent

Sent when a `task` is rescheduled

Contains:

- `task_id` - UUID of the rescheduled task
- `scheduled_at` - Timestamp of when the task was scheduled
- `scheduled_for` - Timestamp of when the task was scheduled for

### TaskExecutedEvent

Sent when a `task` is executed

Contains:

- `task_id` - UUID of the executed task
- `result` - TBD

### ExecuteTaskEvent

Sent when a `task` is to be executed

Contains:

- `task_id` - UUID of the task to be executed

## Notification

### NotificationSentEvent

Sent when a `notification` has been sent

Contains:

- `notification_id` - UUID of the notification task
