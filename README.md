# Account API Client Library

I tried to create a client library that is suitable for projects with Hexagonal Architecture.Every operation has its own package, interface, and adapters. Adapter layers and REST clients are separated and there is a DI process in between them. The components are highly reusable for other purposes and independently configurable.

SOLID principles are taken into account during the design process. Methods have a single responsibility, interfaces are segregated, dependency injection mechanism is existing.

Method input parameters and return variables are designed to increase memory management performance. The purpose is to decrease heap and also increase stack usage.

The library contains both unit and integration tests. A mock rest API is used to simulate account API for integration tests.

## Tests

Unit and Integration Test run command

    make test
