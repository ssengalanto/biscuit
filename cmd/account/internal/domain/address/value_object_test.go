package address_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
	"github.com/stretchr/testify/require"
)

func TestComponents_IsValid(t *testing.T) {
	addr := gofakeit.Address()
	testCases := []struct {
		name    string
		payload address.Components
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name: "valid address components",
			payload: address.Components{
				Street:     addr.Address,
				Unit:       addr.Address,
				City:       addr.City,
				District:   addr.City,
				State:      addr.State,
				Country:    addr.Country,
				PostalCode: addr.Zip,
			},
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("address components should be valid: %s", err)
				require.True(t, result, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "invalid address components",
			payload: address.Components{},
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("address components should be invalid: %s", err)
				require.False(t, result, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			components := tc.payload
			ok, err := components.IsValid()
			tc.assert(t, ok, err)
		})
	}
}

func TestAddress_Update(t *testing.T) {
	current := gofakeit.Address()
	update := gofakeit.Address()

	currentAddr := address.Components{
		Street:     current.Address,
		Unit:       current.Address,
		City:       current.City,
		District:   current.City,
		State:      current.State,
		Country:    current.Country,
		PostalCode: current.Zip,
	}

	testCases := []struct {
		name    string
		current address.Components
		update  address.Components
		assert  func(t *testing.T, expected address.Components, actual address.Components, err error)
	}{
		{
			name:    "update address components success",
			current: currentAddr,
			update: address.Components{
				Street:     update.Address,
				Unit:       update.Address,
				City:       update.City,
				District:   update.City,
				State:      update.State,
				Country:    update.Country,
				PostalCode: update.Zip,
			},
			assert: func(t *testing.T, expected address.Components, actual address.Components, err error) {
				errMsg := fmt.Sprintf("update address components should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			components := tc.current
			newComponents, err := components.Update(tc.update)
			tc.assert(t, tc.update, newComponents, err)
		})
	}
}
