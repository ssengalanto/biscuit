package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	_ "github.com/ssengalanto/biscuit/cmd/account/docs" //notlint:revive //unnecessary
	cmdv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	qv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/query/v1"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
	cache "github.com/ssengalanto/biscuit/cmd/account/internal/infrastructure/cache/redis"
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/http"
	httphv1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/http/handlers/v1"
	"github.com/ssengalanto/biscuit/pkg/behaviours"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"github.com/ssengalanto/midt"
	httpSwagger "github.com/swaggo/http-swagger"
)

// registerHTTPHandlers registers all http handlers in router.
func registerHTTPHandlers(logger interfaces.Logger, mediator *midt.Midt) *chi.Mux {
	r := http.NewRouter()
	r.Mount("/swagger/docs", httpSwagger.WrapHandler)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			getAccountHandler := httphv1.NewGetAccountHandler(logger, mediator)
			r.Get("/account/{id}", getAccountHandler.Handle)

			createAccountHandler := httphv1.NewCreateAccountHandler(logger, mediator)
			r.Post("/account", createAccountHandler.Handle)

			updateAccountHandler := httphv1.NewUpdateAccountHandler(logger, mediator)
			r.Patch("/account/{id}", updateAccountHandler.Handle)

			activateAccountHandler := httphv1.NewActivateAccountHandler(logger, mediator)
			r.Patch("/account/{id}/activate", activateAccountHandler.Handle)

			deactivateAccountHandler := httphv1.NewDeactivateAccountHandler(logger, mediator)
			r.Patch("/account/{id}/deactivate", deactivateAccountHandler.Handle)

			deleteAccountHandler := httphv1.NewDeleteAccountHandler(logger, mediator)
			r.Delete("/account/{id}", deleteAccountHandler.Handle)
		})
	})

	return r
}

// registerMediatorHandlers registers all request, notification and pipeline behaviour handlers in the registry.
func registerMediatorHandlers(
	logger interfaces.Logger,
	repository account.Repository,
	rdb redis.UniversalClient,
) *midt.Midt {
	m := midt.New()
	c := cache.New(logger, rdb)

	registerRequestHandlers(logger, repository, c, m)
	registerPipelineBehaviours(logger, m)

	return m
}

// registerRequestHandlers registers all request handlers in the registry.
func registerRequestHandlers(
	logger interfaces.Logger,
	repository account.Repository,
	cache account.Cache,
	m *midt.Midt,
) {
	// v1 commands
	createAccountCommandHandler := cmdv1.NewCreateAccountCommandHandler(logger, repository, cache)
	err := m.RegisterRequestHandler(createAccountCommandHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}

	updateAccountCommandHandler := cmdv1.NewUpdateAccountCommandHandler(logger, repository, cache)
	err = m.RegisterRequestHandler(updateAccountCommandHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}

	activateAccountCommandHandler := cmdv1.NewActivateAccountCommandHandler(logger, repository, cache)
	err = m.RegisterRequestHandler(activateAccountCommandHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}

	deactivateAccountCommandHandler := cmdv1.NewDeactivateAccountCommandHandler(logger, repository, cache)
	err = m.RegisterRequestHandler(deactivateAccountCommandHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}

	deleteAccountCommandHandler := cmdv1.NewDeleteAccountCommandHandler(logger, repository, cache)
	err = m.RegisterRequestHandler(deleteAccountCommandHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}

	// v1 queries
	getAccountQueryHandler := qv1.NewGetAccountQueryHandler(logger, repository, cache)
	err = m.RegisterRequestHandler(getAccountQueryHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}
}

// registerPipelineBehaviours registers all pipeline behaviour in the registry.
func registerPipelineBehaviours(logger interfaces.Logger, m *midt.Midt) {
	loggerBehaviour := behaviours.NewLoggerBehaviour(logger)
	err := m.RegisterPipelineBehaviour(loggerBehaviour)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}
}
