package person_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
	"github.com/stretchr/testify/require"
)

func TestDetails_IsValid(t *testing.T) {
	testCases := []struct {
		name    string
		payload person.Details
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name: "valid person details",
			payload: person.Details{
				FirstName:   gofakeit.FirstName(),
				LastName:    gofakeit.LastName(),
				Email:       gofakeit.Email(),
				Phone:       gofakeit.Phone(),
				DateOfBirth: gofakeit.Date(),
			},
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("person details should be valid: %s", err)
				require.True(t, result, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name: "invalid person details",
			payload: person.Details{
				FirstName:   "",
				LastName:    "",
				Email:       "",
				Phone:       "",
				DateOfBirth: time.Time{},
			},
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("person details should be invalid: %s", err)
				require.False(t, result, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			details := tc.payload
			ok, err := details.IsValid()
			tc.assert(t, ok, err)
		})
	}
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

func TestDetails_Update(t *testing.T) {
	testCases := []struct {
		name    string
		current person.Details
		update  person.Details
		assert  func(t *testing.T, expected person.Details, actual person.Details, err error)
	}{
		{
			name:    "update person details success",
			current: createPersonDetails(),
			update:  createPersonDetails(),
			assert: func(t *testing.T, expected person.Details, actual person.Details, err error) {
				errMsg := fmt.Sprintf("update person details should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "update person details failed",
			current: createPersonDetails(),
			update: person.Details{
				FirstName:   gofakeit.FirstName(),
				LastName:    gofakeit.LastName(),
				Email:       "invalid-email",
				Phone:       gofakeit.Phone(),
				DateOfBirth: gofakeit.Date(),
			},
			assert: func(t *testing.T, expected person.Details, actual person.Details, err error) {
				errMsg := fmt.Sprintf("update person details should fail: %s", err)
				require.NotEqual(t, expected, actual, errMsg)
				require.NotNil(t, err, errMsg)
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
			name:    "valid avatar",
			payload: gofakeit.URL(),
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("avatar should be valid: %s", err)
				require.True(t, result, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "invalid avatar",
			payload: "invalid-avatar",
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("avatar should be invalid: %s", err)
				require.False(t, result, errMsg)
				require.NotNil(t, err, errMsg)
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
			name:    "change avatar success",
			current: gofakeit.URL(),
			update:  gofakeit.URL(),
			assert: func(t *testing.T, expected person.Avatar, actual person.Avatar, err error) {
				errMsg := fmt.Sprintf("avatar change should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "change avatar failed",
			current: gofakeit.URL(),
			update:  "invalid-avatar",
			assert: func(t *testing.T, expected person.Avatar, actual person.Avatar, err error) {
				errMsg := fmt.Sprintf("avatar change should fail: %s", err)
				require.NotEqual(t, expected, actual, errMsg)
				require.NotNil(t, err, errMsg)
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
	avatar := person.Avatar(gofakeit.URL()).String()
	kind := reflect.TypeOf(avatar).String()
	require.Equal(t, "string", kind, "type should be `string`")
}
