package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ssengalanto/potato-project/pkg/config"
	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/logger"
	"github.com/ssengalanto/potato-project/pkg/pgsql"
)

func main() {
	cfg, err := config.New(constants.Dev, constants.ViperConfigType)
	if err != nil {
		log.Fatal(err)
	}

	struclog, err := logger.New(cfg.GetString(constants.AppEnv), cfg.GetString(constants.LogType))
	if err != nil {
		log.Fatal(err)
	}

	db, err := pgsql.NewConnection(cfg.GetString(constants.PgsqlDSN))
	if err != nil {
		struclog.Fatal("connection failed", map[string]any{"err": err})
	}
	defer db.Close()

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.GetString(constants.AccountServicePort)),
		ReadHeaderTimeout: constants.ReadTimeout,
		Handler:           mux,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Potato Project!! 🚀")
		struclog.Info("working", nil)
	})

	err = server.ListenAndServe()
	if err != nil {
		struclog.Fatal("server shuts down:", map[string]any{"err": err})
	}
}
