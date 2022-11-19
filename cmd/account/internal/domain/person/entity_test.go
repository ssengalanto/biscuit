package person_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	entity := person.New()
	require.NotNilf(t, entity, "entity should not be nil")
}

func TestEntity_UpdateDetails(t *testing.T) {
	update := person.Details{
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		Email:       gofakeit.Email(),
		Phone:       gofakeit.Phone(),
		DateOfBirth: gofakeit.Date(),
	}
	invalidEmail := "invalid-email"

	testCases := []struct {
		name    string
		entity  person.Entity
		details person.UpdateDetailsInput
		assert  func(t *testing.T, expected person.Details, actual person.Details, err error)
	}{
		{
			name: "update person details success",
			entity: person.Entity{
				ID:        uuid.New(),
				AccountID: uuid.New(),
				Details: person.Details{
					FirstName:   gofakeit.FirstName(),
					LastName:    gofakeit.LastName(),
					Email:       gofakeit.Email(),
					Phone:       gofakeit.Phone(),
					DateOfBirth: gofakeit.Date(),
				},
				Avatar: gofakeit.URL(),
			},
			details: person.UpdateDetailsInput{
				FirstName:   &update.FirstName,
				LastName:    &update.LastName,
				Email:       &update.Email,
				Phone:       &update.Phone,
				DateOfBirth: &update.DateOfBirth,
			},
			assert: func(t *testing.T, expected person.Details, actual person.Details, err error) {
				errMsg := fmt.Sprintf("update person details should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name: "update person details failed",
			entity: person.Entity{
				ID:        uuid.New(),
				AccountID: uuid.New(),
				Details: person.Details{
					FirstName:   gofakeit.FirstName(),
					LastName:    gofakeit.LastName(),
					Email:       gofakeit.Email(),
					Phone:       gofakeit.Phone(),
					DateOfBirth: gofakeit.Date(),
				},
				Avatar: gofakeit.URL(),
			},
			details: person.UpdateDetailsInput{
				FirstName:   &update.FirstName,
				LastName:    &update.LastName,
				Email:       &invalidEmail,
				Phone:       &update.Phone,
				DateOfBirth: &update.DateOfBirth,
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
			entity := tc.entity
			err := entity.UpdateDetails(tc.details)
			details := entity.Details
			updateDetails := person.Details{
				FirstName:   *tc.details.FirstName,
				LastName:    *tc.details.LastName,
				Email:       *tc.details.Email,
				Phone:       *tc.details.Phone,
				DateOfBirth: *tc.details.DateOfBirth,
			}
			tc.assert(t,
				details, updateDetails, err)
		})
	}
}
