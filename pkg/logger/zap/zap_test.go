package zap_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/logger/zap"
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
			name: "valid env",
			env:  constants.Dev,
			assert: func(t *testing.T, result *zap.Logger, err error) {
				errMsg := fmt.Sprintf("creating new instance should succeed: %s", err)
				require.NotNil(t, result, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name: "invalid env",
			env:  "invalid",
			assert: func(t *testing.T, result *zap.Logger, err error) {
				errMsg := fmt.Sprintf("creating new instance should fail: %s", err)
				require.Nil(t, result, errMsg)
				require.NotNil(t, err, errMsg)
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
