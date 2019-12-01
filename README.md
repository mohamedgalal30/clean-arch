# What is this?
   - this is an example trying to apply Uncle Bob’s concept of [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) by demonstrating how its principles could be applied to an actual Go application.
# Architecture
The architecture will be separated into four layers: domain, use cases, interfaces and infrastructure.
# The Pattern
| Infrastructure       | Interfaces           | Use Cases            | Domain               |
|----------------------|----------------------|----------------------|----------------------|
| application-agnostic | application-specific | application-specific | application-agnostic |
| business-agnostic    | business-agnostic    | business-specific    | business-specific    |


For every part in every layer, there are three questions of interest: where is it used, where is its interface, where is its implementation?

If we look at the OrderRepository, the answers are as follows: it’s used by the use cases layer, its interface belongs to the domain layer, and its implementation belongs to the interfaces layer.
The Add method of the Order entity, on the other hand, is used by the uses cases layer, too, and also, its interface belongs to the domain layer. But, its implementation belongs there as well, because it doesn’t need anything outside the domain layer itself.

declare interface in the place that you will use it. not implement it.
use cases “…orchestrate the flow of data to and from the entities, and direct those entities to use their enterprise wide business rules to achieve the goals of the use case.”
injection is used to handle dependencies.,it could be mocked in the unit tests, making the web service handler testable in isolation
stuff like connecting to the database daemon through the network, deciding to use a slave for reads and the master for writes, handling timeouts, and so forth, are infrastructural issues.
the repository will have to sort out the details of putting together a single entity from multiple tables, but the clients of the repositories won’t be concerned.

consider updating jira when you start working on a task
