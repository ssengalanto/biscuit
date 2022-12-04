package pgsql

import "fmt"

// Errors used by the pgsql package.
var (
	// ErrInvalidQuery is returned when the query doesn't exist.
	ErrInvalidQuery = fmt.Errorf("invalid query")
)
