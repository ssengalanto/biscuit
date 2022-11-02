package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:              ":8081",
		ReadHeaderTimeout: 3 * time.Second, //nolint:gomnd //unnecessary
		Handler:           mux,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Potato Project!! 🚀")
	})

	log.Fatal(server.ListenAndServe())
}
