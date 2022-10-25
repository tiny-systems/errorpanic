# errorpanic

errorpanic It is a simple golang wrapper for functions which may panic.

* Synchronous call.
* For functions which may return errors.
* If panic happens, it will be handled as an error with the help of the `recover` function
* Error typed panics are supported `panic(err)`

## Installation

Standard `go get`:

```
go get github.com/tinysystemsio/errorpanic
```

## Usage & Example

```go
err := errorpanic.Wrap(someFuncError)
```
Please check [examples](examples/README.md)
