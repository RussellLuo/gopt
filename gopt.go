package gopt

// Option sets an optional parameter for x.
type Option[T any] func(x T)

// Apply applies all options to x.
func Apply[T any](x T, options ...Option[T]) T {
	for _, option := range options {
		option(x)
	}
	return x
}
