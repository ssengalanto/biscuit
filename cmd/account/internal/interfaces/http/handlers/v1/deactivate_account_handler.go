//nolint:godot //unnecessary
package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	cmdv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	qv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/query/v1"
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
	apphttp "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/http"
	"github.com/ssengalanto/biscuit/pkg/http/response/json"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"github.com/ssengalanto/midt"
)

// DeactivateAccountHandler - http handler struct for account deactivation.
type DeactivateAccountHandler struct {
	log      interfaces.Logger
	mediator midt.Mediator
}

// NewDeactivateAccountHandler creates a new http handler for handling account deactivation.
func NewDeactivateAccountHandler(logger interfaces.Logger, mediator midt.Mediator) *DeactivateAccountHandler {
	return &DeactivateAccountHandler{log: logger, mediator: mediator}
}

// Handle
// @Tags account
// @Summary Deactivate an existing account.
// @Description Deactivate an existing account record that matches the provided ID.
// @Accept json
// @Produce json
// @Param id path string true "Account ID" example("0b6ecded-fa9d-4b39-a309-9ef501de15f4")
// @Success 200 {object} GetAccountResponse
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/accounts/{id}/deactivate [patch]
func (d *DeactivateAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")

	dreq := dtov1.DeactivateAccountRequest{ID: id}
	if !apphttp.ValidateRequest(w, d.log, dreq) {
		return
	}

	cmd := cmdv1.NewDeactivateAccountCommand(dreq)

	_, err := d.mediator.Send(ctx, cmd)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	greq := dtov1.GetAccountRequest{ID: id}
	if !apphttp.ValidateRequest(w, d.log, greq) {
		return
	}

	q := qv1.NewGetAccountQuery(greq)

	res, err := d.mediator.Send(ctx, q)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	json.MustEncodeResponse(w, http.StatusOK, res)
}
