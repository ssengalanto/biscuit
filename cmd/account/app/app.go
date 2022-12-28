package app

import (
	"fmt"
	"log"

	repository "github.com/ssengalanto/potato-project/cmd/account/internal/infrastructure/persistence/pgsql"
	"github.com/ssengalanto/potato-project/pkg/config"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/logger"
	"github.com/ssengalanto/potato-project/pkg/pgsql"
	"github.com/ssengalanto/potato-project/pkg/redis"
	"github.com/ssengalanto/potato-project/pkg/server"
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

	rdb := redis.NewUniversalClient(
		fmt.Sprintf("%s:%d", cfg.GetString(constants.RedisURL), cfg.GetInt(constants.RedisPort)),
		cfg.GetInt(constants.RedisDB),
		cfg.GetString(constants.RedisPassword),
	)
	defer rdb.Close()

	repo := repository.NewAccountRepository(db)
	mediatr := RegisterMediatrHandlers(slog, repo, rdb)
	mux := RegisterHTTPHandlers(slog, mediatr)

	svr := server.New(cfg.GetInt(constants.AccountServicePort), mux)
	err = svr.Start()
	if err != nil {
		slog.Info("shutting down http server", nil)
		slog.Fatal("cannot start http server:", map[string]any{"err": err})
	}
}
