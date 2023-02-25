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

func TestNewDeleteAccountCommand(t *testing.T) {
	t.Run("it should create a new delete account command instance", func(t *testing.T) {
		input := dtov1.DeleteAccountRequest{ID: gofakeit.UUID()}
		cmd := cmdv1.NewDeleteAccountCommand(input)
		assert.NotNil(t, cmd)
	})
}

func TestNewDeleteAccountCommandHandler(t *testing.T) {
	t.Run("it should create a new delete account handler instance", func(t *testing.T) {
		logger, repository, cache := createDependencies(t)
		hdlr := cmdv1.NewDeleteAccountCommandHandler(logger, repository, cache)
		assert.NotNil(t, hdlr)
	})
}

func TestDeleteAccountCommandHandler_Name(t *testing.T) {
	t.Run("it should return the correct handler name", func(t *testing.T) {
		logger, repository, cache := createDependencies(t)
		hdlr := cmdv1.NewDeleteAccountCommandHandler(logger, repository, cache)
		n := hdlr.Name()
		assert.Equal(t, fmt.Sprintf("%T", &cmdv1.DeleteAccountCommand{}), n)
	})
}

func TestDeleteAccountCommandHandler_Handle(t *testing.T) {
	ctx := context.Background()
	logger, repository, cache := createDependencies(t)
	hdlr := cmdv1.NewDeleteAccountCommandHandler(logger, repository, cache)

	t.Run("it should return the correct response", func(t *testing.T) {
		id := uuid.New()
		repository.EXPECT().DeleteByID(ctx, gomock.Eq(id))
		cache.EXPECT().Delete(ctx, gomock.Eq(id.String()))
		res, err := hdlr.Handle(ctx, &cmdv1.DeleteAccountCommand{ID: id.String()})
		require.NoError(t, err)

		r, ok := res.(dtov1.DeleteAccountResponse)
		assert.True(t, ok)
		assert.Equal(t, id.String(), r.ID)
	})
	t.Run("it should return an error when an invalid uuid is provided", func(t *testing.T) {
		id := "invalid"
		logger.EXPECT().Error(gomock.Any(), gomock.Any())
		res, err := hdlr.Handle(ctx, &cmdv1.DeleteAccountCommand{ID: id})
		require.Error(t, err)

		r, ok := res.(dtov1.DeleteAccountResponse)
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
