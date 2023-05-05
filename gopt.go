package gopt

type Setter interface {
	Set(string, any)
}

type Option[T Setter] func(x T)

func With[T Setter](name string, value any) Option[T] {
	return func(x T) { x.Set(name, value) }
}

func Apply[T Setter](x T, options ...Option[T]) T {
	for _, option := range options {
		option(x)
	}
	return x
}
