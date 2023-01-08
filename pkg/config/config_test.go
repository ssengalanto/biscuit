package config_test

import (
	"fmt"
	"testing"

	"github.com/ssengalanto/hex/pkg/config"
	"github.com/ssengalanto/hex/pkg/constants"
	"github.com/ssengalanto/hex/pkg/interfaces"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name       string
		env        string
		configType string
		assert     func(t *testing.T, result interfaces.Config, err error)
	}{
		// TODO: disabled temporary due to github workflow env path
		// {
		//	name:       "valid env",
		//	env:        constants.Dev,
		//	configType: constants.ViperConfigType,
		//	assert: func(t *testing.T, result interfaces.Config, err error) {
		//		errMsg := fmt.Sprintf("creating new instance should succeed: %s", err)
		//		require.NotNil(t, result, errMsg)
		//		require.Nil(t, err, errMsg)
		//	},
		// },
		{
			name:       "invalid env",
			env:        "invalid",
			configType: constants.ViperConfigType,
			assert: func(t *testing.T, result interfaces.Config, err error) {
				errMsg := fmt.Sprintf("creating new instance should fail: %s", err)
				require.Nil(t, result, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
		{
			name:       "invalid config type",
			env:        constants.Dev,
			configType: "invalid",
			assert: func(t *testing.T, result interfaces.Config, err error) {
				errMsg := fmt.Sprintf("creating new instance should fail: %s", err)
				require.Nil(t, result, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg, err := config.New(tc.env, tc.configType)
			tc.assert(t, cfg, err)
		})
	}
}
