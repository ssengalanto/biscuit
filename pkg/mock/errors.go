package mock

import "fmt"

// Errors used by the mock package.
var (
	// ErrDatabaseStubConnection is returned when stub database connection failed.
	ErrDatabaseStubConnection = fmt.Errorf("stub database connection failed")
)
