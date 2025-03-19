# Design choices

## 1. State
We're using a stateless, disposable, replicable application. <br>
Runs well in a container; multiple instances can be spawned

## 2. Request logging
Gone with a per-request logger instance.

Pros:
 - Thread-safe by design
 - Logger can have contextual information, like correlation ID / request ID, etc
 - No need to pass app's main logger instance around

Cons:
 - More memory and CPU usage (new logger instance per request)
   - But it doesn't matter because we're not ultra high volume here. 

## 3. Database
 - We're connecting to a relational SQL-like DB via the repository and service pattern
 - This DDD separates DB implementation from business logic, and services from "controllers"
 - Assuming an RDS-style, redundant, managed DB
 - Using DTOs an models to pass data between layers, placed in "models" and "dto" packages
 - See the /database/design folder for the schema and design choices
   - https://dbdiagram.io/d/cushon-67d7694c75d75cc844497b36 for the visual DB layout
  


 

