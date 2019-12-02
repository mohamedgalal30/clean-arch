# What is this?
 - This is an example to illustrate how to apply Uncle Bob’s concept of [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) by demonstrating how its principles could be applied to an actual Go application.
# Architecture
- The architecture will be separated into four layers: domain entities, use cases, interfaces and infrastructure.
- Declare interface in the place that you will use it. not in the place that you will implement it.
- Injection is used to handle dependencies. it could be mocked in the unit tests, making any layer testable in isolation.
# What to put where, and why?
The entity and use cases layers together form the core of our application, 
representing the realities of the business we are operating in. 
Everything else is implementation details that are not related to the nature of our business.

## Entity
This layer, will have Data Structure and business rules.

## Use Cases
- Use cases orchestrate the flow of data to and from the entities, and direct those entities to use their enterprise wide business rules to achieve the goals of the use case.
This layer will act as the business process handler. Any process will be handled here. 
- This layer will decide, which repository will use. And have responsibility to provide data to serve into delivery. 
- Process the data doing calculation or anything will done here.
- Usecase layer will accept sanitized input from Delivery layer, then process the input and pass it to repository layer to store it in DB , or Fetching from another micro service ,etc.
- This layer will depends to entity Layer.

## Interfaces
### Repository
- Repository will store any Database handler. Querying, or Creating/ Inserting into any database will stored here. 
- This layer will act for CRUD to database only. No business process happen here. Only plain function to Database.
- This layer also have responsibility to choose what DB will used in Application. Could be redis, MongoDB, Postgres whatever, will decided here.
- This layer will control the input, and give it directly to ORM services.
- Calling microservices, will be handled here. Create HTTP Request to other services, and sanitize the data. 
- Handle all data input - output no specific logic happen.
- The repository will have to sort out the details of putting together a single entity from multiple tables, but the clients of the repositories won’t be concerned.
- This layer will depends to Connected DB , or other microservices.

### Delivery
- This layer will act as the presenter. Decide how the data will presented. 
Could be as REST API, CLI or gRPC whatever the delivery type. 
- This layer also will accept the input from user. Sanitize the input and sent it to Usecase layer.
- Client will call the resource endpoint over network, and the Delivery layer will get the input or request, 
and sent it to Usecase Layer.
- This layer will depends to Usecase Layer.

## Infrastructure
This layer will have general-purpose code that can be used in any other application, like: 
- Creating web server instance.
- Connecting to the database daemon through the network, deciding to use a slave for reads and the master for writes, handling timeouts.
- Connecting to db or another microservice.
- An Adapter to any package or vendor will be here.
# The Pattern
| Infrastructure       | Interfaces           | Use Cases            | Domain               |
|----------------------|----------------------|----------------------|----------------------|
| application-agnostic | application-specific | application-specific | application-agnostic |
| business-agnostic    | business-agnostic    | business-specific    | business-specific    |

For every part in every layer, there are three questions of interest: where is it used, where is its interface, where is its implementation?

If we look at the TicketDbRepo, the answers are as follows: 
- it’s used by the use cases layer, 
- its interface belongs to the domain layer, 
- and its implementation belongs to the interfaces/repository layer.

## TODO
- transformers between layers
