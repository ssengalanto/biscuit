//nolint:gochecknoglobals,ireturn // unnecessary rules for this package
package config

import (
	"errors"
	"os"
	"strings"
	"sync"
	"time"

	v "github.com/spf13/viper"
)

var (
	configInstance *config
	viperInstance  *v.Viper
	once           sync.Once
	mu             sync.Mutex
)

// Config is an interface consisting of the core config methods.
type Config interface {
	Get(key string) any
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetIntSlice(key string) []int
	GetString(key string) string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	IsSet(key string) bool
	AllKeys() []string
	AllSettings() map[string]any
}

type config struct {
	viper *v.Viper
}

// GetInstance returns a new/existing config instance that implements the config interface.
func GetInstance() (Config, error) {
	mu.Lock()
	defer mu.Unlock()

	if configInstance != nil {
		return configInstance, nil
	}

	c, err := create()
	if err != nil {
		return nil, err
	}

	configInstance = c

	return configInstance, err
}

// create initializes the viper instance, and it creates and configure a new config struct.
func create() (*config, error) {
	pkg := "potato-project"

	if viperInstance == nil {
		once.Do(func() {
			viperInstance = v.New()
		})
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cfg := config{viper: viperInstance}
	cfg.viper.AddConfigPath(strings.SplitAfter(wd, pkg)[0])
	cfg.viper.SetConfigName(".env")
	cfg.viper.SetConfigType("env")
	cfg.viper.AutomaticEnv()

	if err = cfg.viper.ReadInConfig(); err != nil {
		if ok := errors.Is(err, v.ConfigFileNotFoundError{}); ok {
			return nil, ErrConfigFileNotFound
		}

		return nil, ErrCannotReadConfig
	}

	return &cfg, nil
}

// Get can retrieve any value given the key to use.
func (c *config) Get(key string) any {
	return c.viper.Get(key)
}

// GetBool returns the value associated with the key as a boolean.
func (c *config) GetBool(key string) bool {
	return c.viper.GetBool(key)
}

// GetFloat64 returns the value associated with the key as a float64.
func (c *config) GetFloat64(key string) float64 {
	return c.viper.GetFloat64(key)
}

// GetInt returns the value associated with the key as an int.
func (c *config) GetInt(key string) int {
	return c.viper.GetInt(key)
}

// GetIntSlice returns the value associated with the key as a slice of int values.
func (c *config) GetIntSlice(key string) []int {
	return c.viper.GetIntSlice(key)
}

// GetString returns the value associated with the key as a string.
func (c *config) GetString(key string) string {
	return c.viper.GetString(key)
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (c *config) GetStringSlice(key string) []string {
	return c.viper.GetStringSlice(key)
}

// GetTime returns the value associated with the key as time.
func (c *config) GetTime(key string) time.Time {
	return c.viper.GetTime(key)
}

// GetDuration returns the value associated with the key as a duration.
func (c *config) GetDuration(key string) time.Duration {
	return c.viper.GetDuration(key)
}

// IsSet checks to see if the key has been set in any of the data locations.
func (c *config) IsSet(key string) bool {
	return c.viper.IsSet(key)
}

// AllKeys returns all keys holding a value, regardless of where they are set.
func (c *config) AllKeys() []string {
	return c.viper.AllKeys()
}

// AllSettings merges all settings and returns them as a map[string]any.
func (c *config) AllSettings() map[string]any {
	return c.viper.AllSettings()
}
