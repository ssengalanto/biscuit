package app

import (
	"log"

	"github.com/ssengalanto/biscuit/pkg/config"
	"github.com/ssengalanto/biscuit/pkg/constants"
	"github.com/ssengalanto/biscuit/pkg/logger"
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

	mediator := registerMediatorHandlers(slog)
	mux := registerHTTPHandlers(slog, mediator)

	svr := server.New(cfg.GetInt(constants.AuthServicePort), mux)
	err = svr.Start()
	if err != nil {
		slog.Info("shutting down http server", nil)
		slog.Fatal("cannot start http server:", map[string]any{"err": err})
	}
}
