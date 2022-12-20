package pgsql

import "fmt"

// Errors used by the pgsql package.
var (
	// ErrInvalidQuery is returned when the query doesn't exist.
	ErrInvalidQuery = fmt.Errorf("invalid query")
	// ErrExecFailed is returned when the query execution failed.
	ErrExecFailed = fmt.Errorf("exec failed")
)
