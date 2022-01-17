package errors

type errorCode struct {
	code string
}

func Define(code string) *errorCode {
	if code == "" {
		panic("empty error code")
	}

	return &errorCode{
		code: code,
	}
}
