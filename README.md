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
	Host string
	Port int
}

func (s *Server) Set(name string, value any) { gopt.ReflectSet(s, name, value) }

func New(options ...gopt.Option) *Server {
	return gopt.Apply(new(Server), options...)
}

func main() {
	server := New(
		gopt.With("Host", "localhost"),
		gopt.With("Port", 8080),
	)
	fmt.Printf("server: %+v\n", server)

	// Output:
	// server: &{Host:localhost Port:8080}
}
```


## FAQ

1. Why might I want to use this tiny library?

    Traditional Functional Options (see [this][2] and [this][3]) will define many exported functions, which is likely to cause naming conflicts.

2. What if I don't want to expose gopt's API to my own library users?

   One solution is to define your own API like this:

    ```go
    var WithOption = gopt.With
    ```

   Then your library users can use it as follows:

    ```go
    server := New(
    	WithOption("Host", "localhost"),
    	WithOption("Port", 8080),
    )
    ```

3. What if the struct fields are unexported?

    Write the `Set()` method by hand. For example:

    ```go
    type Server struct {
    	host string
    	port int
    }
    
    func (s *Server) Set(name string, value any) {
    	switch name {
    	case "host":
    		s.host = value.(string)
    	case "port":
    		s.port = value.(int)
    	default:
    		panic("no field named " + name)
    	}
    }
    ```
   
    To make the errors more human-readable, if encountered during programming, you can also leverage gopt's helper utilities:

    ```go
    type Server struct {
    	host string
    	port int
    }
    
    func (s *Server) Set(name string, value any) {
    	switch name {
    	case "host":
    		gopt.F(s, name, &s.host).MustSet(value)
    	case "port":
    		gopt.F(s, name, &s.port).MustSet(value)
    	default:
    		panic(gopt.ErrNotFound(s, name))
    	}
    }
    ```

4. What if I don't want to use the slow `gopt.ReflectSet()`?

    Write the `Set()` method by hand.


## License

[MIT](LICENSE)


[1]: https://pkg.go.dev/github.com/RussellLuo/gopt
[2]: https://golang.cafe/blog/golang-functional-options-pattern.html
[3]: https://github.com/tmrts/go-patterns/blob/master/idiom/functional-options.md
