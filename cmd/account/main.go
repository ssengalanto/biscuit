package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ssengalanto/potato-project/pkg/config"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/logger"
)

func main() {
	cfg, err := config.New(constants.Dev, constants.ViperConfigType)
	if err != nil {
		panic("config error")
	}

	log, err := logger.New(cfg.GetString(constants.AppEnv), cfg.GetString(constants.LogType))
	if err != nil {
		panic("log error")
	}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:              ":8081",
		ReadHeaderTimeout: 3 * time.Second, //nolint:gomnd //unnecessary
		Handler:           mux,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Potato Project!! 🚀")
		log.Info("working", nil)
	})

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("server shuts down", nil)
	}
}
