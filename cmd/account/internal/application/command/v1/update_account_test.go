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

func TestNewUpdateAccountCommand(t *testing.T) {
	t.Parallel()
	t.Run("it should create a new update account command instance", func(t *testing.T) {
		t.Parallel()
		req := newUpdateAccountRequest()
		cmd := v1.NewUpdateAccountCommand(gofakeit.UUID(), req)
		assert.NotNil(t, cmd)
	})
}

func TestNewUpdateAccountCommandHandler(t *testing.T) {
	t.Run("it should create a new update account handler instance", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger, repository, cache := createDepedencies(ctrl)
		hdlr := v1.NewUpdateAccountCommandHandler(logger, repository, cache)
		assert.NotNil(t, hdlr)
	})
}

func TestUpdateAccountCommandHandler_Name(t *testing.T) {
	t.Run("it should return the correct handler name", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger, repository, cache := createDepedencies(ctrl)
		hdlr := v1.NewUpdateAccountCommandHandler(logger, repository, cache)
		n := hdlr.Name()
		assert.Equal(t, fmt.Sprintf("%T", &v1.UpdateAccountCommand{}), n)
	})
}

func TestUpdateAccountCommandHandler_Handle(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	logger, repository, cache := createDepedencies(ctrl)
	hdlr := v1.NewUpdateAccountCommandHandler(logger, repository, cache)

	// t.Run("it should return the correct response", func(t *testing.T) {
	//	req := newUpdateAccountRequest()
	//	repository.EXPECT().FindByID(ctx, gomock.Any())
	//	repository.EXPECT().Update(ctx, gomock.Any())
	//	cache.EXPECT().Delete(ctx, gomock.Any())
	//	res, err := hdlr.Handle(ctx, &v1.UpdateAccountCommand{
	//		ID:          gofakeit.UUID(),
	//		FirstName:   req.FirstName,
	//		LastName:    req.LastName,
	//		Phone:       req.Phone,
	//		DateOfBirth: req.DateOfBirth,
	//	})
	//	require.NoError(t, err)
	//
	//	_, ok := res.(dto.UpdateAccountResponse)
	//	assert.True(t, ok)
	// })
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

func newUpdateAccountRequest() dto.UpdateAccountRequest {
	fn := gofakeit.FirstName()
	ln := gofakeit.LastName()
	phone := gofakeit.Phone()
	dob := gofakeit.Date()
	loc := []dto.UpdateAddressRequest{}
	return dto.UpdateAccountRequest{
		FirstName:   &fn,
		LastName:    &ln,
		Phone:       &phone,
		DateOfBirth: &dob,
		Locations:   &loc,
	}
}
