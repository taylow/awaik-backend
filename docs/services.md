# Services

Awaik uses a microservice architecture, and is made up of a handful of services, each responsible for specific domains and functionality.

The choice for microservices in a project of this size was:

- Ease of scaling - as services are event-driven, adding new functionality is as easy as adding a new consumer microservice to the system, that consumes and produces the necessary event for the desired result
- Ease of contribution/ownership - contributors can maintain small portions of code and limit changes to specific features or domains
- Organisation and structure - separating groups of functionality as microservices allows for a clean project structure, and ease of maintenance
- Modularity - being able to work on each part of the system in isolation will result in concise task lists, engineering efforts, and problem domains
- Testability - each service can be tested in isolation without the need of the whole system to be under test (small SUT)

## Microservices

### Health

- [HealthMonitoringService](../services/health/monitoring/README.md)
- [HealthRecoveryService](../services/health/recovery/README.md)

### Notification

- [NotificationService](../services/notification/README.md)

### Task

- [TaskEditingService](../services/task/editing/README.md)
- [TaskExecutionService](../services/task/execution/README.md)
- [TaskSchedulingService](../services/task/scheduling/README.md)

## 3rd Party Services

### Apache Kafka

Apache Kafka is an open source distributed event streaming platform used by thousands of companies for high-performance data pipelines, streaming analytics, data integration, and mission-critical applications.

As Awaik is build on an event-driven architecture, Kafka is used as the backbone for all event-related communication. Consumer groups allow for simplified scaling, as messages can be load balanced across a set of workers from a single subject. The majority of service-to-service communication is done via events, in a reactive manner rather than a direct connection.

### Apache Cassandra

Apache Cassandra is an open source NoSQL distributed database trusted by thousands of companies for scalability and high availability without compromising performance. Linear scalability and proven fault-tolerance on commodity hardware or cloud infrastructure make it the perfect platform for mission-critical data.

Cassandra is used to store most of the data that powers Awaik. From user data, to service monitoring configuration data, Cassandra provides a simple NoSQL means to store mass amounts of data, way beyond the limits that Awaik will reach.

### HashiCorp Vault

Vault is a tool for securely accessing secrets. A secret is anything that you want to tightly control access to, such as API keys, passwords, certificates, and more. Vault provides a unified interface to any secret, while providing tight access control and recording a detailed audit log.

Vault is used to store Awaik user secrets - any sensitive information required in order for a request to be successfully. Providing a secret store allows users to monitor protected endpoints, and allows for future expandability of advance service monitoring.
