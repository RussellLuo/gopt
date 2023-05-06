package gopt_test

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

// ServerOption is a public singleton used to specify options.
var ServerOption serverOption

// severOption is unexported and holds all option definitions.
type serverOption struct{}

func (_ serverOption) WithHost(host string) gopt.Option[*Server] {
	return func(s *Server) { s.host = host }
}

func (_ serverOption) WithPort(port int) gopt.Option[*Server] {
	return func(s *Server) { s.port = port }
}

func Example() {
	server := New(
		ServerOption.WithHost("localhost"),
		ServerOption.WithPort(8080),
	)
	fmt.Printf("server: %+v\n", server)

	// Output:
	// server: &{host:localhost port:8080}
}
