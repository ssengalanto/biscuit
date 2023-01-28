package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	authHttp "github.com/ssengalanto/biscuit/cmd/auth/internal/interfaces/http"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"github.com/ssengalanto/midt"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

// registerHTTPHandlers registers all http handlers in router.
func registerHTTPHandlers(_ interfaces.Logger, _ *midt.Midt) *chi.Mux {
	r := authHttp.NewRouter()
	// TODO: add path to swagger docs
	r.Mount("/swagger/docs", httpSwagger.WrapHandler)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/sample", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				var message Message
				err := json.NewDecoder(r.Body).Decode(&message)
				if err != nil {
					return
				}
				err = json.NewEncoder(w).Encode(message)
				if err != nil {
					return
				}
			})
		})
	})

	return r
}

// registerMediatorHandlers registers all request, notification and pipeline behaviour handlers in the registry.
func registerMediatorHandlers(
	_ interfaces.Logger,
) *midt.Midt {
	m := midt.New()

	return m
}
