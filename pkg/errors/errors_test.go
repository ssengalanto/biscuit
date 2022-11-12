package errors_test

import (
	"errors"
	"fmt"
	"testing"

	apperr "github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	err := apperr.New("test error")
	require.NotNil(t, err, "error shouldn't be nil")
}

func TestError_Error(t *testing.T) {
	msg := "test error"
	err := apperr.New(msg)
	require.EqualError(t, err, msg, "error should have the same error message")
}

func TestError_Wrap(t *testing.T) {
	err := apperr.Wrap(fmt.Errorf("test error: %w", apperr.ErrInternal))
	require.NotNil(t, err, "error shouldn't be nil")
	require.True(t, errors.Is(err, apperr.ErrInternal), "error is not internal")
}
