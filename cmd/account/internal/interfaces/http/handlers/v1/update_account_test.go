package v1_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	cmdv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	qv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/query/v1"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/biscuit/cmd/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateAccountHandler_Handle(t *testing.T) {
	s := setup(t)

	acct, payload := newUpdateAccountRequest()
	id := acct.ID.String()
	url := fmt.Sprintf("/api/v1/accounts/%s", id)

	t.Run("it should return success response", func(t *testing.T) {
		s.mediator.EXPECT().Send(
			gomock.Any(),
			cmdv1.NewUpdateAccountCommand(id, payload)).Times(1).Return(dto.UpdateAccountResponse{ID: id}, nil)
		s.mediator.EXPECT().Send(gomock.Any(), &qv1.GetAccountQuery{ID: id}).Times(1).Return(acct, nil)

		r, err := http.NewRequest(http.MethodPatch, url, parsePayload(payload))
		require.NoError(t, err)

		w := httptest.NewRecorder()
		s.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return an error due to invalid json format", func(t *testing.T) {
		s.logger.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()

		r, err := http.NewRequest(http.MethodPatch, url, strings.NewReader("{}{}"))
		require.NoError(t, err)

		w := httptest.NewRecorder()
		s.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("it should return an error due to update account command failure", func(t *testing.T) {
		s.mediator.EXPECT().Send(
			gomock.Any(),
			cmdv1.NewUpdateAccountCommand(id, payload)).Times(1).Return(nil, errors.New("error"))

		r, err := http.NewRequest(http.MethodPatch, url, parsePayload(payload))
		require.NoError(t, err)

		w := httptest.NewRecorder()
		s.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
	t.Run("it should return an error due to get account query failure", func(t *testing.T) {
		s.mediator.EXPECT().Send(
			gomock.Any(),
			cmdv1.NewUpdateAccountCommand(id, payload)).Times(1).Return(dto.UpdateAccountResponse{ID: id}, nil)
		s.mediator.EXPECT().Send(
			gomock.Any(),
			&qv1.GetAccountQuery{ID: id}).Times(1).Return(nil, errors.New("error"))

		r, err := http.NewRequest(http.MethodPatch, url, parsePayload(payload))
		require.NoError(t, err)

		w := httptest.NewRecorder()
		s.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
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
