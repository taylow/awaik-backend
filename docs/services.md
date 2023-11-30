# Services

Awaik uses a microservice architecture, and is made up of a handful of services, each responsible for specific domains and functionality.

The choice for microservices in a project of this size was:

- Ease of scaling - as services are event-driven, adding new functionality is as easy as adding a new consumer microservice to the system, that consumes and produces the necessary event for the desired result
- Ease of contribution/ownership - contributors can maintain small portions of code and limit changes to specific features or domains
- Organisation and structure - separating groups of functionality as microservices allows for a clean project structure, and ease of maintenance
- 


## Services

### Health

- [HealthMonitoringService](../services/health/monitoring/README.md)
- [HealthRecoveryService](../services/health/recovery/README.md)

### Notification

- [NotificationService](../services/notification/README.md)

### Task

- [TaskEditingService](../services/task/editing/README.md)
- [TaskExecutionService](../services/task/execution/README.md)
- [TaskSchedulingService](../services/task/scheduling/README.md)
