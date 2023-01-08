package errors

import (
	"errors"
	"net/http"

	apperror "github.com/ssengalanto/hex/pkg/errors"
)

type HTTPError struct {
	Error Err `json:"error"`
}

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewHTTPError returns a new HTTP error which accepts an error parameter that will be mapped
// to the corresponding application error.
func NewHTTPError(err error) *HTTPError {
	code := http.StatusInternalServerError

	switch {
	case errors.Is(err, apperror.ErrInvalid):
		code = http.StatusBadRequest
	case errors.Is(err, apperror.ErrUnauthorized):
		code = http.StatusUnauthorized
	case errors.Is(err, apperror.ErrForbidden):
		code = http.StatusForbidden
	case errors.Is(err, apperror.ErrNotFound):
		code = http.StatusNotFound
	case errors.Is(err, apperror.ErrTimeout):
		code = http.StatusRequestTimeout
	case errors.Is(err, apperror.ErrTemporaryDisabled):
		code = http.StatusServiceUnavailable
	case errors.Is(err, apperror.ErrInternal):
		code = http.StatusInternalServerError
	}

	httpError := &HTTPError{
		Error: Err{
			Code:    code,
			Message: http.StatusText(code),
		},
	}

	return httpError
}
