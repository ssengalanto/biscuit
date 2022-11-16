package address_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	entity := address.New()
	require.NotNilf(t, entity, "entity should not be nil")
}

func TestEntity_Update(t *testing.T) {
	addr := gofakeit.Address()

	testCases := []struct {
		name    string
		payload address.UpdateAddressInput
		assert  func(t *testing.T, expected address.UpdateAddressInput, actual address.UpdateAddressInput, err error)
	}{
		{
			name: "update address success",
			payload: address.UpdateAddressInput{
				Components: address.Components{
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
				Geometry: address.Geometry{
					Lat: gofakeit.Latitude(),
					Lng: gofakeit.Longitude(),
				},
			},
			assert: func(t *testing.T, expected address.UpdateAddressInput, actual address.UpdateAddressInput, err error) {
				errMsg := fmt.Sprintf("update address should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name: "update address failed",
			payload: address.UpdateAddressInput{
				Components: address.Components{
					PlaceID: gofakeit.UUID(),
				},
				Geometry: address.Geometry{},
			},
			assert: func(t *testing.T, expected address.UpdateAddressInput, actual address.UpdateAddressInput, err error) {
				errMsg := fmt.Sprintf("update address should fail: %s", err)
				require.NotEqual(t, expected, actual, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entity := address.New()
			err := entity.Update(tc.payload)
			tc.assert(t,
				address.UpdateAddressInput{
					Components: entity.Components,
					Geometry:   entity.Geometry,
				}, tc.payload, err)
		})
	}
}
