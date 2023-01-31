package token_test

import (
	"reflect"
	"testing"

	"github.com/ssengalanto/biscuit/cmd/auth/internal/domain/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGrantType_IsValid(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name:    "it should be a valid grant type",
			payload: token.GrantTypePassword,
			assert: func(t *testing.T, result bool, err error) {
				assert.True(t, result)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should fail the validation due to empty value",
			payload: "",
			assert: func(t *testing.T, result bool, err error) {
				assert.False(t, result)
				require.Error(t, err)
			},
		},
		{
			name:    "it should be an invalid grant type",
			payload: "invalid-grant-type",
			assert: func(t *testing.T, result bool, err error) {
				assert.False(t, result)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			gt := token.GrantType(tc.payload)
			ok, err := gt.IsValid()
			tc.assert(t, ok, err)
		})
	}
}

func TestGrantType_String(t *testing.T) {
	t.Parallel()
	t.Run("it should convert email to string", func(t *testing.T) {
		t.Parallel()
		gt := token.GrantType("password").String()
		kind := reflect.TypeOf(gt).String()
		require.Equal(t, "string", kind)
	})
}
