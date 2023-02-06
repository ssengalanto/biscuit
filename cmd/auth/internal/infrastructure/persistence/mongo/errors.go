package mongo

import "fmt"

// Errors used by the mongo package.
var (
	// ErrDeleteRecordFailed is returned when deleting a record failed.
	ErrDeleteRecordFailed = fmt.Errorf("record deletion failed")
)
