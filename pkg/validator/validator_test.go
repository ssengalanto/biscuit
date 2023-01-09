package validator_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/biscuit/pkg/validator"
	"github.com/stretchr/testify/require"
)

func TestVar(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, err error)
	}{
		{
			name:    "validation passed",
			payload: "with value",
			assert: func(t *testing.T, err error) {
				require.Nil(t, err, fmt.Sprintf("validation should pass: %s", err))
			},
		},
		{
			name:    "validation failed",
			payload: "",
			assert: func(t *testing.T, err error) {
				require.NotNil(t, err, fmt.Sprintf("validation should fail: %s", err))
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.Var("test", tc.payload, "required")
			tc.assert(t, err)
		})
	}
}

func TestStruct(t *testing.T) {
	type user struct {
		FirstName string `validate:"required"`
		LastName  string `validate:"required"`
		Email     string `validate:"required,email"`
	}

	testCases := []struct {
		name    string
		payload user
		assert  func(t *testing.T, err error)
	}{
		{
			name: "validation passed",
			payload: user{
				FirstName: gofakeit.FirstName(),
				LastName:  gofakeit.LastName(),
				Email:     gofakeit.Email(),
			},
			assert: func(t *testing.T, err error) {
				require.Nil(t, err, fmt.Sprintf("validation should pass: %s", err))
			}},
		{
			name: "validation failed",
			payload: user{
				FirstName: gofakeit.FirstName(),
				LastName:  "",
				Email:     "invalid",
			},
			assert: func(t *testing.T, err error) {
				require.NotNil(t, err, fmt.Sprintf("validation should fail: %s", err))
			}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.Struct(tc.payload)
			tc.assert(t, err)
		})
	}
}
