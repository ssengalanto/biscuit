package logger_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/potato-project/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	log, err := logger.New()
	require.NotNil(t, log)
	require.Nil(t, err)
}

func TestLogger_Info(t *testing.T) {
	msg := gofakeit.Word()
	log, observedLogs := logger.NewTestInstance(zap.InfoLevel)
	log.Info(msg, nil)

	allLogs := observedLogs.All()
	assert.Equal(t, msg, allLogs[0].Message)
	require.Equal(t, 1, observedLogs.Len())
}

func TestLogger_Error(t *testing.T) {
	msg := gofakeit.Word()
	log, observedLogs := logger.NewTestInstance(zap.ErrorLevel)
	log.Error(msg, nil)

	allLogs := observedLogs.All()
	assert.Equal(t, msg, allLogs[0].Message)
	require.Equal(t, 1, observedLogs.Len())
}

func TestLogger_Debug(t *testing.T) {
	msg := gofakeit.Word()
	log, observedLogs := logger.NewTestInstance(zap.DebugLevel)
	log.Debug(msg, nil)

	allLogs := observedLogs.All()
	assert.Equal(t, msg, allLogs[0].Message)
	require.Equal(t, 1, observedLogs.Len())
}

func TestLogger_Warn(t *testing.T) {
	msg := gofakeit.Word()
	log, observedLogs := logger.NewTestInstance(zap.InfoLevel)
	log.Warn(msg, nil)

	allLogs := observedLogs.All()
	assert.Equal(t, msg, allLogs[0].Message)
	require.Equal(t, 1, observedLogs.Len())
}
