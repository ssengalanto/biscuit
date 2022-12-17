package app

import (
	"log"

	"github.com/ssengalanto/potato-project/cmd/account/internal/application/command"
	repository "github.com/ssengalanto/potato-project/cmd/account/internal/infrastructure/persistence/pgsql"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/http"
	"github.com/ssengalanto/potato-project/pkg/config"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/logger"
	"github.com/ssengalanto/potato-project/pkg/pgsql"
	"github.com/ssengalanto/potato-project/pkg/server"
)

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
	router := http.NewRouter()

	command.RegisterHandlers(slog, repo)
	http.RegisterHandlers(slog, router)

	svr := server.New(cfg.GetInt(constants.AccountServicePort), router)
	err = svr.Start()
	if err != nil {
		slog.Info("shutting down http server", nil)
		slog.Fatal("cannot start http server:", map[string]any{"err": err})
	}
}
