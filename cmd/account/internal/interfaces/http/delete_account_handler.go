//nolint:godot //unnecessary
package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ssengalanto/potato-project/cmd/account/internal/application/command"
	"github.com/ssengalanto/potato-project/cmd/account/internal/application/query"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/http/response/json"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"github.com/ssengalanto/potato-project/pkg/mediatr"
)

// DeleteAccountHandler - http handler struct for account deletion.
type DeleteAccountHandler struct {
	log      interfaces.Logger
	mediator *mediatr.Mediatr
}

// NewDeleteAccountHandler creates a new http handler for handling account deletion.
func NewDeleteAccountHandler(logger interfaces.Logger, mediator *mediatr.Mediatr) *DeleteAccountHandler {
	return &DeleteAccountHandler{log: logger, mediator: mediator}
}

// Handle
// @Tags account
// @Summary Delete an account
// @Description Delete an existing account record
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} dto.GetAccountResponseDto
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /api/v1/account/{id} [delete]
func (c *DeleteAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.RequestTimeout)
	defer cancel()

	id := chi.URLParam(r, "id")

	getreq := dto.GetAccountRequestDto{ID: id}
	if !validateRequest(w, c.log, getreq) {
		return
	}

	q := query.NewGetAccountQuery(getreq)

	response, err := c.mediator.Send(ctx, q)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	delreq := dto.DeleteAccountRequestDto{ID: id}
	if !validateRequest(w, c.log, delreq) {
		return
	}

	cmd := command.NewDeleteAccountCommand(delreq)

	_, err = c.mediator.Send(ctx, cmd)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	json.MustEncodeResponse(w, http.StatusCreated, response)
}
