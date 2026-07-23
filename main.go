package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileHandler := http.FileServer(
		http.Dir("."),
	)

	mux.Handle("/", fileHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
