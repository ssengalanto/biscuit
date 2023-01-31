package token

import "fmt"

// Errors used by the token entity.
var (
	// ErrInvalidGrantType is returned when the grant type is invalid.
	ErrInvalidGrantType = fmt.Errorf("invalid grant type")
)
