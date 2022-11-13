package account_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	entity := account.New()
	require.NotNilf(t, entity, "entity should not be nil")
}

func TestEntity_IsActive(t *testing.T) {
	entity := account.New()
	require.False(t, entity.IsActive(), "Entity.Active should be false")
}

func TestEntity_Activate(t *testing.T) {
	entity := account.New()
	entity.Activate()
	require.True(t, entity.IsActive(), "Entity.Active should be true")
}

func TestEntity_Deactivate(t *testing.T) {
	entity := account.New()
	entity.Deactivate()
	require.False(t, entity.IsActive(), "Entity.Active should be false")
}

func TestEntity_LoginTimestamp(t *testing.T) {
	entity := account.New()
	entity.LoginTimestamp()
	require.False(t, entity.LastLoginAt.IsZero(), "Entity.LastLoginAt should not have zero value")
}

func TestEntity_UpdateEmail(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, expected account.Email, actual account.Email, err error)
	}{
		{
			name:    "update email success",
			payload: gofakeit.Email(),
			assert: func(t *testing.T, expected account.Email, actual account.Email, err error) {
				errMsg := fmt.Sprintf("update email should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "update email failed",
			payload: "invalid-email",
			assert: func(t *testing.T, expected account.Email, actual account.Email, err error) {
				errMsg := fmt.Sprintf("update email should fail: %s", err)
				require.NotEqual(t, expected, actual, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entity := account.New()
			err := entity.UpdateEmail(tc.payload)
			tc.assert(t, entity.Email, account.Email(tc.payload), err)
		})
	}
}

func TestEntity_UpdatePassword(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, expected account.Password, actual account.Password, err error)
	}{
		{
			name:    "update password success",
			payload: gofakeit.Password(true, true, true, true, false, 10),
			assert: func(t *testing.T, expected account.Password, actual account.Password, err error) {
				errMsg := fmt.Sprintf("update password should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "update password failed",
			payload: gofakeit.Password(true, true, true, true, false, 5),
			assert: func(t *testing.T, expected account.Password, actual account.Password, err error) {
				errMsg := fmt.Sprintf("update password should fail: %s", err)
				require.NotEqual(t, expected, actual, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entity := account.New()
			err := entity.UpdatePassword(tc.payload)
			tc.assert(t, entity.Password, account.Password(tc.payload), err)
		})
	}
}
