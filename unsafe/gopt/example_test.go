package gopt_test

import (
	"fmt"

	"github.com/RussellLuo/unsafe/gopt"
)

type Server struct {
	Host string
	Port int
}

func (s *Server) Set(name string, value any) { gopt.ReflectSet(s, name, value) }

func New(options ...gopt.Option) *Server {
	return gopt.Apply(new(Server), options...)
}

func Example() {
	server := New(
		gopt.With("Host", "localhost"),
		gopt.With("Port", 8080),
	)
	fmt.Printf("server: %+v\n", server)

	// Output:
	// server: &{Host:localhost Port:8080}
}
