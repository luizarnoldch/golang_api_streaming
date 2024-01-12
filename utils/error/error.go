package error

import (
	"errors"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) ToString() string {
	return e.Message
}

func (e *Error) ToError() error {
	return errors.New(e.ToString())
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func NewNotFoundError(message string) *Error {
	return &Error{
		Code: http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError(message string) *Error {
	return &Error{
		Code: http.StatusInternalServerError,
		Message: message,
	}
}

func NewValidationError(message string) *Error {
	return &Error{
		Code: http.StatusUnprocessableEntity,
		Message: message,
	}
}