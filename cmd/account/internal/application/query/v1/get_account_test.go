package v1_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	v1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/query/v1"
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/biscuit/cmd/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGetAccountCommand(t *testing.T) {
	t.Run("it should create a new get account command instance", func(t *testing.T) {
		input := dto.GetAccountRequest{ID: gofakeit.UUID()}
		cmd := v1.NewGetAccountQuery(input)
		assert.NotNil(t, cmd)
	})
}

func TestNewGetAccountCommandHandler(t *testing.T) {
	t.Run("it should create a new get account handler instance", func(t *testing.T) {
		logger, repository, cache := createDepedencies(t)
		hdlr := v1.NewGetAccountQueryHandler(logger, repository, cache)
		assert.NotNil(t, hdlr)
	})
}

func TestGetAccountCommandHandler_Name(t *testing.T) {
	t.Run("it should return the correct handler name", func(t *testing.T) {
		logger, repository, cache := createDepedencies(t)
		hdlr := v1.NewGetAccountQueryHandler(logger, repository, cache)
		n := hdlr.Name()
		assert.Equal(t, fmt.Sprintf("%T", &v1.GetAccountQuery{}), n)
	})
}

func TestGetAccountCommandHandler_Handle(t *testing.T) {
	ctx := context.Background()
	logger, repository, cache := createDepedencies(t)
	hdlr := v1.NewGetAccountQueryHandler(logger, repository, cache)

	t.Run("it should return the correct response", func(t *testing.T) {
		entity := mock.NewAccountEntity()
		cache.EXPECT().Get(ctx, gomock.Any())
		repository.EXPECT().FindByID(ctx, gomock.Any()).Return(entity, nil)
		cache.EXPECT().Set(ctx, gomock.Any(), gomock.Any())
		res, err := hdlr.Handle(ctx, &v1.GetAccountQuery{ID: entity.ID.String()})
		require.NoError(t, err)

		r, ok := res.(dto.GetAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, entity.ID.String(), r.ID)
	})
	t.Run("it should return an error when an invalid uuid is provided", func(t *testing.T) {
		id := "invalid"
		cache.EXPECT().Get(ctx, gomock.Any())
		logger.EXPECT().Error(gomock.Any(), gomock.Any())
		res, err := hdlr.Handle(ctx, &v1.GetAccountQuery{ID: id})
		require.Error(t, err)

		r, ok := res.(dto.GetAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, "", r.ID)
	})
	t.Run("it should panic when an invalid query is provided", func(t *testing.T) {
		assert.Panics(t, func() {
			req := struct{}{}
			_, _ = hdlr.Handle(ctx, &req)
		})
	})
}
