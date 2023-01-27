package constants

import "time"

// Constants for application configurations.
const (
	MaxHeaderBytes      = 1 << 20
	RateLimit           = 120
	IdleTimeout         = time.Minute
	ReadTimeout         = 10 * time.Second
	WriteTimeout        = 30 * time.Second
	Timeout             = time.Second * 60
	ShutdownGracePeriod = 15 * time.Second
	Dev                 = "development"
	Test                = "test"
	Prod                = "production"
)

// Constants for internal packages.
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
	RedisURL           = "REDIS_URL"
	RedisDB            = "REDIS_DB"
	RedisPassword      = "REDIS_PASSWORD"
	RedisInsightPort   = "REDIS_INSIGHT_PORT"
	RedisInsightURL    = "REDIS_INSIGHT_URL"
	LogType            = "LOG_TYPE"
	ConfigType         = "CONFIG_TYPE"
)
