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
	ZapLogType      = "zap"
	ViperConfigType = "viper"
	PgsqlDriver     = "pgx"
	SlqMockDriver   = "sqlmock"
)

// Constants for environment variable keys.
const (
	AppName            = "APP_NAME"
	AppEnv             = "APP_ENV"
	AccountServicePort = "ACCOUNT_SERVICE_PORT"
	PgsqlUser          = "PGSQL_USER"
	PgsqlPassword      = "PGSQL_PASSWORD"
	PgsqlDB            = "PGSQL_DB"
	PgsqlPort          = "PGSQL_PORT"
	PgsqlDSN           = "PGSQL_DSN"
	RedisPort          = "REDIS_PORT"
	RedisInsightPort   = "REDIS_INSIGHT_PORT"
	LogType            = "LOG_TYPE"
	ConfigType         = "CONFIG_TYPE"
)
