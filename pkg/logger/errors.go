package logger

import "fmt"

// Errors used by the logger package.
var (
	// ErrLoggerInitializationFailed is returned when logger initialization failed.
	ErrLoggerInitializationFailed = fmt.Errorf("logger initialization failed")
)
