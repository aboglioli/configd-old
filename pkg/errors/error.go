package errors

import (
	"fmt"
)

type err struct {
	code     *errorCode
	message  string
	cause    error
	metadata map[string]interface{}
}

func (ec *errorCode) New(message string, opts ...Option) *err {
	e := &err{
		code:    ec,
		message: message,
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

func (e *err) Error() string {
	return fmt.Sprintf("(%s) %s", e.code.code, e.message)
}
