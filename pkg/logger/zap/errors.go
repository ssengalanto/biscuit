package zap

import "fmt"

// Errors used by the zap package.
var (
	// ErrZapInitializationFailed is returned when zap initialization failed.
	ErrZapInitializationFailed = fmt.Errorf("zap initialization failed")
)
