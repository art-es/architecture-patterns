# Layered Pattern

Base layers:
- Presentation Layer
- Application Layer
- Business Logic Layer
- Data Access Layer (Persistence & Database Layers)
Additional
- Infrastructure Layer (contains adapters, bridges, etc. of external libraries)

Among the advantages of the layered pattern is that a lower layer can be used by different higher layers.

## Presentation Layer

The presentation layer is responsible for handling the user interface.

_`./presentation/...` packages_

## Application Layer

The application layer sits between the presentation layer and the business layer. On the one hand, it provides an abstraction so that the presentation layer doesn’t need to know the business layer. 

_`./application/...` packages_

## Business Logic Layer

In the business logic layer contains the base logic how to handle the data.

The business layer doesn’t need to be concerned about how to format customer data for display.

In some cases, the business layer and persistence layer are combined into a single business layer.

_`./domain/...` packages_

## Data Access Layer

A data access layer is a layer of an programming entity which provides simplified access to data stored in persistent storage of some kind, such as an entity-relational database.

_`./persistence/...` packages_
