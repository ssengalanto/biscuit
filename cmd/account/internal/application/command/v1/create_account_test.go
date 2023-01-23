package v1_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	v1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCreateAccountCommand(t *testing.T) {
	t.Parallel()
	t.Run("it should create a new create account command instance", func(t *testing.T) {
		t.Parallel()
		input := newCreateAccountRequest()
		cmd := v1.NewCreateAccountCommand(input)
		assert.NotNil(t, cmd)
	})
}

func TestNewCreateAccountCommandHandler(t *testing.T) {
	t.Parallel()
	t.Run("it should create a new create account handler instance", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		logger, repository, cache := createDepedencies(ctrl)
		hdlr := v1.NewCreateAccountCommandHandler(logger, repository, cache)
		assert.NotNil(t, hdlr)
	})
}

func TestCreateAccountCommandHandler_Name(t *testing.T) {
	t.Parallel()
	t.Run("it should return the correct handler name", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		logger, repository, cache := createDepedencies(ctrl)
		hdlr := v1.NewCreateAccountCommandHandler(logger, repository, cache)
		n := hdlr.Name()
		assert.Equal(t, fmt.Sprintf("%T", &v1.CreateAccountCommand{}), n)
	})
}

func TestCreateAccountCommandHandler_Handle(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	logger, repository, cache := createDepedencies(ctrl)
	hdlr := v1.NewCreateAccountCommandHandler(logger, repository, cache)

	t.Run("it should return the correct response", func(t *testing.T) {
		req := newCreateAccountRequest()
		repository.EXPECT().Create(ctx, gomock.Any())
		res, err := hdlr.Handle(ctx, &v1.CreateAccountCommand{
			Email:       req.Email,
			Password:    req.Password,
			Active:      req.Active,
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			Phone:       req.Phone,
			DateOfBirth: req.DateOfBirth,
			Locations:   req.Locations,
		})
		require.NoError(t, err)

		_, ok := res.(dto.CreateAccountResponse)
		assert.True(t, ok)
	})
	t.Run("it should return an error when an invalid command is provided", func(t *testing.T) {
		t.Parallel()
		req := struct{}{}
		logger.EXPECT().Error(gomock.Any(), gomock.Any())
		res, err := hdlr.Handle(ctx, &req)
		require.Error(t, err)

		r, ok := res.(dto.CreateAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, "", r.ID)
	})
	t.Run("it should return an error when an invalid account is provided", func(t *testing.T) {
		t.Parallel()
		req := newCreateAccountRequest()
		logger.EXPECT().Error(gomock.Any(), gomock.Any())
		res, err := hdlr.Handle(ctx, &v1.CreateAccountCommand{
			Email:       "",
			Password:    req.Password,
			Active:      req.Active,
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			Phone:       req.Phone,
			DateOfBirth: req.DateOfBirth,
			Locations:   req.Locations,
		})
		require.Error(t, err)

		r, ok := res.(dto.CreateAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, "", r.ID)
	})
	t.Run("it should return an error when an invalid person is provided", func(t *testing.T) {
		t.Parallel()
		req := newCreateAccountRequest()
		logger.EXPECT().Error(gomock.Any(), gomock.Any())
		res, err := hdlr.Handle(ctx, &v1.CreateAccountCommand{
			Email:       req.Email,
			Password:    req.Password,
			Active:      req.Active,
			FirstName:   "",
			LastName:    req.LastName,
			Phone:       req.Phone,
			DateOfBirth: req.DateOfBirth,
			Locations:   req.Locations,
		})
		require.Error(t, err)

		r, ok := res.(dto.CreateAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, "", r.ID)
	})
}

func newCreateAccountRequest() dto.CreateAccountRequest {
	addr := gofakeit.Address()
	return dto.CreateAccountRequest{
		Email:       gofakeit.Email(),
		Password:    gofakeit.Password(true, true, true, true, false, 10),
		Active:      true,
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		Phone:       gofakeit.Phone(),
		DateOfBirth: gofakeit.Date(),
		Locations: []dto.CreateAddressRequest{{
			Street:     addr.Street,
			Unit:       addr.Street,
			City:       addr.City,
			District:   addr.City,
			State:      addr.State,
			Country:    addr.Country,
			PostalCode: addr.Zip,
		}},
	}
}
