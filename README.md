## Hypothetical Retail ISA Account service

## Overview
The company has a hypothetical relational database where details of users, employers, accounts, products and investments are kept.

### Retail vs employer access
This service provide an embryonic REST API for retail customers. The assumption is that employers have a separate API that links to the same database when required. This satisfies the condition that the retail and employer based business is separate.

### Quick start:
 - Ensure Docker is installed (required for database)
 - Ensure prerequisite tooling is installed (see /docs/prerequisites.md)
 - See `config/config.json` for db and http settings
 - ``make run`` (starts db container, runs migrations, starts the app) 
 - ``make help`` (for more commands)

### Documentation
See the /docs/ folder for details on:
 - assumptions
 - design choices
 - business logic
 - framework design
 - prerequisites (for setting up the project)

### Endpoints
```  
 - GET  /isa-funds                  - list of available funds
 - GET  /user/product/available     - list of available products for a specific user
 - GET  /user/account/balances      - list of account balances for a specific user
 - POST /user/account/check-deposit - check if a deposit is allowed into a given account
 - POST /user/account/open-retail   - open a new retail account for a user
```

Opening accounts for an ISA product supports multiple funds being attached to the account in the future,
by accepting an array of funds and weightings; **however** the validator currently only allows a single fund, 
to satisfy the brief.


### Code generation and automation
- API boilerplate is generated from openapi spec
- Genesis database migration with seeds is generated (using genesis.sh) 
- DB models are then generated from the database schema (using sqlc)
- API documentation (HTML and .MD versions) is generated from the openapi spec
  - They live in /docs/api

### Known issues / what's missing
- Database mocking for repository unit tests
- Repository mocking for service unit tests
- Service and Router mocking for endpoint unit tests
- A full integration test suite in /integration_tests
- A full docker-compose setup for production


