package config

import (
	"fmt"
	"strings"

	"github.com/ssengalanto/biscuit/pkg/constants"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
)

// New creates a new config instance for specific config type and environment.
func New(env, configType string) (interfaces.Config, error) {
	config, err := buildConfig(env, configType)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// buildConfig builds a config instance for specific config type and environment.
func buildConfig(env, configType string) (interfaces.Config, error) {
	var config interfaces.Config
	var err error

	if configType == "" {
		configType = constants.ViperConfigType
	}

	buildProviders := getBuildProviders()
	lastIdx := len(buildProviders) - 1
	for i, provider := range buildProviders {
		matched := provider.configType() == strings.ToLower(configType)
		outOfScope := i == lastIdx && !matched

		if outOfScope {
			return nil,
				fmt.Errorf("%w: invalid config type with value of `%s`, must be one of the ff: viper, dotenv",
					ErrConfigInitializationFailed, configType)
		}

		if !matched {
			continue
		}

		config, err = provider.build(env)
		if err != nil {
			return nil, err
		}
		break
	}
	return config, nil
}
