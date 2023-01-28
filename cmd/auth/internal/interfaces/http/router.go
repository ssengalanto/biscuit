package http

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/ssengalanto/biscuit/pkg/constants"
)

// NewRouter creates a new http router.
func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/heartbeat"))
	router.Use(httprate.LimitByIP(constants.RateLimit, time.Minute))
	router.Use(middleware.Timeout(constants.Timeout))

	return router
}
