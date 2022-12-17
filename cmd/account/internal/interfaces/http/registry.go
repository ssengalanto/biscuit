package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
)

func RegisterHandlers(logger interfaces.Logger, router *chi.Mux) {
	httpHandlers := NewAccountHandler(logger)
	router.Post("/account", httpHandlers.CreateAccount)
}
