//nolint:godot //unnecessary
package http

import (
	"context"
	"net/http"

	"github.com/ssengalanto/potato-project/cmd/account/internal/application/command"
	"github.com/ssengalanto/potato-project/cmd/account/internal/application/query"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/ssengalanto/potato-project/pkg/http/response/json"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"github.com/ssengalanto/potato-project/pkg/mediatr"
)

type CreateAccountHandler struct {
	log      interfaces.Logger
	mediator *mediatr.Mediatr
}

// NewCreateAccountHandler creates a new http handler for handling account creation.
func NewCreateAccountHandler(logger interfaces.Logger, mediator *mediatr.Mediatr) *CreateAccountHandler {
	return &CreateAccountHandler{log: logger, mediator: mediator}
}

// Handle
// @Tags account
// @Summary Create a new account
// @Description Creates a new account
// @Accept json
// @Produce json
// @Param CreateAccountRequestDto body dto.CreateAccountRequestDto true "Account data"
// @Success 201 {object} dto.CreateAccountResponseDto
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /api/v1/account [post]
func (c *CreateAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.RequestTimeout)
	defer cancel()

	var request dto.CreateAccountRequestDto

	err := json.DecodeRequest(w, r, &request)
	if err != nil {
		c.log.Error("invalid request", map[string]any{"error": err})
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

	resource, err := c.mediator.Send(ctx, cmd)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	resourceID, ok := resource.(string)
	if !ok {
		c.log.Error("invalid resource id", map[string]any{"id": resourceID})
		json.MustEncodeError(w, errors.ErrInvalid)
	}

	q := query.NewGetAccountQuery(dto.GetAccountRequestDto{ID: resourceID})
	response, err := c.mediator.Send(ctx, q)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	json.MustEncodeResponse(w, http.StatusCreated, response)
}
