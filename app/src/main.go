package main

import (
	"fmt"
	"log"
	"net/http"
)

const httpPort = "8080"

type Config struct{}

func main() {
	var app Config

	log.Printf("Starting broker service on port %s\n", httpPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
