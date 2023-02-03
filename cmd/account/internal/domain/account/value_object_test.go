package account_test

import (
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmail_IsValid(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name:    "it should be a valid email",
			payload: gofakeit.Email(),
			assert: func(t *testing.T, result bool, err error) {
				assert.True(t, result)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should be an invalid email",
			payload: "invalid-email",
			assert: func(t *testing.T, result bool, err error) {
				assert.False(t, result)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email := account.Email(tc.payload)
			ok, err := email.IsValid()
			tc.assert(t, ok, err)
		})
	}
}

func TestEmail_Update(t *testing.T) {
	testCases := []struct {
		name    string
		current string
		update  string
		assert  func(t *testing.T, expected account.Email, actual account.Email, err error)
	}{
		{
			name:    "it should change the email successfully",
			current: gofakeit.Email(),
			update:  gofakeit.Email(),
			assert: func(t *testing.T, expected account.Email, actual account.Email, err error) {
				assert.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should fail to change the email",
			current: gofakeit.Email(),
			update:  "invalid-email",
			assert: func(t *testing.T, expected account.Email, actual account.Email, err error) {
				assert.NotEqual(t, expected, actual)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := account.Email(tc.current)
			ne, err := e.Update(tc.update)
			tc.assert(t, account.Email(tc.update), ne, err)
		})
	}
}

func TestEmail_String(t *testing.T) {
	t.Run("it should convert email to string", func(t *testing.T) {
		email := account.Email(gofakeit.Email()).String()
		kind := reflect.TypeOf(email).String()
		require.Equal(t, "string", kind)
	})
}

func TestPassword_IsValid(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name:    "it should be a valid password",
			payload: createValidPassword(),
			assert: func(t *testing.T, result bool, err error) {
				assert.True(t, result)
				require.Nil(t, err)
			},
		},
		{
			name:    "it should be an invalid password",
			payload: createInvalidPassword(),
			assert: func(t *testing.T, result bool, err error) {
				assert.False(t, result)
				require.NotNil(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			password := account.Password(tc.payload)
			ok, err := password.IsValid()
			tc.assert(t, ok, err)
		})
	}
}

func TestPassword_Update(t *testing.T) {
	testCases := []struct {
		name    string
		current string
		update  string
		assert  func(t *testing.T, expected account.Password, actual account.Password, err error)
	}{
		{
			name:    "it should change the password successfully",
			current: createValidPassword(),
			update:  createValidPassword(),
			assert: func(t *testing.T, expected account.Password, actual account.Password, err error) {
				assert.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should fail to change the password",
			current: createValidPassword(),
			update:  createInvalidPassword(),
			assert: func(t *testing.T, expected account.Password, actual account.Password, err error) {
				assert.NotEqual(t, expected, actual)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pw := account.Password(tc.current)
			npw, err := pw.Update(tc.update)
			tc.assert(t, account.Password(tc.update), npw, err)
		})
	}
}

func TestPassword_Hash(t *testing.T) {
	t.Run("it should hash the password", func(t *testing.T) {
		pw := account.Password(createValidPassword())
		hpw, err := pw.Hash()
		assert.NotEqual(t, pw, hpw)
		require.NoError(t, err)
	})
}

func TestPassword_Check(t *testing.T) {
	t.Run("it should match the password", func(t *testing.T) {
		pw := account.Password(createValidPassword())
		hpw, err := pw.Hash()
		require.NoError(t, err)

		err = hpw.Check(pw.String())
		require.NoError(t, err)
	})
	t.Run("it should fail the password check", func(t *testing.T) {
		pw := account.Password(createValidPassword())
		hpw, err := pw.Hash()
		require.NoError(t, err)

		err = hpw.Check(pw.String())
		require.NoError(t, err)
	})
}

func TestPassword_String(t *testing.T) {
	t.Run("it should convert password to string", func(t *testing.T) {
		password := account.Password(createValidPassword()).String()
		kind := reflect.TypeOf(password).String()
		require.Equal(t, "string", kind)
	})
}

func createValidPassword() string {
	return gofakeit.Password(true, true, true, true, false, 10)
}

func createInvalidPassword() string {
	return gofakeit.Password(true, true, true, true, false, 5)
}
