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

func TestEntity_Update(t *testing.T) {
	addr := gofakeit.Address()

	testCases := []struct {
		name    string
		payload address.Components
		assert  func(t *testing.T, expected address.Components, actual address.Components, err error)
	}{
		{
			name: "update address success",
			payload: address.Components{
				Street:     addr.Address,
				Unit:       addr.Address,
				City:       addr.City,
				District:   addr.City,
				State:      addr.State,
				Country:    addr.Country,
				PostalCode: addr.Zip,
			},
			assert: func(t *testing.T, expected address.Components, actual address.Components, err error) {
				errMsg := fmt.Sprintf("update address should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name: "update address failed",
			payload: address.Components{
				PostalCode: addr.Zip,
			},
			assert: func(t *testing.T, expected address.Components, actual address.Components, err error) {
				errMsg := fmt.Sprintf("update address should fail: %s", err)
				require.NotEqual(t, expected, actual, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entity := newAddress()
			err := entity.Update(tc.payload)
			tc.assert(t, entity.Components, tc.payload, err)
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
