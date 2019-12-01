<<<<<<< HEAD

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
=======
# clean-arch
>>>>>>> 1cc1692d6953203fd7ee3914c9a52400cdcc4bd0
