package config_test

import (
	"testing"

	"github.com/ssengalanto/potato-project/pkg/config"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/stretchr/testify/require"
)

func TestGetInstance(t *testing.T) {
	c1, err := config.GetInstance()
	require.NotNil(t, c1)
	require.Nil(t, err)

	c2, _ := config.GetInstance()
	require.Equal(t, &c1, &c2, "should return the same instance")
}

func TestConfig_Get(t *testing.T) {
	c, err := config.GetInstance()
	require.Nil(t, err)

	v := c.Get(constants.AppName)
	require.Equal(t, v, "potato-project")
}
