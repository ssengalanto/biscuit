package app

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/ssengalanto/potato-project/cmd/account/docs" //notlint:revive //unnecessary
	cmdv1 "github.com/ssengalanto/potato-project/cmd/account/internal/application/command/v1"
	qv1 "github.com/ssengalanto/potato-project/cmd/account/internal/application/query/v1"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/http"
	httphv1 "github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/http/handlers/v1"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"github.com/ssengalanto/potato-project/pkg/mediatr"
	httpSwagger "github.com/swaggo/http-swagger"
)

// RegisterHTTPHandlers - http registry.
func RegisterHTTPHandlers(logger interfaces.Logger, mediator *mediatr.Mediatr) *chi.Mux {
	r := http.NewRouter()
	r.Mount("/swagger", httpSwagger.WrapHandler)

	// v1 routes
	r.Route("/api/v1", func(r chi.Router) {
		getAccountHandler := httphv1.NewGetAccountHandler(logger, mediator)
		r.Get("/account/{id}", getAccountHandler.Handle)

		createAccountHandler := httphv1.NewCreateAccountHandler(logger, mediator)
		r.Post("/account", createAccountHandler.Handle)

		deleteAccountHandler := httphv1.NewDeleteAccountHandler(logger, mediator)
		r.Delete("/account/{id}", deleteAccountHandler.Handle)
	})

	return r
}

// RegisterMediatrHandlers - mediatr registry.
func RegisterMediatrHandlers(logger interfaces.Logger, repository account.Repository) *mediatr.Mediatr {
	m := mediatr.New()

	// v1 commands
	createAccountCommandHandler := cmdv1.NewCreateAccountCommandHandler(logger, repository)
	err := m.RegisterRequestHandler(createAccountCommandHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}

	deleteAccountCommandHandler := cmdv1.NewDeleteAccountCommandHandler(logger, repository)
	err = m.RegisterRequestHandler(deleteAccountCommandHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}

	// v1 queries
	getAccountQueryHandler := qv1.NewGetAccountQueryHandler(logger, repository)
	err = m.RegisterRequestHandler(getAccountQueryHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}

	return m
}
