This is just a little bit of PHP goodness to 
make error handling more palatable at scale, in a commercial Go app.

PHP does Exception and error handling exceptionally well, excuse pun.

This is my quick port of PHP's Exception paradigm to Go, with a bit of extra sugar.

It supports:
- Exception chaining
- Arbitrary data
- Arbitrary error codes
- Backtraces / stack traces (requires compiler support) with graceful fallback
- Pretty-printing of exception chains, traces, etc
- A catalogue of common, user-defined formal exception types
- Emulation for "Throw" that works elegantly with Go. 

The idea is not to break Go's idiomatic error handling, 
but to provide a more expressive way to handle nuanced and contextual errors in 
a concurrent Go app at scale. 

Will open-source this soon, still needs some work.

