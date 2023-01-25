package v1_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	cmdv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	qv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/query/v1"
	acctmock "github.com/ssengalanto/biscuit/cmd/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeactivateAccountHandler_Handle(t *testing.T) {
	s := setup(t)

	acct := acctmock.NewAccountEntity()
	id := acct.ID.String()
	url := fmt.Sprintf("/api/v1/accounts/%s/deactivate", id)

	t.Run("it should return success response", func(t *testing.T) {
		s.mediator.EXPECT().Send(gomock.Any(), &cmdv1.DeactivateAccountCommand{ID: id}).Times(1)
		s.mediator.EXPECT().Send(gomock.Any(), &qv1.GetAccountQuery{ID: id}).Times(1).Return(acct, nil)

		r, err := http.NewRequest(http.MethodPatch, url, nil)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		s.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return an error due to an invalid uuid", func(t *testing.T) {
		s.logger.EXPECT().Error(gomock.Any(), gomock.Any()).Times(1)

		invalid := fmt.Sprintf("/api/v1/accounts/%s/deactivate", "invalid")
		r, err := http.NewRequest(http.MethodPatch, invalid, nil)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		s.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("it should return an error due to deactivate account command failure", func(t *testing.T) {
		s.mediator.EXPECT().Send(
			gomock.Any(),
			&cmdv1.DeactivateAccountCommand{ID: id},
		).Times(1).Return(nil, errors.New("error"))

		r, err := http.NewRequest(http.MethodPatch, url, nil)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		s.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
	t.Run("it should return an error due to get account query failure", func(t *testing.T) {
		s.mediator.EXPECT().Send(gomock.Any(), &cmdv1.DeactivateAccountCommand{ID: id}).Times(1)
		s.mediator.EXPECT().Send(
			gomock.Any(),
			&qv1.GetAccountQuery{ID: id},
		).Times(1).Return(nil, errors.New("error"))

		r, err := http.NewRequest(http.MethodPatch, url, nil)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		s.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
