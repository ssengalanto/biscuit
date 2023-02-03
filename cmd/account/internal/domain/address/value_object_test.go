package address_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/address"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComponents_IsValid(t *testing.T) {
	t.Run("it should be a valid address components", func(t *testing.T) {
		addr := gofakeit.Address()
		components := address.Components{
			Street:     addr.Address,
			Unit:       addr.Address,
			City:       addr.City,
			District:   addr.City,
			State:      addr.State,
			Country:    addr.Country,
			PostalCode: addr.Zip,
		}
		err := components.IsValid()
		require.NoError(t, err)
	})
	t.Run("it should be an invalid address components", func(t *testing.T) {
		components := address.Components{}
		err := components.IsValid()
		assert.NotNil(t, err)
		require.Error(t, err)
	})
}

func TestAddress_Update(t *testing.T) {
	t.Run("it should update the address components successfully", func(t *testing.T) {
		current := gofakeit.Address()
		components := address.Components{
			Street:     current.Address,
			Unit:       current.Address,
			City:       current.City,
			District:   current.City,
			State:      current.State,
			Country:    current.Country,
			PostalCode: current.Zip,
		}

		update := gofakeit.Address()
		payload := address.Components{
			Street:     update.Address,
			Unit:       update.Address,
			City:       update.City,
			District:   update.City,
			State:      update.State,
			Country:    update.Country,
			PostalCode: update.Zip,
		}

		newComponents, err := components.Update(payload)
		assert.Equal(t, payload, newComponents)
		require.NoError(t, err)
	})
	t.Run("it should fail to update the address components", func(t *testing.T) {
		current := gofakeit.Address()
		components := address.Components{
			Street:     current.Address,
			Unit:       current.Address,
			City:       current.City,
			District:   current.City,
			State:      current.State,
			Country:    current.Country,
			PostalCode: current.Zip,
		}

		update := gofakeit.Address()
		payload := address.Components{
			Street:     "",
			Unit:       update.Address,
			City:       update.City,
			District:   update.City,
			State:      update.State,
			Country:    update.Country,
			PostalCode: update.Zip,
		}

		newComponents, err := components.Update(payload)
		assert.Equal(t, address.Components{}, newComponents)
		require.Error(t, err)
	})
}
