package server

import "fmt"

// Errors used by the server package.

// ErrServerClosed is returned when the server failed to start.
var ErrServerClosed = fmt.Errorf("server is closed")
