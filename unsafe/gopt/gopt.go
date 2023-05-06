package gopt

type Setter interface {
	Set(string, any)
}

type Option func(Setter)

func With(name string, value any) Option {
	return func(s Setter) { s.Set(name, value) }
}

// Apply applies all options to s. To be able to return a specific type,
// it uses generics instead of interfaces.
func Apply[T Setter](s T, options ...Option) T {
	for _, option := range options {
		option(s)
	}
	return s
}
