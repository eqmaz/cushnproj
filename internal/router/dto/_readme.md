### DTOs for the router module

This package is a collection of formal response shapes for the 
various API responses. These are used to ensure that the API
responses are consistent and predictable.

It also makes testing easier as we can assert that the response
shape is as expected.

The API responses are therefore always correct by design as
these DTOs act like a kind of "contract". 

Ideally all API handlers should return a DTO object result
rather than a raw object or arbitrary data structure.

Being explicit about the response shape also reduces the
risk of data leakage in the API responses.

