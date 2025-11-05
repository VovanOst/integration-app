package domain

import "fmt"

// CustomError — пользовательская ошибка
type CustomError struct {
	message string
}

// NewError создает новую ошибку
func NewError(message string) error {
	return &CustomError{message: message}
}

// NewErrorf создает ошибку с форматированием
func NewErrorf(format string, args ...interface{}) error {
	return NewError(fmt.Sprintf(format, args...))
}

func (e *CustomError) Error() string {
	return e.message
}

// Predefined errors
var (
	ErrNotFound       = NewError("not found")
	ErrInvalidInput   = NewError("invalid input")
	ErrAlreadyExists  = NewError("already exists")
	ErrUnauthorized   = NewError("unauthorized")
	ErrInternalServer = NewError("internal server error")
)
