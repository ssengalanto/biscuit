package logger_test

import (
	"testing"

	"github.com/ssengalanto/biscuit/pkg/constants"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"github.com/ssengalanto/biscuit/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name    string
		env     string
		logType string
		assert  func(t *testing.T, result interfaces.Logger, err error)
	}{
		{
			name:    "it should build for development environment",
			env:     constants.Dev,
			logType: constants.ZapLogType,
			assert: func(t *testing.T, result interfaces.Logger, err error) {
				assert.NotNil(t, result)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should build for test environment",
			env:     constants.Test,
			logType: constants.ZapLogType,
			assert: func(t *testing.T, result interfaces.Logger, err error) {
				assert.NotNil(t, result)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should build for prod environment",
			env:     constants.Prod,
			logType: constants.ZapLogType,
			assert: func(t *testing.T, result interfaces.Logger, err error) {
				assert.NotNil(t, result)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should fail to build due to an invalid env",
			env:     "invalid",
			logType: constants.ZapLogType,
			assert: func(t *testing.T, result interfaces.Logger, err error) {
				assert.Nil(t, result)
				require.Error(t, err)
			},
		},
		{
			name:    "it should fail to build due to an invalid log type",
			env:     constants.Dev,
			logType: "invalid",
			assert: func(t *testing.T, result interfaces.Logger, err error) {
				assert.Nil(t, result)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			log, err := logger.New(tc.env, tc.logType)
			tc.assert(t, log, err)
		})
	}
}
