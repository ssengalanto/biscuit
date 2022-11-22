package logger

import (
	"github.com/ssengalanto/potato-project/pkg/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type buildProvider interface {
	env() string
	build() (*zap.Logger, error)
}

// development - buildProvider for development environment.
type development struct{}

func (d development) env() string {
	return constants.Dev
}
func (d development) build() (*zap.Logger, error) {
	cfg := createDevelopmentConfig()

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

// testing - buildProvider for testing environment.
type testing struct{}

func (t testing) env() string {
	return constants.Test
}
func (t testing) build() (*zap.Logger, error) {
	cfg := createDevelopmentConfig()

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

// production - buildProvider for production environment.
type production struct{}

func (p production) env() string {
	return constants.Prod
}
func (p production) build() (*zap.Logger, error) {
	cfg := createProductionConfig()

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

// createDevelopmentConfig creates a new zap.Config for development environment.
func createDevelopmentConfig() zap.Config {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return cfg
}

// createProductionConfig creates a new zap.Config for production environment.
func createProductionConfig() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return cfg
}

// getBuildProviders returns a slice of buildProvider.
func getBuildProviders() []buildProvider {
	return []buildProvider{
		development{}, testing{}, production{},
	}
}
