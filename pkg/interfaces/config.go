package interfaces

import "time"

// Config is an interface consisting of the core config methods.
type Config interface {
	Get(key string) any
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetString(key string) string
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
}
