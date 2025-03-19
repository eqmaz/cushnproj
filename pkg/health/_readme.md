# A basic health package

This package that spins up a health listener on one or more of:
 - an HTTP server (using Fiber)
 - on GRPC server

The patten here is passive (reactive) health checking. The health check is only
triggered when a request is made to the health endpoint. 

There are two methods:
- ping: returns a "pong" or 200 OK if the service is healthy
- health: Does a full health check by running a lambda function and returns a 200 OK if the service is healthy, otherwise details the problems

