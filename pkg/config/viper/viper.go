package viper

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	viper *viper.Viper
}

// New creates a new Config instance.
func New(env string) (*Config, error) {
	config, err := buildViper(env)
	if err != nil {
		return nil, err
	}

	return &Config{
		viper: config,
	}, nil
}

// buildViper builds a new viper instance for specific environment with predefined configuration.
func buildViper(env string) (*viper.Viper, error) {
	var config *viper.Viper
	var err error

	buildProviders := getBuildProviders()
	lastIdx := len(buildProviders) - 1
	for i, provider := range buildProviders {
		matched := provider.env() == strings.ToLower(env)
		outOfScope := i == lastIdx && !matched

		if outOfScope {
			return nil,
				fmt.Errorf("%w: invalid env with value of `%s`, must be one of the ff: `development`, `testing`, `production`",
					ErrViperInitializationFailed, env)
		}

		if !matched {
			continue
		}

		config, err = provider.build()
		if err != nil {
			return nil, err
		}
		break
	}
	return config, nil
}

// Get can retrieve any value given the key to use.
func (c *Config) Get(key string) any {
	return c.viper.Get(key)
}

// GetBool returns the value associated with the key as a boolean.
func (c *Config) GetBool(key string) bool {
	return c.viper.GetBool(key)
}

// GetFloat64 returns the value associated with the key as a float64.
func (c *Config) GetFloat64(key string) float64 {
	return c.viper.GetFloat64(key)
}

// GetInt returns the value associated with the key as an int.
func (c *Config) GetInt(key string) int {
	return c.viper.GetInt(key)
}

// GetString returns the value associated with the key as a string.
func (c *Config) GetString(key string) string {
	return c.viper.GetString(key)
}

// GetTime returns the value associated with the key as time.
func (c *Config) GetTime(key string) time.Time {
	return c.viper.GetTime(key)
}

// GetDuration returns the value associated with the key as a duration.
func (c *Config) GetDuration(key string) time.Duration {
	return c.viper.GetDuration(key)
}
