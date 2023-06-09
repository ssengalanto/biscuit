package v1_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	cmdv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCreateAccountCommand(t *testing.T) {
	t.Run("it should create a new create account command instance", func(t *testing.T) {
		input := newCreateAccountRequest()
		cmd := cmdv1.NewCreateAccountCommand(input)
		assert.NotNil(t, cmd)
	})
}

func TestNewCreateAccountCommandHandler(t *testing.T) {
	t.Run("it should create a new create account handler instance", func(t *testing.T) {
		logger, repository, cache := createDependencies(t)
		hdlr := cmdv1.NewCreateAccountCommandHandler(logger, repository, cache)
		assert.NotNil(t, hdlr)
	})
}

func TestCreateAccountCommandHandler_Name(t *testing.T) {
	t.Run("it should return the correct handler name", func(t *testing.T) {
		logger, repository, cache := createDependencies(t)
		hdlr := cmdv1.NewCreateAccountCommandHandler(logger, repository, cache)
		n := hdlr.Name()
		assert.Equal(t, fmt.Sprintf("%T", &cmdv1.CreateAccountCommand{}), n)
	})
}

func TestCreateAccountCommandHandler_Handle(t *testing.T) {
	ctx := context.Background()
	logger, repository, cache := createDependencies(t)
	hdlr := cmdv1.NewCreateAccountCommandHandler(logger, repository, cache)
	t.Run("it should return the correct response", func(t *testing.T) {
		req := newCreateAccountRequest()
		repository.EXPECT().Create(ctx, gomock.Any())
		res, err := hdlr.Handle(ctx, &cmdv1.CreateAccountCommand{
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

		_, ok := res.(dtov1.CreateAccountResponse)
		assert.True(t, ok)
	})
	t.Run("it should return an error when an invalid account is provided", func(t *testing.T) {
		req := newCreateAccountRequest()
		logger.EXPECT().Error(gomock.Any(), gomock.Any())
		res, err := hdlr.Handle(ctx, &cmdv1.CreateAccountCommand{
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

		r, ok := res.(dtov1.CreateAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, "", r.ID)
	})
	t.Run("it should return an error when an invalid person is provided", func(t *testing.T) {
		req := newCreateAccountRequest()
		logger.EXPECT().Error(gomock.Any(), gomock.Any())
		res, err := hdlr.Handle(ctx, &cmdv1.CreateAccountCommand{
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

		r, ok := res.(dtov1.CreateAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, "", r.ID)
	})
	t.Run("it should panic when an invalid command is provided", func(t *testing.T) {
		assert.Panics(t, func() {
			req := struct{}{}
			_, _ = hdlr.Handle(ctx, &req)
		})
	})
}

func newCreateAccountRequest() dtov1.CreateAccountRequest {
	addr := gofakeit.Address()
	return dtov1.CreateAccountRequest{
		Email:       gofakeit.Email(),
		Password:    gofakeit.Password(true, true, true, true, false, 10),
		Active:      true,
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		Phone:       gofakeit.Phone(),
		DateOfBirth: gofakeit.Date(),
		Locations: []dtov1.CreateAddressRequest{{
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
