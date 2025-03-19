# Assumptions 
### (for the purpose of this excersie)

1. **Assuming a gateway in front**, that handles:
   - Authentication (all requests are authenticated with some JWT)
   - Rate limiting (all requests are within the rate limit)
   - Ddos protection (all requests are not malicious)
   - 

2. **Passive (reactive) health** pattern
   - We assume there's a central health monitor that "pings" periodically
   - The /health endpoint is on a different server/port than the main service
   - Another approach could be to actively push to a central monitor on a periodic timer
   - 
   
3. **Single application** (for hypothetical example)
   - In the real world we'd have segregated API services (apps) for retail and employer domains
     - They'd also likely support different authentication mechanisms
   - This project simply demonstrates separation of concerns through business logic, code organisation and database tables
   - 

4. **Single Relational DB** (for hypothetical example)
   - This project demonstrates a single database with multiple tables for simplicity
   - In the real world we'd have separate databases relational and time-series problems
   - A real-world scenario may have DB access layers as separate services, with buffering, caching, reconciliation, auditing and so on.
   - Real world financial data may have some level of pooled multi-tenancy via a tenant ID as a partition key, and similar strategies
   - 

5. **Joint accounts not supported**
   - We assume that an account can only have one user (account holder)
   - Currently joint ISAs are not allowed, but laws can change, and products can evolve
   - In the real world we may design the db to allow for multiple account holders (in case policy changes)
   - 

6. **Simplified models**
   - In real life ISA accounts would have account numbers, sort codes, IBANs, UUIDs, and so on (for simplicity here we're using a primary key ID)
   - You'd never normally expose database IDs to the outside world
   - 


