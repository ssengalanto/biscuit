package v1_test

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	v1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/http/handlers/v1"
	"github.com/ssengalanto/biscuit/pkg/mock"
)

type testSetup struct {
	mux      *chi.Mux
	logger   *mock.MockLogger
	mediator *mock.MockMediator
}

func setup(t *testing.T) *testSetup {
	ctrl := gomock.NewController(t)
	logger := mock.NewMockLogger(ctrl)
	mediator := mock.NewMockMediator(ctrl)

	getAccountHandler := v1.NewGetAccountHandler(logger, mediator)
	createAccountHandler := v1.NewCreateAccountHandler(logger, mediator)
	activateAccountHandler := v1.NewActivateAccountHandler(logger, mediator)
	deactivateAccountHandler := v1.NewDeactivateAccountHandler(logger, mediator)
	updateAccountHandler := v1.NewUpdateAccountHandler(logger, mediator)
	deleteAccountHandler := v1.NewDeleteAccountHandler(logger, mediator)

	mux := chi.NewRouter()
	mux.Get("/api/v1/accounts/{id}", getAccountHandler.Handle)
	mux.Post("/api/v1/accounts", createAccountHandler.Handle)
	mux.Patch("/api/v1/accounts/{id}", updateAccountHandler.Handle)
	mux.Patch("/api/v1/accounts/{id}/activate", activateAccountHandler.Handle)
	mux.Patch("/api/v1/accounts/{id}/deactivate", deactivateAccountHandler.Handle)
	mux.Delete("/api/v1/accounts/{id}", deleteAccountHandler.Handle)

	return &testSetup{
		mux:      mux,
		logger:   logger,
		mediator: mediator,
	}
}

func parsePayload(payload any) io.Reader {
	res, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return strings.NewReader(string(res))
}
