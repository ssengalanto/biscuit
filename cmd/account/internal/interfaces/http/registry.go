package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"github.com/ssengalanto/potato-project/pkg/mediatr"
)

func RegisterHandlers(logger interfaces.Logger, router *chi.Mux, mediator *mediatr.Mediatr) {
	httpHandlers := NewAccountHandler(logger, mediator)
	router.Post("/account", httpHandlers.CreateAccount)
}
