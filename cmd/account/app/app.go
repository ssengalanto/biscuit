package app

import (
	"log"

	_ "github.com/ssengalanto/potato-project/cmd/account/docs" //notlint:revive //unnecessary
	"github.com/ssengalanto/potato-project/cmd/account/internal/application/command"
	"github.com/ssengalanto/potato-project/cmd/account/internal/application/query"
	repository "github.com/ssengalanto/potato-project/cmd/account/internal/infrastructure/persistence/pgsql"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/http"
	"github.com/ssengalanto/potato-project/pkg/config"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/logger"
	"github.com/ssengalanto/potato-project/pkg/mediatr"
	"github.com/ssengalanto/potato-project/pkg/pgsql"
	"github.com/ssengalanto/potato-project/pkg/server"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Run bootstrap and runs the application.
func Run() {
	cfg, err := config.New(constants.Dev, constants.ViperConfigType)
	if err != nil {
		log.Fatal(err)
	}

	slog, err := logger.New(cfg.GetString(constants.AppEnv), cfg.GetString(constants.LogType))
	if err != nil {
		log.Fatal(err)
	}

	db, err := pgsql.NewConnection(cfg.GetString(constants.PgsqlDSN))
	if err != nil {
		slog.Fatal("connection failed", map[string]any{"err": err})
	}
	defer db.Close()

	repo := repository.NewAccountRepository(db)
	mediator := mediatr.NewMediatr()

	router := http.NewRouter()
	router.Mount("/swagger", httpSwagger.WrapHandler)

	command.RegisterHandlers(slog, repo, mediator)
	query.RegisterHandlers(slog, repo, mediator)
	http.RegisterHandlers(slog, router, mediator)

	svr := server.New(cfg.GetInt(constants.AccountServicePort), router)
	err = svr.Start()
	if err != nil {
		slog.Info("shutting down http server", nil)
		slog.Fatal("cannot start http server:", map[string]any{"err": err})
	}
}
