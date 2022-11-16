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
		PlaceID: gofakeit.UUID(),
		AddressLine1: address.Names{
			ShortName: current.Street,
			LongName:  current.Street,
		},
		AddressLine2: address.Names{
			ShortName: current.Street,
			LongName:  current.Street,
		},
		City: address.Names{
			ShortName: current.City,
			LongName:  current.City,
		},
		State: address.Names{
			ShortName: current.State,
			LongName:  current.State,
		},
		Country: address.Names{
			ShortName: current.Country,
			LongName:  current.Country,
		},
		PostalCode: address.Names{
			ShortName: current.Zip,
			LongName:  current.Zip,
		},
		FormattedAddress: current.Address,
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
				PlaceID: gofakeit.UUID(),
				AddressLine1: address.Names{
					ShortName: update.Street,
					LongName:  update.Street,
				},
				AddressLine2: address.Names{
					ShortName: update.Street,
					LongName:  update.Street,
				},
				City: address.Names{
					ShortName: update.City,
					LongName:  update.City,
				},
				State: address.Names{
					ShortName: update.State,
					LongName:  update.State,
				},
				Country: address.Names{
					ShortName: update.Country,
					LongName:  update.Country,
				},
				PostalCode: address.Names{
					ShortName: update.Zip,
					LongName:  update.Zip,
				},
				FormattedAddress: update.Address,
			},
			assert: func(t *testing.T, expected address.Components, actual address.Components, err error) {
				errMsg := fmt.Sprintf("update address components should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "update address components failed",
			current: currentAddr,
			update: address.Components{
				PlaceID: gofakeit.UUID(),
			},
			assert: func(t *testing.T, expected address.Components, actual address.Components, err error) {
				errMsg := fmt.Sprintf("update address components should fail: %s", err)
				require.NotEqual(t, expected, actual, errMsg)
				require.NotNil(t, err, errMsg)
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
			name:    "invalid address geometry",
			payload: address.Geometry{},
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

func TestGeometry_Update(t *testing.T) {
	current := address.Geometry{
		Lat: gofakeit.Latitude(),
		Lng: gofakeit.Longitude(),
	}
	update := address.Geometry{
		Lat: gofakeit.Latitude(),
		Lng: gofakeit.Longitude(),
	}

	testCases := []struct {
		name    string
		current address.Geometry
		update  address.Geometry
		assert  func(t *testing.T, expected address.Geometry, actual address.Geometry, err error)
	}{
		{
			name:    "update address geometry success",
			current: current,
			update:  update,
			assert: func(t *testing.T, expected address.Geometry, actual address.Geometry, err error) {
				errMsg := fmt.Sprintf("update address geometry should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "update address geometry failed",
			current: current,
			update: address.Geometry{
				Lat: gofakeit.Longitude(),
				Lng: 0,
			},
			assert: func(t *testing.T, expected address.Geometry, actual address.Geometry, err error) {
				errMsg := fmt.Sprintf("update address geometry should fail: %s", err)
				require.NotEqual(t, expected, actual, errMsg)
				require.NotNil(t, err, errMsg)
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
