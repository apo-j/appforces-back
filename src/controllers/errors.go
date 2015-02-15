package controllers

import (
	"fmt"
)

const (
	// Error codes
	ErrCodeNotExist      = 1
	ErrCodeAlreadyExists = 2
)

// The serializable Error structure.
type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// NewError creates an error instance with the specified code and message.
func NewError(code int, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

