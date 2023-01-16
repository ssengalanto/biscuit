package app

import (
	"fmt"
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

	db, err := pgsql.NewConnection(cfg.GetString(constants.PgsqlDSN))
	if err != nil {
		slog.Fatal(err.Error(), map[string]any{"err": err})
	}
	defer db.Close()

	rdb, err := redis.NewUniversalClient(
		fmt.Sprintf("%s:%d", cfg.GetString(constants.RedisURL), cfg.GetInt(constants.RedisPort)),
		cfg.GetInt(constants.RedisDB),
		cfg.GetString(constants.RedisPassword),
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
