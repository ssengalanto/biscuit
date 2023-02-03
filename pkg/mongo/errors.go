package mongo

import (
	"fmt"
)

// Errors used by the mongo package.

// ErrConnectionFailed is returned when mongo database connection failed.
var ErrConnectionFailed = fmt.Errorf("mongo database connection failed")
