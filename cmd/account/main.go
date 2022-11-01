package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Potato Project! 🚀")
	})

	log.Fatal(http.ListenAndServe(":8081", mux))
}
