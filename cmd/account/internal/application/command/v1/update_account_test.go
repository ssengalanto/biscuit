package v1_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	v1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/biscuit/cmd/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUpdateAccountCommand(t *testing.T) {
	t.Parallel()
	t.Run("it should create a new update account command instance", func(t *testing.T) {
		t.Parallel()
		cmd := v1.NewUpdateAccountCommand(gofakeit.UUID(), dto.UpdateAccountRequest{})
		assert.NotNil(t, cmd)
	})
}

func TestNewUpdateAccountCommandHandler(t *testing.T) {
	t.Run("it should create a new update account handler instance", func(t *testing.T) {
		logger, repository, cache := createDependencies(t)
		hdlr := v1.NewUpdateAccountCommandHandler(logger, repository, cache)
		assert.NotNil(t, hdlr)
	})
}

func TestUpdateAccountCommandHandler_Name(t *testing.T) {
	t.Run("it should return the correct handler name", func(t *testing.T) {
		logger, repository, cache := createDependencies(t)
		hdlr := v1.NewUpdateAccountCommandHandler(logger, repository, cache)
		n := hdlr.Name()
		assert.Equal(t, fmt.Sprintf("%T", &v1.UpdateAccountCommand{}), n)
	})
}

func TestUpdateAccountCommandHandler_Handle(t *testing.T) {
	ctx := context.Background()
	logger, repository, cache := createDependencies(t)
	hdlr := v1.NewUpdateAccountCommandHandler(logger, repository, cache)

	t.Run("it should return the correct response", func(t *testing.T) {
		entity, req := newUpdateAccountRequest()
		repository.EXPECT().FindByID(ctx, gomock.Any()).Return(entity, nil)
		repository.EXPECT().Update(ctx, gomock.Any())
		cache.EXPECT().Delete(ctx, gomock.Any())
		res, err := hdlr.Handle(ctx, &v1.UpdateAccountCommand{
			ID:          entity.ID.String(),
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			Phone:       req.Phone,
			DateOfBirth: req.DateOfBirth,
			Locations:   req.Locations,
		})
		require.NoError(t, err)

		r, ok := res.(dto.UpdateAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, entity.ID.String(), r.ID)
	})
	t.Run("it should return an error when an invalid uuid is provided", func(t *testing.T) {
		id := "invalid"
		logger.EXPECT().Error(gomock.Any(), gomock.Any())
		res, err := hdlr.Handle(ctx, &v1.UpdateAccountCommand{ID: id})
		require.Error(t, err)

		r, ok := res.(dto.UpdateAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, "", r.ID)
	})
	t.Run("it should return an error when an invalid command is provided", func(t *testing.T) {
		req := struct{}{}
		logger.EXPECT().Error(gomock.Any(), gomock.Any())
		res, err := hdlr.Handle(ctx, &req)
		require.Error(t, err)

		r, ok := res.(dto.UpdateAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, "", r.ID)
	})
}

func newUpdateAccountRequest() (account.Entity, dto.UpdateAccountRequest) {
	var loc []dto.UpdateAddressRequest
	entity := mock.NewAccountEntity()
	fn := gofakeit.FirstName()
	ln := gofakeit.LastName()
	phone := gofakeit.Phone()
	dob := gofakeit.Date()

	for _, addr := range *entity.Person.Address {
		a := gofakeit.Address()
		req := dto.UpdateAddressRequest{
			ID:         addr.ID.String(),
			Street:     &a.Street,
			Unit:       &a.Street,
			City:       &a.City,
			District:   &a.City,
			State:      &a.State,
			Country:    &a.Country,
			PostalCode: &a.Zip,
		}
		loc = append(loc, req)
	}

	return entity, dto.UpdateAccountRequest{
		FirstName:   &fn,
		LastName:    &ln,
		Phone:       &phone,
		DateOfBirth: &dob,
		Locations:   &loc,
	}
}
