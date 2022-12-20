//nolint:godot //unnecessary
package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ssengalanto/potato-project/cmd/account/internal/application/query"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/http/response/json"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"github.com/ssengalanto/potato-project/pkg/mediatr"
)

type GetAccountHandler struct {
	log      interfaces.Logger
	mediator *mediatr.Mediatr
}

// NewGetAccountHandler creates a new http handler for handling account retrieval.
func NewGetAccountHandler(logger interfaces.Logger, mediator *mediatr.Mediatr) *GetAccountHandler {
	return &GetAccountHandler{log: logger, mediator: mediator}
}

// Handle
// @Tags account
// @Summary Get account by ID
// @Description Get account record by account ID.
// @Accept json
// @Produce json
// @Param GetAccountRequestDto body dto.GetAccountRequestDto true "Account Id"
// @Success 201 {object} dto.GetAccountResponseDto
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /api/v1/account/:id [get]
func (c *GetAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.RequestTimeout)
	defer cancel()

	id := chi.URLParam(r, "id")

	request := dto.GetAccountRequestDto{ID: id}

	q := query.NewGetAccountQuery(request)

	response, err := c.mediator.Send(ctx, q)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	json.MustEncodeResponse(w, http.StatusOK, response)
}
