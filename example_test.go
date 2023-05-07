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

// ServerOption holds all option factories for Server.
type ServerOption struct{}

func (ServerOption) WithHost(host string) gopt.Option[*Server] {
	return func(s *Server) { s.host = host }
}

func (ServerOption) WithPort(port int) gopt.Option[*Server] {
	return func(s *Server) { s.port = port }
}

func Example() {
	server := New(
		ServerOption{}.WithHost("localhost"),
		ServerOption{}.WithPort(8080),
	)
	fmt.Printf("server: %+v\n", server)

	// Output:
	// server: &{host:localhost port:8080}
}
