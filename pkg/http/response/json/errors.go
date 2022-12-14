package json

import "fmt"

// Errors used by the HTTP json package.
var (
	// ErrBadlyFormedJSON is returned when JSON is badly formed.
	ErrBadlyFormedJSON = fmt.Errorf("body contains badly-formed")
	// ErrIncorrectJSONType is returned when JSON contains incorrect type.
	ErrIncorrectJSONType = fmt.Errorf("body contains incorrect JSON type")
	// ErrEmptyBody is returned when JSON body is empty.
	ErrEmptyBody = fmt.Errorf("body must not be empty")
	// ErrUnknownField is returned when JSON body contains unknown field.
	ErrUnknownField = fmt.Errorf("body must not be empty")
	// ErrBodySizeExceeded is returned when JSON body size exceeded the limit.
	ErrBodySizeExceeded = fmt.Errorf("body size exceeded")
	// ErrExtraneousBody is returned when JSON body contains more than a single JSON.
	ErrExtraneousBody = fmt.Errorf("body must only contain a single JSON value")
)
