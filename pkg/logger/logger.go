package logger

import (
	"fmt"
	"strings"

	"github.com/ssengalanto/hex/pkg/constants"
	"github.com/ssengalanto/hex/pkg/interfaces"
)

// New creates a new logger instance for specific log type and environment.
func New(env, logType string) (interfaces.Logger, error) {
	logger, err := buildLogger(env, logType)
	if err != nil {
		return nil, err
	}

	return logger, nil
}

// buildLogger builds a logger for specific log type and environment.
func buildLogger(env, logType string) (interfaces.Logger, error) {
	var logger interfaces.Logger
	var err error

	if logType == "" {
		logType = constants.ZapLogType
	}

	buildProviders := getBuildProviders()
	lastIdx := len(buildProviders) - 1
	for i, provider := range buildProviders {
		matched := provider.logType() == strings.ToLower(logType)
		outOfScope := i == lastIdx && !matched

		if outOfScope {
			return nil,
				fmt.Errorf("%w: invalid log type with value of `%s`, must be one of the ff: `zap`",
					ErrLoggerInitializationFailed, env)
		}

		if !matched {
			continue
		}

		logger, err = provider.build(env)
		if err != nil {
			return nil, err
		}
		break
	}
	return logger, nil
}
