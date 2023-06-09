//nolint:godot //unnecessary
package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	cmdv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	qv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/query/v1"
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
	apphttp "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/http"
	"github.com/ssengalanto/biscuit/pkg/errors"
	"github.com/ssengalanto/biscuit/pkg/http/response/json"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"github.com/ssengalanto/midt"
)

// UpdateAccountHandler - http handler struct for updating account.
type UpdateAccountHandler struct {
	log      interfaces.Logger
	mediator midt.Mediator
}

// NewUpdateAccountHandler creates a new http handler for handling account updates.
func NewUpdateAccountHandler(logger interfaces.Logger, mediator midt.Mediator) *UpdateAccountHandler {
	return &UpdateAccountHandler{log: logger, mediator: mediator}
}

// Handle
// @Tags account
// @Summary Update an existing account.
// @Description Updates an existing account in the database with the provided request body.
// @Accept json
// @Produce json
// @Param id path string true "Account ID" example("0b6ecded-fa9d-4b39-a309-9ef501de15f4")
// @Param UpdateAccountRequest body UpdateAccountRequest true "Account data"
// @Success 200 {object} GetAccountResponse
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/accounts/{id} [patch]
func (u *UpdateAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")

	var request dtov1.UpdateAccountRequest

	err := json.DecodeRequest(w, r, &request)
	if err != nil {
		u.log.Error("invalid request body format", map[string]any{"error": err})
		json.MustEncodeError(w, errors.ErrInvalid)
		return
	}

	if !apphttp.ValidateRequest(w, u.log, request) {
		return
	}

	cmd := cmdv1.NewUpdateAccountCommand(
		id,
		dtov1.UpdateAccountRequest{
			FirstName:   request.FirstName,
			LastName:    request.LastName,
			Phone:       request.Phone,
			DateOfBirth: request.DateOfBirth,
			Locations:   request.Locations,
		},
	)

	rr, err := u.mediator.Send(ctx, cmd)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	rsc := rr.(dtov1.UpdateAccountResponse) //nolint:errcheck //intentional panic

	q := qv1.NewGetAccountQuery(dtov1.GetAccountRequest{ID: rsc.ID}) //nolint:gosimple //explicit
	res, err := u.mediator.Send(ctx, q)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	json.MustEncodeResponse(w, http.StatusOK, res)
}
