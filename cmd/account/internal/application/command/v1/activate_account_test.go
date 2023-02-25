package v1_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	cmdv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewActivateAccountCommand(t *testing.T) {
	t.Run("it should create a new activate account command instance", func(t *testing.T) {
		input := dtov1.ActivateAccountRequest{ID: gofakeit.UUID()}
		cmd := cmdv1.NewActivateAccountCommand(input)
		assert.NotNil(t, cmd)
	})
}

func TestNewActivateAccountCommandHandler(t *testing.T) {
	t.Run("it should create a new activate account handler instance", func(t *testing.T) {
		logger, repository, cache := createDependencies(t)
		hdlr := cmdv1.NewActivateAccountCommandHandler(logger, repository, cache)
		assert.NotNil(t, hdlr)
	})
}

func TestActivateAccountCommandHandler_Name(t *testing.T) {
	t.Run("it should return the correct handler name", func(t *testing.T) {
		logger, repository, cache := createDependencies(t)
		hdlr := cmdv1.NewActivateAccountCommandHandler(logger, repository, cache)
		n := hdlr.Name()
		assert.Equal(t, fmt.Sprintf("%T", &cmdv1.ActivateAccountCommand{}), n)
	})
}

func TestActivateAccountCommandHandler_Handle(t *testing.T) {
	ctx := context.Background()
	logger, repository, cache := createDependencies(t)
	hdlr := cmdv1.NewActivateAccountCommandHandler(logger, repository, cache)

	t.Run("it should return the correct response", func(t *testing.T) {
		id := uuid.New()
		repository.EXPECT().FindByID(ctx, gomock.Eq(id))
		repository.EXPECT().Update(ctx, gomock.Any())
		cache.EXPECT().Delete(ctx, gomock.Eq(id.String()))
		res, err := hdlr.Handle(ctx, &cmdv1.ActivateAccountCommand{ID: id.String()})
		require.NoError(t, err)

		r, ok := res.(dtov1.ActivateAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, id.String(), r.ID)
	})
	t.Run("it should return an error when an invalid uuid is provided", func(t *testing.T) {
		id := "invalid"
		logger.EXPECT().Error(gomock.Any(), gomock.Any())
		res, err := hdlr.Handle(ctx, &cmdv1.ActivateAccountCommand{ID: id})
		require.Error(t, err)

		r, ok := res.(dtov1.ActivateAccountResponse)
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
