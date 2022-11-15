package account_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/stretchr/testify/require"
)

func TestEmail_IsValid(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name:    "valid email",
			payload: gofakeit.Email(),
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("email should be valid: %s", err)
				require.True(t, result, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "invalid email",
			payload: "invalid-email",
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("email should be invalid: %s", err)
				require.False(t, result, errMsg)
				require.NotNil(t, err, errMsg)
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
			name:    "change email success",
			current: gofakeit.Email(),
			update:  gofakeit.Email(),
			assert: func(t *testing.T, expected account.Email, actual account.Email, err error) {
				errMsg := fmt.Sprintf("email change should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "change email failed",
			current: gofakeit.Email(),
			update:  "invalid-email",
			assert: func(t *testing.T, expected account.Email, actual account.Email, err error) {
				errMsg := fmt.Sprintf("email change should fail: %s", err)
				require.NotEqual(t, expected, actual, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email := account.Email(tc.current)
			newEmail, err := email.Update(tc.update)
			tc.assert(t, account.Email(tc.update), newEmail, err)
		})
	}
}

func TestEmail_String(t *testing.T) {
	email := account.Email(gofakeit.Email()).String()
	kind := reflect.TypeOf(email).String()
	require.Equal(t, "string", kind, "type should be `string`")
}

func createValidPassword() string {
	return gofakeit.Password(true, true, true, true, false, 10)
}

func createInvalidPassword() string {
	return gofakeit.Password(true, true, true, true, false, 5)
}

func TestPassword_IsValid(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name:    "valid password",
			payload: createValidPassword(),
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("password should be valid: %s", err)
				require.True(t, result, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "invalid password",
			payload: createInvalidPassword(),
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("password should be valid: %s", err)
				require.False(t, result, errMsg)
				require.NotNil(t, err, errMsg)
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
			name:    "change password success",
			current: createValidPassword(),
			update:  createValidPassword(),
			assert: func(t *testing.T, expected account.Password, actual account.Password, err error) {
				errMsg := fmt.Sprintf("password change should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "change password failed",
			current: createValidPassword(),
			update:  createInvalidPassword(),
			assert: func(t *testing.T, expected account.Password, actual account.Password, err error) {
				errMsg := fmt.Sprintf("password change should fail: %s", err)
				require.NotEqual(t, expected, actual, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			password := account.Password(tc.current)
			newPassword, err := password.Update(tc.update)
			tc.assert(t, account.Password(tc.update), newPassword, err)
		})
	}
}

func TestPassword_String(t *testing.T) {
	password := account.Password(createValidPassword()).String()
	kind := reflect.TypeOf(password).String()
	require.Equal(t, "string", kind, "type should be `string`")
}
