package http

import (
	"net/http"

	"github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/ssengalanto/potato-project/pkg/http/response/json"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"github.com/ssengalanto/potato-project/pkg/validator"
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
