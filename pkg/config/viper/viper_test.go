package viper_test

import (
	"fmt"
	"testing"

	"github.com/ssengalanto/potato-project/pkg/config/viper"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name   string
		env    string
		assert func(t *testing.T, result *viper.Config, err error)
	}{
		{
			name: "valid env",
			env:  constants.Dev,
			assert: func(t *testing.T, result *viper.Config, err error) {
				errMsg := fmt.Sprintf("creating new instance should succeed: %s", err)
				require.NotNil(t, result, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name: "invalid env",
			env:  "invalid",
			assert: func(t *testing.T, result *viper.Config, err error) {
				errMsg := fmt.Sprintf("creating new instance should fail: %s", err)
				require.Nil(t, result, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			log, err := viper.New(tc.env)
			tc.assert(t, log, err)
		})
	}
}
