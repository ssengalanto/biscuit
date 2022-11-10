package errors_test

import (
	"errors"
	"fmt"
	"testing"

	apperr "github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	err := apperr.New("test error")
	assert.NotNil(t, err, "error shouldn't be nil")
}

func TestError(t *testing.T) {
	msg := "test error"
	err := apperr.New(msg)
	assert.EqualError(t, err, msg, "error should have the same error message")
}

func TestWrap(t *testing.T) {
	err := apperr.Wrap(fmt.Errorf("test error: %w", apperr.ErrInternal))
	assert.NotNil(t, err, "error shouldn't be nil")
	assert.True(t, errors.Is(err, apperr.ErrInternal), "error is not internal")
}
