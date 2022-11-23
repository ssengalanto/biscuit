//nolint:gosec // unnecessary rules for this package
package constants

import "time"

// Constants for application invariants.
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

// Constants for packages.
const (
	ZapLogType = "zap"
)

// Constants for environment variable keys.
const (
	AppName            = "APP_NAME"
	AppEnv             = "APP_ENV"
	AccountServicePort = "ACCOUNT_SERVICE_PORT"
	PostgresUser       = "POSTGRES_USER"
	PostgresPassword   = "POSTGRES_PASSWORD"
	PostgresDB         = "POSTGRES_DB"
	PostgresPort       = "POSTGRES_PORT"
	PostgresDSN        = "POSTGRES_DSN"
	RedisPort          = "REDIS_PORT"
	RedisInsightPort   = "REDIS_INSIGHT_PORT"
	LogType            = "LOG_TYPE"
)
