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
				PlaceID: gofakeit.UUID(),
				AddressLine1: address.Names{
					ShortName: addr.Street,
					LongName:  addr.Street,
				},
				AddressLine2: address.Names{
					ShortName: addr.Street,
					LongName:  addr.Street,
				},
				City: address.Names{
					ShortName: addr.City,
					LongName:  addr.City,
				},
				State: address.Names{
					ShortName: addr.State,
					LongName:  addr.State,
				},
				Country: address.Names{
					ShortName: addr.Country,
					LongName:  addr.Country,
				},
				PostalCode: address.Names{
					ShortName: addr.Zip,
					LongName:  addr.Zip,
				},
				FormattedAddress: addr.Address,
			},
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("address components should be valid: %s", err)
				require.True(t, result, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name: "invalid address components",
			payload: address.Components{
				PlaceID:          "",
				AddressLine1:     address.Names{},
				AddressLine2:     address.Names{},
				City:             address.Names{},
				State:            address.Names{},
				Country:          address.Names{},
				PostalCode:       address.Names{},
				FormattedAddress: "",
			},
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

func TestGeometry_IsValid(t *testing.T) {
	testCases := []struct {
		name    string
		payload address.Geometry
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name: "valid address geometry",
			payload: address.Geometry{
				Lat: gofakeit.Latitude(),
				Lng: gofakeit.Longitude(),
			},
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("address geometry should be valid: %s", err)
				require.True(t, result, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name: "invalid address geometry",
			payload: address.Geometry{
				Lat: 0,
				Lng: 0,
			},
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("address geometry should be invalid: %s", err)
				require.False(t, result, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			geometry := tc.payload
			ok, err := geometry.IsValid()
			tc.assert(t, ok, err)
		})
	}
}
