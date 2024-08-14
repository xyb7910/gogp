package option

type Option[T any] func(t *T)

// WithOption add options to T
func WithOption[T any](t *T, opts ...Option[T]) {
	for _, opt := range opts {
		opt(t)
	}
}

type OptionErr[T any] func(t *T) error

// WithOptionErr add options to T and return error
func WithOptionErr[T any](t *T, opts ...OptionErr[T]) error {
	for _, opt := range opts {
		if err := opt(t); err != nil {
			return err
		}
	}
	return nil
}
