package logger

import (
	"github.com/ssengalanto/potato-project/pkg/config"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type Logger struct {
	log *zap.Logger
}

// New creates a new Logger instance.
func New() (Logger, error) {
	var zapCfg zap.Config

	env := getAppEnv()
	if env == constants.Prod {
		zapCfg = zap.NewProductionConfig()
	} else {
		zapCfg = zap.NewDevelopmentConfig()
	}

	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	zapLogger, err := zapCfg.Build()
	if err != nil {
		return Logger{}, ErrLoggerInitializationFailed
	}

	return Logger{
		log: zapLogger,
	}, nil
}

// Info logs a message at info level.
func (l *Logger) Info(msg string, fields interfaces.Fields) {
	zapFields := mapToZapFields(fields)
	l.log.Info(msg, zapFields...)
}

// Error logs a message at error level.
func (l *Logger) Error(msg string, fields interfaces.Fields) {
	zapFields := mapToZapFields(fields)
	l.log.Error(msg, zapFields...)
}

// Debug logs a message at debug level.
func (l *Logger) Debug(msg string, fields interfaces.Fields) {
	zapFields := mapToZapFields(fields)
	l.log.Debug(msg, zapFields...)
}

// Warn logs a message at warn level.
func (l *Logger) Warn(msg string, fields interfaces.Fields) {
	zapFields := mapToZapFields(fields)
	l.log.Warn(msg, zapFields...)
}

// Fatal logs a message at fatal level.
func (l *Logger) Fatal(msg string, fields interfaces.Fields) {
	zapFields := mapToZapFields(fields)
	l.log.Fatal(msg, zapFields...)
}

// Panic logs a message at panic level.
func (l *Logger) Panic(msg string, fields interfaces.Fields) {
	zapFields := mapToZapFields(fields)
	l.log.Panic(msg, zapFields...)
}

// mapToZapFields maps the logger fields to zap fields.
func mapToZapFields(fields map[string]any) []zap.Field {
	var zapFields []zap.Field
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}

	return zapFields
}

// getAppEnv gets the application environment from .env.
func getAppEnv() string {
	cfg, err := config.GetInstance()
	if err != nil {
		return ""
	}

	return cfg.GetString(constants.AppEnv)
}

// NewTestInstance creates a new Logger instance for testing.
// Use only for testing purposes.
func NewTestInstance(level zapcore.Level) (Logger, *observer.ObservedLogs) {
	observedZapCore, observedLogs := observer.New(level)
	observedLogger := zap.New(observedZapCore)

	return Logger{
		log: observedLogger,
	}, observedLogs
}
