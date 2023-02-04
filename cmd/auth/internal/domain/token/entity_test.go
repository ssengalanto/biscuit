package token_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/auth/internal/domain/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("it should create a new token entity", func(t *testing.T) {
		tk := newToken()
		assert.NotNil(t, tk)
	})
}

func TestEntity_GenerateAccessToken(t *testing.T) {
	t.Run("it should generate a new access token", func(t *testing.T) {
		tk := newToken()
		err := tk.GenerateAccessToken(gofakeit.Word(), newBase64RSAPrivateKey(), 15*time.Minute)
		assert.NotEmpty(t, tk.AccessToken)
		require.NoError(t, err)
	})

	t.Run("it should fail to generate a new access token", func(t *testing.T) {
		tk := newToken()
		err := tk.GenerateAccessToken(gofakeit.Word(), newBase64RSAPublicKey(), 15*time.Minute)
		assert.Empty(t, tk.AccessToken)
		require.Error(t, err)
	})
}

func TestEntity_GenerateRefreshToken(t *testing.T) {
	t.Run("it should generate a new refresh token", func(t *testing.T) {
		tk := newToken()
		err := tk.GenerateRefreshToken(gofakeit.Word(), newBase64RSAPrivateKey(), 60*time.Minute)
		assert.NotEmpty(t, tk.RefreshToken)
		require.NoError(t, err)
	})

	t.Run("it should fail to generate a new refresh token", func(t *testing.T) {
		tk := newToken()
		err := tk.GenerateRefreshToken(gofakeit.Word(), newBase64RSAPublicKey(), 60*time.Minute)
		assert.Empty(t, tk.RefreshToken)
		require.Error(t, err)
	})
}

func newToken() token.Entity {
	return token.New(uuid.New(), gofakeit.UUID())
}
