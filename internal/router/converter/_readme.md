### Converter package

Oftentimes we need to convert and transform input or output data back and forth between
database models and internal structures, to structures that are suitable for the client.

This package contains the converters that are used to convert data between these different
structures. It's a cleaner approach, which moves the burden of conversion from the service
layer to the converter layer. Another benefit is that we can reuse converters when needed.

