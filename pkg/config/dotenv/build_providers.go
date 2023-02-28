package dotenv

import (
	"github.com/ssengalanto/biscuit/pkg/constants"
)

type buildProvider interface {
	env() string
	build() (*Config, error)
}

// development - buildProvider for development environment.
type development struct{}

func (d development) env() string {
	return constants.Dev
}

func (d development) build() (*Config, error) {
	c, err := createViperInstance(d.env())
	if err != nil {
		return nil, err
	}

	return c, nil
}

// testing - buildProvider for testing environment.
type testing struct{}

func (t testing) env() string {
	return constants.Test
}
func (t testing) build() (*Config, error) {
	c, err := createViperInstance(t.env())
	if err != nil {
		return nil, err
	}

	return c, nil
}

// production - buildProvider for production environment.
type production struct{}

func (p production) env() string {
	return constants.Prod
}
func (p production) build() (*Config, error) {
	c, err := createViperInstance(p.env())
	if err != nil {
		return nil, err
	}

	return c, nil
}

// createViperInstance creates a new viper instance for specific environment.
func createViperInstance(env string) (*Config, error) {
	return New(env)
}

// getBuildProviders returns a slice of buildProvider.
func getBuildProviders() []buildProvider {
	return []buildProvider{
		development{}, testing{}, production{},
	}
}
