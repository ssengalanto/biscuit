package http

import (
	"net/http"

	"github.com/ssengalanto/biscuit/pkg/errors"
	"github.com/ssengalanto/biscuit/pkg/http/response/json"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"github.com/ssengalanto/biscuit/pkg/validator"
)

func ValidateRequest(w http.ResponseWriter, log interfaces.Logger, req any) bool {
	err := validator.Struct(req)
	if err != nil {
		log.Error("invalid request", map[string]any{"error": err})
		json.MustEncodeError(w, errors.ErrInvalid)
		return false
	}
	return true
}
