package v1_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	cmdv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
	"github.com/ssengalanto/biscuit/cmd/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUpdateAccountCommand(t *testing.T) {
	t.Run("it should create a new update account command instance", func(t *testing.T) {
		cmd := cmdv1.NewUpdateAccountCommand(gofakeit.UUID(), dtov1.UpdateAccountRequest{})
		assert.NotNil(t, cmd)
	})
}

func TestNewUpdateAccountCommandHandler(t *testing.T) {
	t.Run("it should create a new update account handler instance", func(t *testing.T) {
		logger, repository, cache := createDependencies(t)
		hdlr := cmdv1.NewUpdateAccountCommandHandler(logger, repository, cache)
		assert.NotNil(t, hdlr)
	})
}

func TestUpdateAccountCommandHandler_Name(t *testing.T) {
	t.Run("it should return the correct handler name", func(t *testing.T) {
		logger, repository, cache := createDependencies(t)
		hdlr := cmdv1.NewUpdateAccountCommandHandler(logger, repository, cache)
		n := hdlr.Name()
		assert.Equal(t, fmt.Sprintf("%T", &cmdv1.UpdateAccountCommand{}), n)
	})
}

func TestUpdateAccountCommandHandler_Handle(t *testing.T) {
	ctx := context.Background()
	logger, repository, cache := createDependencies(t)
	hdlr := cmdv1.NewUpdateAccountCommandHandler(logger, repository, cache)

	t.Run("it should return the correct response", func(t *testing.T) {
		entity, req := newUpdateAccountRequest()
		repository.EXPECT().FindByID(ctx, gomock.Any()).Return(entity, nil)
		repository.EXPECT().Update(ctx, gomock.Any())
		cache.EXPECT().Delete(ctx, gomock.Any())
		res, err := hdlr.Handle(ctx, &cmdv1.UpdateAccountCommand{
			ID:          entity.ID.String(),
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			Phone:       req.Phone,
			DateOfBirth: req.DateOfBirth,
			Locations:   req.Locations,
		})
		require.NoError(t, err)

		r, ok := res.(dtov1.UpdateAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, entity.ID.String(), r.ID)
	})
	t.Run("it should return an error when an invalid uuid is provided", func(t *testing.T) {
		id := "invalid"
		logger.EXPECT().Error(gomock.Any(), gomock.Any())
		res, err := hdlr.Handle(ctx, &cmdv1.UpdateAccountCommand{ID: id})
		require.Error(t, err)

		r, ok := res.(dtov1.UpdateAccountResponse)
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

func newUpdateAccountRequest() (account.Entity, dtov1.UpdateAccountRequest) {
	var loc []dtov1.UpdateAddressRequest
	entity := mock.NewAccountEntity()
	fn := gofakeit.FirstName()
	ln := gofakeit.LastName()
	phone := gofakeit.Phone()
	dob := gofakeit.Date()

	for _, addr := range *entity.Person.Address {
		a := gofakeit.Address()
		req := dtov1.UpdateAddressRequest{
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

	return entity, dtov1.UpdateAccountRequest{
		FirstName:   &fn,
		LastName:    &ln,
		Phone:       &phone,
		DateOfBirth: &dob,
		Locations:   &loc,
	}
}
