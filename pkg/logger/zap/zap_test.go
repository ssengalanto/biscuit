package zap_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/biscuit/pkg/constants"
	"github.com/ssengalanto/biscuit/pkg/logger/zap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name   string
		env    string
		assert func(t *testing.T, result *zap.Logger, err error)
	}{
		{
			name: "it should build for development environment",
			env:  constants.Dev,
			assert: func(t *testing.T, result *zap.Logger, err error) {
				assert.NotNil(t, result)
				require.NoError(t, err)
			},
		},
		{
			name: "it should build for test environment",
			env:  constants.Test,
			assert: func(t *testing.T, result *zap.Logger, err error) {
				assert.NotNil(t, result)
				require.NoError(t, err)
			},
		},
		{
			name: "it should build for prod environment",
			env:  constants.Prod,
			assert: func(t *testing.T, result *zap.Logger, err error) {
				assert.NotNil(t, result)
				require.NoError(t, err)
			},
		},
		{
			name: "it should fail to build due to an invalid env",
			env:  "invalid",
			assert: func(t *testing.T, result *zap.Logger, err error) {
				assert.Nil(t, result)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			log, err := zap.New(tc.env)
			tc.assert(t, log, err)
		})
	}
}

func TestLogger_Info(t *testing.T) {
	msg := gofakeit.Word()
	log, observedLogs := zap.NewTestInstance(zapcore.InfoLevel)
	log.Info(msg, nil)

	allLogs := observedLogs.All()
	assert.Equal(t, msg, allLogs[0].Message)
	require.Equal(t, 1, observedLogs.Len())
}

func TestLogger_Error(t *testing.T) {
	msg := gofakeit.Word()
	log, observedLogs := zap.NewTestInstance(zapcore.ErrorLevel)
	log.Error(msg, nil)

	allLogs := observedLogs.All()
	assert.Equal(t, msg, allLogs[0].Message)
	require.Equal(t, 1, observedLogs.Len())
}

func TestLogger_Debug(t *testing.T) {
	msg := gofakeit.Word()
	log, observedLogs := zap.NewTestInstance(zapcore.DebugLevel)
	log.Debug(msg, nil)

	allLogs := observedLogs.All()
	assert.Equal(t, msg, allLogs[0].Message)
	require.Equal(t, 1, observedLogs.Len())
}

func TestLogger_Warn(t *testing.T) {
	msg := gofakeit.Word()
	log, observedLogs := zap.NewTestInstance(zapcore.WarnLevel)
	log.Warn(msg, nil)

	allLogs := observedLogs.All()
	assert.Equal(t, msg, allLogs[0].Message)
	require.Equal(t, 1, observedLogs.Len())
}
