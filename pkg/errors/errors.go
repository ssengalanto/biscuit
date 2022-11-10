package errors

import (
	"errors"
)

// Errors used in the application.
var (
	ErrInvalid           = errors.New("invalid")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
	ErrNotFound          = errors.New("not found")
	ErrInternal          = errors.New("internal server error")
	ErrUnknown           = errors.New("unknown error")
	ErrTemporaryDisabled = errors.New("temporary disabled")
	ErrTimeout           = errors.New("timeout")
)

type Error struct {
	err error
}

// New returns a new error.
func New(message string) *Error {
	return newErr(errors.New(message))
}

// Wrap returns an error that wraps the target error.
func Wrap(err error) *Error {
	return newErr(err)
}

// Unwrap return an unwrap error.
func (e *Error) Unwrap() error {
	return e.err
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.err.Error()
}

func newErr(err error) *Error {
	if err == nil {
		err = ErrUnknown
	}

	return &Error{
		err: err,
	}
}
