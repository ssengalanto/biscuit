//nolint:godot //unnecessary
package v1

import (
	"net/http"

	cmdv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	qv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/query/v1"
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
	apphttp "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/http"
	"github.com/ssengalanto/biscuit/pkg/errors"
	"github.com/ssengalanto/biscuit/pkg/http/response/json"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"github.com/ssengalanto/midt"
)

// CreateAccountHandler - http handler struct for account creation.
type CreateAccountHandler struct {
	log      interfaces.Logger
	mediator midt.Mediator
}

// NewCreateAccountHandler creates a new http handler for handling account creation.
func NewCreateAccountHandler(logger interfaces.Logger, mediator midt.Mediator) *CreateAccountHandler {
	return &CreateAccountHandler{log: logger, mediator: mediator}
}

// Handle
// @Tags account
// @Summary Create a new account.
// @Description Creates a new account in the database with the provided request body.
// @Accept json
// @Produce json
// @Param CreateAccountRequest body CreateAccountRequest true "Account data"
// @Success 201 {object} GetAccountResponse
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/accounts [post]
func (c *CreateAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request dtov1.CreateAccountRequest

	err := json.DecodeRequest(w, r, &request)
	if err != nil {
		c.log.Error("invalid request body format", map[string]any{"error": err})
		json.MustEncodeError(w, errors.ErrInvalid)
		return
	}

	if !apphttp.ValidateRequest(w, c.log, request) {
		return
	}

	cmd := cmdv1.NewCreateAccountCommand(dtov1.CreateAccountRequest{
		Email:       request.Email,
		Password:    request.Password,
		Active:      request.Active,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Phone:       request.Phone,
		DateOfBirth: request.DateOfBirth,
		Locations:   request.Locations,
	})

	rr, err := c.mediator.Send(ctx, cmd)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	rsc := rr.(dtov1.CreateAccountResponse) //nolint:errcheck //intentional panic

	q := qv1.NewGetAccountQuery(dtov1.GetAccountRequest{ID: rsc.ID}) //nolint:gosimple //explicit
	res, err := c.mediator.Send(ctx, q)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	json.MustEncodeResponse(w, http.StatusCreated, res)
}
