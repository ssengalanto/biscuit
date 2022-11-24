package config

import "fmt"

// Errors used by the config package.
var (
	// ErrConfigInitializationFailed is returned when config initialization failed.
	ErrConfigInitializationFailed = fmt.Errorf("config initialization failed")
)
