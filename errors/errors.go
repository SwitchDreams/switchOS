package errors

type IError interface {
	Error() string
}

type OSError struct {
	message string
}

func (o OSError) Message() string {
	return o.message
}

func (o OSError) Error() string {
	return o.message
}

func NewOSError(msg string) *OSError {
	return &OSError{
		message: msg,
	}
}

func WrapError(err error, msg string) *OSError {
	if err == nil {
		return NewOSError(msg)
	}

	if e, ok := err.(*OSError); ok {
		return NewOSError(
			msg + ": " + e.Message(),
		)
	}
	e := msg + ": " + err.Error()
	return NewOSError(e)
}
