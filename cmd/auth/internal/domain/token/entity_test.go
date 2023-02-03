package token_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/auth/internal/domain/token"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("it should create a new token entity", func(t *testing.T) {
		t.Parallel()
		acct := token.New(uuid.New(), gofakeit.UUID())
		assert.NotNil(t, acct)
	})
}
