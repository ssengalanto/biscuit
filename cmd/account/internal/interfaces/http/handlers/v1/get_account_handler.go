//nolint:godot //unnecessary
package v1

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	qv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/query/v1"
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
	apphttp "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/http"
	"github.com/ssengalanto/biscuit/pkg/constants"
	"github.com/ssengalanto/biscuit/pkg/http/response/json"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"github.com/ssengalanto/midt"
)

// GetAccountHandler - http handler struct for account retrieval.
type GetAccountHandler struct {
	log      interfaces.Logger
	mediator *midt.Midt
}

// NewGetAccountHandler creates a new http handler for handling account retrieval.
func NewGetAccountHandler(logger interfaces.Logger, mediator *midt.Midt) *GetAccountHandler {
	return &GetAccountHandler{log: logger, mediator: mediator}
}

// Handle
// @Tags account
// @Summary Retrieve an existing account.
// @Description Retrieves an existing account record that matches the provided ID.
// @Accept json
// @Produce json
// @Param id path string true "Account ID" example("0b6ecded-fa9d-4b39-a309-9ef501de15f4")
// @Success 200 {object} GetAccountResponse
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/account/{id} [get]
func (c *GetAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.RequestTimeout)
	defer cancel()

	id := chi.URLParam(r, "id")

	request := dto.GetAccountRequest{ID: id}
	if !apphttp.ValidateRequest(w, c.log, request) {
		return
	}

	q := qv1.NewGetAccountQuery(request)

	response, err := c.mediator.Send(ctx, q)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	json.MustEncodeResponse(w, http.StatusOK, response)
}
