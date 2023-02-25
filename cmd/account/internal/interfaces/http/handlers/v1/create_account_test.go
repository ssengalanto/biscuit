package v1_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	cmdv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	qv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/query/v1"
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
	acctmock "github.com/ssengalanto/biscuit/cmd/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountHandler_Handle(t *testing.T) {
	s := setup(t)

	acct := acctmock.NewAccountEntity()
	id := acct.ID.String()
	url := "/api/v1/accounts"

	addr := gofakeit.Address()
	payload := dtov1.CreateAccountRequest{
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

	t.Run("it should return success response", func(t *testing.T) {
		s.mediator.EXPECT().Send(
			gomock.Any(),
			cmdv1.NewCreateAccountCommand(payload)).Times(1).Return(dtov1.CreateAccountResponse{ID: id},
			nil,
		)
		s.mediator.EXPECT().Send(gomock.Any(), &qv1.GetAccountQuery{ID: id}).Times(1).Return(acct, nil)

		r, err := http.NewRequest(http.MethodPost, url, parsePayload(payload))
		require.NoError(t, err)

		w := httptest.NewRecorder()
		s.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
	})
	t.Run("it should return an error due to invalid json format", func(t *testing.T) {
		s.logger.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()

		r, err := http.NewRequest(http.MethodPost, url, strings.NewReader("{}{}"))
		require.NoError(t, err)

		w := httptest.NewRecorder()
		s.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("it should return an error due to create account command failure", func(t *testing.T) {
		s.mediator.EXPECT().Send(
			gomock.Any(),
			cmdv1.NewCreateAccountCommand(payload)).Times(1).Return(nil, errors.New("error"))

		r, err := http.NewRequest(http.MethodPost, url, parsePayload(payload))
		require.NoError(t, err)

		w := httptest.NewRecorder()
		s.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
	t.Run("it should return an error due to get account query failure", func(t *testing.T) {
		s.mediator.EXPECT().Send(
			gomock.Any(),
			cmdv1.NewCreateAccountCommand(payload)).Times(1).Return(dtov1.CreateAccountResponse{ID: id},
			nil,
		)
		s.mediator.EXPECT().Send(
			gomock.Any(),
			&qv1.GetAccountQuery{ID: id}).Times(1).Return(nil, errors.New("error"))

		r, err := http.NewRequest(http.MethodPost, url, parsePayload(payload))
		require.NoError(t, err)

		w := httptest.NewRecorder()
		s.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
