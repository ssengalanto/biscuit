package http

import (
	"github.com/ssengalanto/biscuit/pkg/constants"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

// NewRouter creates a new http router.
func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/heartbeat"))
	router.Use(httprate.LimitByIP(constants.RateLimit, time.Minute))
	router.Use(middleware.Recoverer)

	return router
}
