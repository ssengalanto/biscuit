package person_test

import (
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/person"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDetails_IsValid(t *testing.T) {
	testCases := []struct {
		name    string
		payload person.Details
		assert  func(t *testing.T, err error)
	}{
		{
			name: "it should be a valid person details",
			payload: person.Details{
				FirstName:   gofakeit.FirstName(),
				LastName:    gofakeit.LastName(),
				Email:       gofakeit.Email(),
				Phone:       gofakeit.Phone(),
				DateOfBirth: gofakeit.Date(),
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name:    "it should be an invalid person details",
			payload: person.Details{},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			details := tc.payload
			err := details.IsValid()
			tc.assert(t, err)
		})
	}
}

func TestDetails_Update(t *testing.T) {
	testCases := []struct {
		name    string
		current person.Details
		update  person.Details
		assert  func(t *testing.T, expected person.Details, actual person.Details, err error)
	}{
		{
			name:    "it should update the person details successfully",
			current: createPersonDetails(),
			update:  createPersonDetails(),
			assert: func(t *testing.T, expected person.Details, actual person.Details, err error) {
				require.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should fail to update person details",
			current: createPersonDetails(),
			update: person.Details{
				FirstName:   gofakeit.FirstName(),
				LastName:    gofakeit.LastName(),
				Email:       "invalid-email",
				Phone:       gofakeit.Phone(),
				DateOfBirth: gofakeit.Date(),
			},
			assert: func(t *testing.T, expected person.Details, actual person.Details, err error) {
				assert.NotEqual(t, expected, actual)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			details := tc.current
			newDetails, err := details.Update(tc.update)
			tc.assert(t,
				newDetails, tc.update, err)
		})
	}
}

func TestAvatar_IsValid(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name:    "it should be a valid avatar",
			payload: gofakeit.URL(),
			assert: func(t *testing.T, result bool, err error) {
				require.True(t, result)
				require.Nil(t, err)
			},
		},
		{
			name:    "it should be an invalid avatar",
			payload: "invalid-avatar",
			assert: func(t *testing.T, result bool, err error) {
				require.False(t, result)
				require.NotNil(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			avatar := person.Avatar(tc.payload)
			ok, err := avatar.IsValid()
			tc.assert(t, ok, err)
		})
	}
}

func TestAvatar_Update(t *testing.T) {
	testCases := []struct {
		name    string
		current string
		update  string
		assert  func(t *testing.T, expected person.Avatar, actual person.Avatar, err error)
	}{
		{
			name:    "it should change the avatar successfully",
			current: gofakeit.URL(),
			update:  gofakeit.URL(),
			assert: func(t *testing.T, expected person.Avatar, actual person.Avatar, err error) {
				require.Equal(t, expected, actual)
				require.Nil(t, err)
			},
		},
		{
			name:    "it should fail to change the avatar",
			current: gofakeit.URL(),
			update:  "invalid-avatar",
			assert: func(t *testing.T, expected person.Avatar, actual person.Avatar, err error) {
				require.NotEqual(t, expected, actual)
				require.NotNil(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			avatar := person.Avatar(tc.current)
			newAvatar, err := avatar.Update(tc.update)
			tc.assert(t, person.Avatar(tc.update), newAvatar, err)
		})
	}
}

func TestAvatar_String(t *testing.T) {
	t.Run("it should convert avatar to string", func(t *testing.T) {
		avatar := person.Avatar(gofakeit.URL()).String()
		kind := reflect.TypeOf(avatar).String()
		require.Equal(t, "string", kind)
	})
}

func createPersonDetails() person.Details {
	return person.Details{
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		Email:       gofakeit.Email(),
		Phone:       gofakeit.Phone(),
		DateOfBirth: gofakeit.Date(),
	}
}
