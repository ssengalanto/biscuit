package viper

import "fmt"

// Errors used by the viper package.
var (
	// ErrViperInitializationFailed is returned when viper initialization failed.
	ErrViperInitializationFailed = fmt.Errorf("viper initialization failed")
	// ErrConfigFileNotFound is returned when config file is not found.
	ErrConfigFileNotFound = fmt.Errorf("config file not found")
	// ErrCannotReadConfig is returned when it can't read the config file.
	ErrCannotReadConfig = fmt.Errorf("cannot read config file")
)
