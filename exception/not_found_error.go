package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundError(err error) NotFoundError {
	return NotFoundError{err.Error()}
}
