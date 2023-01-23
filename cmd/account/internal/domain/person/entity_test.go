package person_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/person"
	"github.com/ssengalanto/biscuit/cmd/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("it should create a new person", func(t *testing.T) {
		t.Parallel()
		entity := mock.NewPerson()
		assert.NotNil(t, entity)
	})
}

func TestEntity_UpdateDetails(t *testing.T) {
	t.Parallel()
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
			name: "it should update the person details successfully",
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
				Avatar: person.Avatar(gofakeit.URL()),
			},
			details: person.UpdateDetailsInput{
				FirstName:   &update.FirstName,
				LastName:    &update.LastName,
				Email:       &update.Email,
				Phone:       &update.Phone,
				DateOfBirth: &update.DateOfBirth,
			},
			assert: func(t *testing.T, expected person.Details, actual person.Details, err error) {
				assert.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
		{
			name: "it should fail to update the person details",
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
				Avatar: person.Avatar(gofakeit.URL()),
			},
			details: person.UpdateDetailsInput{
				FirstName:   &update.FirstName,
				LastName:    &update.LastName,
				Email:       &invalidEmail,
				Phone:       &update.Phone,
				DateOfBirth: &update.DateOfBirth,
			},
			assert: func(t *testing.T, expected person.Details, actual person.Details, err error) {
				assert.NotEqual(t, expected, actual)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
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

func TestEntity_UpdateAvatar(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, expected person.Avatar, actual person.Avatar, err error)
	}{
		{
			name:    "it should update the avatar successfully",
			payload: gofakeit.URL(),
			assert: func(t *testing.T, expected person.Avatar, actual person.Avatar, err error) {
				assert.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should fail to update the avatar",
			payload: "invalid-avatar",
			assert: func(t *testing.T, expected person.Avatar, actual person.Avatar, err error) {
				assert.NotEqual(t, expected, actual)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			entity := mock.NewPerson()
			err := entity.UpdateAvatar(tc.payload)
			tc.assert(t, entity.Avatar, person.Avatar(tc.payload), err)
		})
	}
}

func Test_IsValid(t *testing.T) {
	t.Parallel()
	t.Run("it should be a valid person", func(t *testing.T) {
		t.Parallel()
		entity := mock.NewPerson()
		err := entity.IsValid()
		require.NoError(t, err)
	})
	t.Run("it should be an invalid person", func(t *testing.T) {
		t.Parallel()
		entity := mock.NewPerson()
		entity.Details.FirstName = ""
		err := entity.IsValid()
		require.Error(t, err)
	})
}
