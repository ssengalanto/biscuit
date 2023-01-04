package address_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	entity := newAddress()
	require.NotNilf(t, entity, "entity should not be nil")
}

func TestEntity_UpdateComponents(t *testing.T) {
	testCases := []struct {
		name   string
		assert func(t *testing.T)
	}{
		{
			name: "update address success",
			assert: func(t *testing.T) {
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

				entity := newAddress()
				err := entity.UpdateComponents(payload)
				errMsg := fmt.Sprintf("update address should succeed: %s", err)
				require.Equal(t, entity.Components, address.Components{
					Street:     *payload.Street,
					Unit:       *payload.Unit,
					City:       *payload.City,
					District:   *payload.District,
					State:      *payload.State,
					Country:    *payload.Country,
					PostalCode: *payload.PostalCode,
				}, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.assert(t)
		})
	}
}

func newAddress() address.Entity {
	addr := gofakeit.Address()
	return address.New(uuid.New(), address.Components{
		Street:     addr.Address,
		Unit:       addr.Address,
		City:       addr.City,
		District:   addr.City,
		State:      addr.State,
		Country:    addr.Country,
		PostalCode: addr.Zip,
	})
}
