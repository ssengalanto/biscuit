package config

import (
	"github.com/ssengalanto/potato-project/pkg/config/viper"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
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

// getBuildProviders returns a slice of buildProvider.
func getBuildProviders() []buildProvider {
	return []buildProvider{
		viperConfig{},
	}
}
