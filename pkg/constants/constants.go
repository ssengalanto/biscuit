package constants

import "time"

const (
	MaxHeaderBytes      = 1 << 20
	IdleTimeout         = time.Minute
	ReadTimeout         = 10 * time.Second
	WriteTimeout        = 30 * time.Second
	ShutdownGracePeriod = 15 * time.Second
	Dev                 = "development"
	Test                = "test"
	Prod                = "production"
)
