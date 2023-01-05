package redis

import (
	"fmt"
)

// Errors used by the redis package.

// ErrConnectionFailed is returned when redis client connection failed.
var ErrConnectionFailed = fmt.Errorf("redis client connection failed")
