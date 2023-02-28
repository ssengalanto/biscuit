package app

import (
	"flag"
	"log"

	repository "github.com/ssengalanto/biscuit/cmd/account/internal/infrastructure/persistence/pgsql"
	"github.com/ssengalanto/biscuit/pkg/config"
	"github.com/ssengalanto/biscuit/pkg/constants"
	"github.com/ssengalanto/biscuit/pkg/logger"
	"github.com/ssengalanto/biscuit/pkg/pgsql"
	"github.com/ssengalanto/biscuit/pkg/redis"
	"github.com/ssengalanto/biscuit/pkg/server"
)

//nolint:gochecknoglobals //intentional for flag vars
var (
	flEnv = flag.String("env", constants.Dev, "The environment in which the application operates.")
	flCfg = flag.String("cfg", constants.DotEnvConfigType, "The config module that is set as the application's default.")
)

// Run bootstrap and runs the application.
func Run() {
	flag.Parse()

	d := getDefaults()
	c, err := config.New(d.env, d.cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog, err := logger.New(c.GetString(constants.AppEnv), c.GetString(constants.LogType))
	if err != nil {
		log.Fatal(err)
	}

	db, err := pgsql.NewConnection(
		c.GetString(constants.PgsqlUser),
		c.GetString(constants.PgsqlPassword),
		c.GetString(constants.PgsqlHost),
		c.GetString(constants.PgsqlPort),
		c.GetString(constants.PgsqlDBName),
		c.GetString(constants.PgsqlQueryParams),
	)
	if err != nil {
		slog.Fatal(err.Error(), map[string]any{"err": err})
	}
	defer db.Close()

	rdb, err := redis.NewUniversalClient(
		c.GetString(constants.RedisHost),
		c.GetString(constants.RedisPort),
		c.GetString(constants.RedisPassword),
		c.GetInt(constants.RedisDB),
	)
	if err != nil {
		slog.Fatal(err.Error(), map[string]any{"err": err})
	}
	defer rdb.Close()

	repo := repository.NewAccountRepository(slog, db)
	mediator := registerMediatorHandlers(slog, repo, rdb)
	mux := registerHTTPHandlers(slog, mediator)

	svr := server.New(c.GetInt(constants.AccountServicePort), mux)
	err = svr.Start()
	if err != nil {
		slog.Info("shutting down http server", nil)
		slog.Fatal("cannot start http server:", map[string]any{"err": err})
	}
}
