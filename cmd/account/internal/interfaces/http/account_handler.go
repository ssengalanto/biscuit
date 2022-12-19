package http

import (
	"context"
	"net/http"

	"github.com/ssengalanto/potato-project/cmd/account/internal/application/command"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/ssengalanto/potato-project/pkg/http/response/json"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"github.com/ssengalanto/potato-project/pkg/mediatr"
)

type AccountHandler struct {
	log      interfaces.Logger
	mediator *mediatr.Mediatr
}

// NewAccountHandler creates a new account handler.
func NewAccountHandler(logger interfaces.Logger, mediator *mediatr.Mediatr) *AccountHandler {
	return &AccountHandler{log: logger, mediator: mediator}
}

func (a *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.RequestTimeout)
	defer cancel()

	var request dto.CreateAccountRequestDto

	err := json.DecodeRequest(w, r, &request)
	if err != nil {
		a.log.Error("invalid request", map[string]any{"error": err})
		json.MustEncodeError(w, errors.ErrInvalid)
		return
	}

	cmd := command.NewCreateAccountCommand(dto.CreateAccountRequestDto{
		Email:       request.Email,
		Password:    request.Password,
		Active:      request.Active,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Phone:       request.Phone,
		DateOfBirth: request.DateOfBirth,
	})

	response, err := a.mediator.Send(ctx, cmd)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	json.MustEncodeResponse(w, http.StatusCreated, response)
}
