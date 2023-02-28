package config

import (
	"github.com/ssengalanto/biscuit/pkg/config/dotenv"
	"github.com/ssengalanto/biscuit/pkg/config/viper"
	"github.com/ssengalanto/biscuit/pkg/constants"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
)

type buildProvider interface {
	configType() string
	build(env string) (interfaces.Config, error)
}

// viperConfig - buildProvider for viper config.
type viperConfig struct{}

func (v viperConfig) configType() string {
	return constants.ViperConfigType
}

func (v viperConfig) build(env string) (interfaces.Config, error) {
	config, err := viper.New(env)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// dotenvConfig - buildProvider for dotenv config.
type dotenvConfig struct{}

func (d dotenvConfig) configType() string {
	return constants.DotEnvConfigType
}

func (d dotenvConfig) build(env string) (interfaces.Config, error) {
	config, err := dotenv.New(env)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// getBuildProviders returns a slice of buildProvider.
func getBuildProviders() []buildProvider {
	return []buildProvider{
		viperConfig{},
		dotenvConfig{},
	}
}
