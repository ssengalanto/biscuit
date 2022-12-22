package viper

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/ssengalanto/potato-project/pkg/constants"
)

type buildProvider interface {
	env() string
	build() (*viper.Viper, error)
}

// development - buildProvider for development environment.
type development struct{}

func (d development) env() string {
	return constants.Dev
}
func (d development) build() (*viper.Viper, error) {
	v, err := createViperInstance(d.env())
	if err != nil {
		return nil, err
	}

	return v, nil
}

// testing - buildProvider for testing environment.
type testing struct{}

func (t testing) env() string {
	return constants.Test
}
func (t testing) build() (*viper.Viper, error) {
	v, err := createViperInstance(t.env())
	if err != nil {
		return nil, err
	}

	return v, nil
}

// production - buildProvider for production environment.
type production struct{}

func (p production) env() string {
	return constants.Prod
}
func (p production) build() (*viper.Viper, error) {
	v, err := createViperInstance(p.env())
	if err != nil {
		return nil, err
	}

	return v, nil
}

// createViperInstance creates a new viper instance for specific environment.
func createViperInstance(env string) (*viper.Viper, error) {
	pkg := "potato-project"

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	v := viper.New()

	// build path
	v.AddConfigPath(".")
	// local path
	v.AddConfigPath(strings.SplitAfter(wd, pkg)[0])
	v.SetConfigName(fmt.Sprintf(".env.%s", env))
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err = v.ReadInConfig(); err != nil {
		if ok := errors.Is(err, viper.ConfigFileNotFoundError{}); ok {
			return nil, ErrConfigFileNotFound
		}

		return nil, ErrCannotReadConfig
	}

	return v, nil
}

// getBuildProviders returns a slice of buildProvider.
func getBuildProviders() []buildProvider {
	return []buildProvider{
		development{}, testing{}, production{},
	}
}
