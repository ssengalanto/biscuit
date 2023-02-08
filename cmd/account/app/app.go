package app

import (
	"log"

	repository "github.com/ssengalanto/biscuit/cmd/account/internal/infrastructure/persistence/pgsql"
	"github.com/ssengalanto/biscuit/pkg/config"
	"github.com/ssengalanto/biscuit/pkg/constants"
	"github.com/ssengalanto/biscuit/pkg/logger"
	"github.com/ssengalanto/biscuit/pkg/pgsql"
	"github.com/ssengalanto/biscuit/pkg/redis"
	"github.com/ssengalanto/biscuit/pkg/server"
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

	db, err := pgsql.NewConnection(
		cfg.GetString(constants.PgsqlUser),
		cfg.GetString(constants.PgsqlPassword),
		cfg.GetString(constants.PgsqlHost),
		cfg.GetString(constants.PgsqlPort),
		cfg.GetString(constants.PgsqlDBName),
		cfg.GetString(constants.PgsqlSslMode),
	)
	if err != nil {
		slog.Fatal(err.Error(), map[string]any{"err": err})
	}
	defer db.Close()

	rdb, err := redis.NewUniversalClient(
		cfg.GetString(constants.RedisHost),
		cfg.GetString(constants.RedisPort),
		cfg.GetString(constants.RedisPassword),
		cfg.GetInt(constants.RedisDB),
	)
	if err != nil {
		slog.Fatal(err.Error(), map[string]any{"err": err})
	}
	defer rdb.Close()

	repo := repository.NewAccountRepository(slog, db)
	mediator := registerMediatorHandlers(slog, repo, rdb)
	mux := registerHTTPHandlers(slog, mediator)

	svr := server.New(cfg.GetInt(constants.AccountServicePort), mux)
	err = svr.Start()
	if err != nil {
		slog.Info("shutting down http server", nil)
		slog.Fatal("cannot start http server:", map[string]any{"err": err})
	}
}
