package logger_test

import (
	"fmt"
	"testing"

	"github.com/ssengalanto/hex/pkg/constants"
	"github.com/ssengalanto/hex/pkg/interfaces"
	"github.com/ssengalanto/hex/pkg/logger"
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
			name:    "valid env",
			env:     constants.Dev,
			logType: constants.ZapLogType,
			assert: func(t *testing.T, result interfaces.Logger, err error) {
				errMsg := fmt.Sprintf("creating new instance should succeed: %s", err)
				require.NotNil(t, result, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "invalid env",
			env:     "invalid",
			logType: constants.ZapLogType,
			assert: func(t *testing.T, result interfaces.Logger, err error) {
				errMsg := fmt.Sprintf("creating new instance should fail: %s", err)
				require.Nil(t, result, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
		{
			name:    "invalid log type",
			env:     constants.Dev,
			logType: "invalid",
			assert: func(t *testing.T, result interfaces.Logger, err error) {
				errMsg := fmt.Sprintf("creating new instance should fail: %s", err)
				require.Nil(t, result, errMsg)
				require.NotNil(t, err, errMsg)
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
