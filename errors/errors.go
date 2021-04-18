package errors

// IOSError implements the error interface, so that functions who return this type
// can return both the standard and custom errors
type IOSError interface {
	Error() string
}

// OSError represents the applications custom error type
type OSError struct {
	message string
}

// Message returns the error message
func (o OSError) Message() string {
	return o.message
}

// Error implements IOSError interface
func (o OSError) Error() string {
	return o.message
}

// NewOSError creates a new OSError instance
func NewOSError(msg string) *OSError {
	return &OSError{
		message: msg,
	}
}

// WrapError concatenates two errors, custom or standard
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
