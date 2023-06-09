package address_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/address"
	"github.com/ssengalanto/biscuit/cmd/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("it should create a new address", func(t *testing.T) {
		entity := mock.NewAddress()
		assert.NotNil(t, entity)
	})
}

func TestEntity_UpdateComponents(t *testing.T) {
	t.Run("it should update the address successfully", func(t *testing.T) {
		addr := gofakeit.Address()
		payload := address.UpdateComponentsInput{
			Street:     &addr.Address,
			Unit:       &addr.Address,
			City:       &addr.City,
			District:   &addr.City,
			State:      &addr.State,
			Country:    &addr.Country,
			PostalCode: &addr.Zip,
		}
		entity := mock.NewAddress()
		err := entity.UpdateComponents(payload)
		assert.Equal(t, entity.Components, address.Components{
			Street:     *payload.Street,
			Unit:       *payload.Unit,
			City:       *payload.City,
			District:   *payload.District,
			State:      *payload.State,
			Country:    *payload.Country,
			PostalCode: *payload.PostalCode,
		})
		require.NoError(t, err)
	})
}

func Test_IsValid(t *testing.T) {
	t.Run("it should be a valid address", func(t *testing.T) {
		entity := mock.NewAddress()
		err := entity.IsValid()
		require.NoError(t, err)
	})
	t.Run("it should be an invalid address", func(t *testing.T) {
		entity := mock.NewAddress()
		entity.Components.Street = ""
		err := entity.IsValid()
		require.Error(t, err)
	})
}
