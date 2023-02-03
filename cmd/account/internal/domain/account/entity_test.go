package account_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/address"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/person"
	"github.com/ssengalanto/biscuit/cmd/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("it should create a new account", func(t *testing.T) {
		acct := account.New(
			gofakeit.Email(),
			gofakeit.Password(true, true, true, true, false, 10),
			gofakeit.Bool(),
		)
		assert.NotNil(t, acct)
	})
}

func TestEntity_IsActive(t *testing.T) {
	t.Run("it should return the account active state correctly", func(t *testing.T) {
		entity := mock.NewAccountEntity()

		entity.Deactivate()
		assert.False(t, entity.IsActive())

		entity.Activate()
		assert.True(t, entity.IsActive())
	})
}

func TestEntity_Activate(t *testing.T) {
	t.Run("it should activate the account", func(t *testing.T) {
		entity := mock.NewAccountEntity()
		entity.Activate()
		assert.True(t, entity.IsActive())
	})
}

func TestEntity_Deactivate(t *testing.T) {
	t.Run("it should deactivate the account", func(t *testing.T) {
		entity := mock.NewAccountEntity()
		entity.Deactivate()
		assert.False(t, entity.IsActive())
	})
}

func TestEntity_LoginTimestamp(t *testing.T) {
	t.Run("it should set the last login at with current timestamp", func(t *testing.T) {
		entity := mock.NewAccountEntity()
		entity.LoginTimestamp()
		assert.False(t, entity.LastLoginAt.IsZero())
	})
}

func TestEntity_UpdateEmail(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, expected account.Email, actual account.Email, err error)
	}{
		{
			name:    "it should update the email successfully",
			payload: gofakeit.Email(),
			assert: func(t *testing.T, expected account.Email, actual account.Email, err error) {
				assert.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should fail to update the email",
			payload: "invalid-email",
			assert: func(t *testing.T, expected account.Email, actual account.Email, err error) {
				assert.NotEqual(t, expected, actual)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entity := mock.NewAccountEntity()
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
			name:    "it should update the password successfully",
			payload: gofakeit.Password(true, true, true, true, false, 10),
			assert: func(t *testing.T, expected account.Password, actual account.Password, err error) {
				assert.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should fail to update the password",
			payload: gofakeit.Password(true, true, true, true, false, 5),
			assert: func(t *testing.T, expected account.Password, actual account.Password, err error) {
				assert.NotEqual(t, expected, actual)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entity := mock.NewAccountEntity()
			err := entity.UpdatePassword(tc.payload)
			tc.assert(t, entity.Password, account.Password(tc.payload), err)
		})
	}
}

func TestEntity_UpdateDetails(t *testing.T) {
	entity := mock.NewAccountEntity()
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
		entity  account.Entity
		details person.UpdateDetailsInput
		assert  func(t *testing.T, expected person.Details, actual person.Details, err error)
	}{
		{
			name:   "it should update person details successfully",
			entity: entity,
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
			name:   "it should fail to update person details",
			entity: entity,
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
			acct := tc.entity
			err := acct.UpdatePersonDetails(tc.details)
			details := acct.Person.Details
			updateDetails := person.Details{
				FirstName:   *tc.details.FirstName,
				LastName:    *tc.details.LastName,
				Email:       *tc.details.Email,
				Phone:       *tc.details.Phone,
				DateOfBirth: *tc.details.DateOfBirth,
			}
			tc.assert(t, details, updateDetails, err)
		})
	}
}

func TestEntity_UpdatePersonAvatar(t *testing.T) {
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
			entity := mock.NewAccountEntity()
			err := entity.UpdatePersonAvatar(tc.payload)
			tc.assert(t, entity.Person.Avatar, person.Avatar(tc.payload), err)
		})
	}
}

func TestEntity_UpdatePersonAddress(t *testing.T) {
	entity := mock.NewAccountEntity()
	addr := gofakeit.Address()
	testCases := []struct {
		name    string
		entity  account.Entity
		payload []account.UpdateAddressInput
		assert  func(t *testing.T, expected []address.Entity, actual []address.Entity, err error)
	}{
		{
			name:   "it should partially update the address successfully",
			entity: entity,
			payload: []account.UpdateAddressInput{
				{
					ID: (*entity.Person.Address)[0].ID.String(),
					Components: address.UpdateComponentsInput{
						Street:     &addr.Street,
						Unit:       nil,
						City:       &addr.City,
						District:   &addr.City,
						State:      &addr.State,
						Country:    &addr.Country,
						PostalCode: nil,
					},
				},
			},
			assert: func(t *testing.T, expected []address.Entity, actual []address.Entity, err error) {
				assert.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var addrs []address.Entity
			err := tc.entity.UpdatePersonAddress(tc.payload)
			for _, item := range *entity.Person.Address {
				for _, payload := range tc.payload {
					if item.ID.String() == payload.ID {
						addr := address.Entity{
							ID:         item.ID,
							PersonID:   item.PersonID,
							Components: item.Components,
						}
						addrs = append(addrs, addr)
					}
				}
			}
			tc.assert(t, addrs, *entity.Person.Address, err)
		})
	}
}

func TestEntity_HashPassword(t *testing.T) {
	t.Run("it should hash the password", func(t *testing.T) {
		entity := mock.NewAccountEntity()
		err := entity.HashPassword()
		require.NoError(t, err)
	})
}

func TestEntity_CheckPassword(t *testing.T) {
	t.Run("it should match the password", func(t *testing.T) {
		entity := mock.NewAccountEntity()
		pw := entity.Password.String()

		err := entity.HashPassword()

		require.NoError(t, err)

		match := entity.CheckPassword(pw)
		assert.True(t, match)
	})
}

func TestEntity_IsValid(t *testing.T) {
	testCases := []struct {
		name   string
		entity account.Entity
		assert func(t *testing.T, err error)
	}{
		{
			name:   "it should be a valid account",
			entity: mock.NewAccountEntity(),
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name:   "it should have an invalid email",
			entity: account.New("invalid", gofakeit.Password(true, true, true, true, false, 10), true),
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
		{
			name:   "it should have an invalid password",
			entity: account.New(gofakeit.Email(), "invalid", true),
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.entity.IsValid()
			tc.assert(t, err)
		})
	}
}

func TestAggregateAccount(t *testing.T) {
	t.Run("it should aggregate account, person and addresses", func(t *testing.T) {
		acct := mock.NewAccount()
		pers := mock.NewPerson()
		addrs := mock.NewAddresses()

		expected := acct
		expected.Person = &pers
		expected.Person.Address = &addrs

		entity := account.AggregateAccount(acct, pers, addrs)
		require.Equal(t, expected, entity)
	})
}
