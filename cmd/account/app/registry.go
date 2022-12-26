package app

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/ssengalanto/potato-project/cmd/account/docs" //notlint:revive //unnecessary
	"github.com/ssengalanto/potato-project/cmd/account/internal/application/command"
	"github.com/ssengalanto/potato-project/cmd/account/internal/application/query"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/http"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"github.com/ssengalanto/potato-project/pkg/mediatr"
	httpSwagger "github.com/swaggo/http-swagger"
)

// RegisterHTTPHandlers - http registry.
func RegisterHTTPHandlers(logger interfaces.Logger, mediator *mediatr.Mediatr) *chi.Mux {
	r := http.NewRouter()
	r.Mount("/swagger", httpSwagger.WrapHandler)

	getAccountHandler := http.NewGetAccountHandler(logger, mediator)
	r.Get("/account/{id}", getAccountHandler.Handle)

	createAccountHandler := http.NewCreateAccountHandler(logger, mediator)
	r.Post("/account", createAccountHandler.Handle)

	deleteAccountHandler := http.NewDeleteAccountHandler(logger, mediator)
	r.Delete("/account/{id}", deleteAccountHandler.Handle)

	return r
}

// RegisterMediatrHandlers - mediatr registry.
func RegisterMediatrHandlers(logger interfaces.Logger, repository account.Repository) *mediatr.Mediatr {
	m := mediatr.New()

	// commands
	createAccountCommandHandler := command.NewCreateAccountCommandHandler(logger, repository)
	err := m.RegisterRequestHandler(createAccountCommandHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}

	deleteAccountCommandHandler := command.NewDeleteAccountCommandHandler(logger, repository)
	err = m.RegisterRequestHandler(deleteAccountCommandHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}

	// queries
	getAccountQueryHandler := query.NewGetAccountQueryHandler(logger, repository)
	err = m.RegisterRequestHandler(getAccountQueryHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}

	return m
}
