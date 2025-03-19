# About this app framework

## Introduction
This is my simple app framework designed to make it conformable to build a small-ish app with a REST or GRPC server. 
The emphasis is on scalability, speed of implementation and ease of use. It follows a lot of Go ecosystem standards
but introduces a few of my own ideas as well. 

The idea is to make life easy when making design choices around REST frameworks, logging, error handling and configuration
so being able to swap out implementations without affecting any business logic. 

## Patterns
- Application container for encapsulating dependencies
- Dependency injection of DB into repositories; of repositories into services
- Using the Single Responsibility Principle to keep code clean and maintainable
- Keeping functions small to keep code readable, maintainable and testable
- Using interfaces to allow for easier mocking and testing
- Service layer to encapsulate business logic
- Endpoint handlers to encapsulate REST logic
  - They reach out to validator layer, and convertor layer for formatting and validation
- Strictly typed models for use with database repositories (these are auto generated)
- Strict formal DTOs for both input (requests) and output (responses)
- Allowing multiple loggers, for context-based logging
- Config is a singleton, and is loaded at app start
- Decoupling of router implementation from the app
- OpenApi spec for API documentation and code generation
- integration testable, work in progress
- .devcontainer compatibility
- Auto build number version bumping (using jq) on ``make build``

## Considerations

### App Container (/internal/application)
There's a bare-bones "app container" to encapsulate the app's main dependencies.
It can be extended in any way.

### Local Packages (/pkg)
- config: 
  - Config loader, with defaults. Supports JSON files and environment variables.
- console: 
  - Lightweight console output helpers, with colors.
- database:
  - Small DB wrapper around sqlx, allowing to swap out the DB implementation easily.
- e: 
  - Error handling, inspired by PHP's Exceptions.
- health: 
  - Health check server, with REST and GRPC endpoints.
- logger: 
  - Logging, with structured logs.
- optional: 
  - Optional values, inspired by Rust's Option.
- util: 
  - Miscellaneous utility functions

### Router
- Using Fiber for now
- The REST framework can be easily swapped out with minimal effort, fiber, gin etc
- We assume some "correlation id" or "request id" is injected into each request by the gateway and is passed in the header of REST requests. This is used to track the request through the system. It's injected into all logs within the request lifetime.




