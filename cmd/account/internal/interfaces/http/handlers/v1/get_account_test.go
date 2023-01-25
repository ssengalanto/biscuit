package v1_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	qv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/query/v1"
	v1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/http/handlers/v1"
	acctmock "github.com/ssengalanto/biscuit/cmd/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGetAccountHandler(t *testing.T) {
	t.Run("it should create a new get account http handler", func(t *testing.T) {
		logger, mediator := createDependencies(t)
		hdlr := v1.NewGetAccountHandler(logger, mediator)
		assert.NotNil(t, hdlr)
	})
}

func TestGetAccountHandler_Handle(t *testing.T) {
	logger, mediator := createDependencies(t)
	hdlr := v1.NewGetAccountHandler(logger, mediator)
	assert.NotNil(t, hdlr)

	mux := chi.NewRouter()
	mux.Get("/api/v1/accounts/{id}", hdlr.Handle)

	acct := acctmock.NewAccountEntity()
	id := acct.ID.String()
	url := fmt.Sprintf("/api/v1/accounts/%s", id)

	t.Run("it should return success response", func(t *testing.T) {
		mediator.EXPECT().Send(gomock.Any(), &qv1.GetAccountQuery{ID: id}).Times(1).Return(acct, nil)

		r, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return an error due to an invalid uuid", func(t *testing.T) {
		logger.EXPECT().Error(gomock.Any(), gomock.Any()).Times(1)

		invalid := fmt.Sprintf("/api/v1/accounts/%s", "invalid")
		r, err := http.NewRequest(http.MethodGet, invalid, nil)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("it should return an error due to get account query failure", func(t *testing.T) {
		mediator.EXPECT().Send(
			gomock.Any(),
			&qv1.GetAccountQuery{ID: id},
		).Times(1).Return(nil, errors.New("error"))

		r, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
