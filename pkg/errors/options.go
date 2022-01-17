package errors

type Option func(e *err)

func WithCause(cause error) Option {
	return func(e *err) {
		e.cause = cause
	}
}

func WithMetadata(k string, v interface{}) Option {
	return func(e *err) {
		if e.metadata == nil {
			e.metadata = make(map[string]interface{})
		}

		e.metadata[k] = v
	}
}
