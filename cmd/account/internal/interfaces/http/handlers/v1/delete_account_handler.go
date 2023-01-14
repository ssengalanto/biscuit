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

// DeleteAccountHandler - http handler struct for account deletion.
type DeleteAccountHandler struct {
	log      interfaces.Logger
	mediator *midt.Midt
}

// NewDeleteAccountHandler creates a new http handler for handling account deletion.
func NewDeleteAccountHandler(logger interfaces.Logger, mediator *midt.Midt) *DeleteAccountHandler {
	return &DeleteAccountHandler{log: logger, mediator: mediator}
}

// Handle
// @Tags account
// @Summary Delete an account
// @Description Delete an existing account record
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} dto.GetAccountResponse
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /api/v1/account/{id} [delete]
func (c *DeleteAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.RequestTimeout)
	defer cancel()

	id := chi.URLParam(r, "id")

	getreq := dto.GetAccountRequest{ID: id}
	if !apphttp.ValidateRequest(w, c.log, getreq) {
		return
	}

	q := qv1.NewGetAccountQuery(getreq)

	response, err := c.mediator.Send(ctx, q)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	delreq := dto.DeleteAccountRequest{ID: id}
	if !apphttp.ValidateRequest(w, c.log, delreq) {
		return
	}

	cmd := cmdv1.NewDeleteAccountCommand(delreq)

	_, err = c.mediator.Send(ctx, cmd)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	json.MustEncodeResponse(w, http.StatusOK, response)
}
