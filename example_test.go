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

func (ServerOption) Host(host string) gopt.Option[*Server] {
	return func(s *Server) { s.host = host }
}

func (ServerOption) Port(port int) gopt.Option[*Server] {
	return func(s *Server) { s.port = port }
}

func Example() {
	server := New(
		ServerOption{}.Host("localhost"),
		ServerOption{}.Port(8080),
	)
	fmt.Printf("server: %+v\n", server)

	// Output:
	// server: &{host:localhost port:8080}
}
