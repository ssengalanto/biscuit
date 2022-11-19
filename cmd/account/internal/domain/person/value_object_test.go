package person_test

import (
	"fmt"
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
