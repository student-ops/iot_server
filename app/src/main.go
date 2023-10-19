package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const httpPort = "8080"

var token = os.Getenv("INFLUXDB_TOKEN")
var bucket = os.Getenv("INFLUXDB_BUCKET")
var org = os.Getenv("INFLUXDB_ORG")
var dbUrl = os.Getenv("DB_URL")
var (
	clientOptions = influxdb2.DefaultOptions()
	client        = influxdb2.NewClientWithOptions(dbUrl, token, clientOptions)
)

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
