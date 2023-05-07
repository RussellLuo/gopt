# gopt

[![Go Reference](https://pkg.go.dev/badge/github.com/RussellLuo/gopt/vulndb.svg)][1]

Generic Functional Options for Go.


## Installation


```bash
$ go get -u github.com/RussellLuo/gopt
```


## Quick Start

```go
package main

import (
	"fmt"

	"github.com/RussellLuo/gopt"
)

type Server struct {
	host string
	port int
}

func New(options ...gopt.Option[*Server]) *Server {
	return gopt.Apply(new(Server), options...)
}

// ServerOption holds all option factories for Server.
type ServerOption struct{}

func (ServerOption) WithHost(host string) gopt.Option[*Server] {
	return func(s *Server) { s.host = host }
}

func (ServerOption) WithPort(port int) gopt.Option[*Server] {
	return func(s *Server) { s.port = port }
}

func main() {
	server := New(
		ServerOption{}.WithHost("localhost"),
		ServerOption{}.WithPort(8080),
	)
	fmt.Printf("server: %+v\n", server)

	// Output:
	// server: &{host:localhost port:8080}
}
```


## FAQ

1. Why might I want to use this tiny library?

    Traditional Functional Options will define many exported functions, which is likely to cause naming conflicts.

    Reference articles:
    - [Golang Functional Options Pattern][2]
    - Go Patterns' [Functional Options][3]

    Reference projects:
    - gRPC-Go's [DialOption][4]
    - Go kit's [ServerOption][5]

2. Why is `ServerOption{}` used in the example?

    `ServerOption` is a named alias of [the empty struct][6], whose instances are fungible and consume no space.


## Unsafe gopt

[Unsafe gopt](unsafe/gopt) is an early implementation, which is deprecated and for reference only.


## License

[MIT](LICENSE)


[1]: https://pkg.go.dev/github.com/RussellLuo/gopt
[2]: https://golang.cafe/blog/golang-functional-options-pattern.html
[3]: https://github.com/tmrts/go-patterns/blob/master/idiom/functional-options.md
[4]: https://github.com/grpc/grpc-go/blob/v1.55.0/dialoptions.go#L83
[5]: https://github.com/go-kit/kit/blob/v0.12.0/transport/http/server.go#L47
[6]: https://dave.cheney.net/2014/03/25/the-empty-struct
