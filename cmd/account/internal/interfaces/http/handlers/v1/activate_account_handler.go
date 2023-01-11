//nolint:godot //unnecessary
package v1

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	cmdv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	qv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/query/v1"
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
	apphttp "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/http"
	"github.com/ssengalanto/biscuit/pkg/constants"
	"github.com/ssengalanto/biscuit/pkg/http/response/json"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"github.com/ssengalanto/midt"
)

// ActivateAccountHandler - http handler struct for account activation.
type ActivateAccountHandler struct {
	log      interfaces.Logger
	mediator *midt.Midt
}

// NewActivateAccountHandler creates a new http handler for handling account activation.
func NewActivateAccountHandler(logger interfaces.Logger, mediator *midt.Midt) *ActivateAccountHandler {
	return &ActivateAccountHandler{log: logger, mediator: mediator}
}

// Handle
// @Tags account
// @Summary Activate an account
// @Description Activate an existing account record
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} dto.GetAccountResponseDto
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /api/v1/account/{id}/activate [patch]
func (a *ActivateAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.RequestTimeout)
	defer cancel()

	id := chi.URLParam(r, "id")

	actreq := dto.ActivateAccountRequestDto{ID: id}
	if !apphttp.ValidateRequest(w, a.log, actreq) {
		return
	}

	cmd := cmdv1.NewActivateAccountCommand(actreq)

	_, err := a.mediator.Send(ctx, cmd)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	getreq := dto.GetAccountRequestDto{ID: id}
	if !apphttp.ValidateRequest(w, a.log, getreq) {
		return
	}

	q := qv1.NewGetAccountQuery(getreq)

	response, err := a.mediator.Send(ctx, q)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	json.MustEncodeResponse(w, http.StatusOK, response)
}
