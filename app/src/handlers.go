package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/tsawler/toolbox"
)

var tools toolbox.Tools

type RequestPayload struct {
	// SendAt       string                 `json:"sendAt"`
	Surroundings []SurroundingsPalyload `json:"surroundings,omitempty"`
}
type SurroundingsPalyload struct {
	Number      int       `json:"number"`
	Timestamp   time.Time `json:"timestamp"`
	Rssi int `json:rssi`
	Tempreture  float64   `json:"tempreture"`
	Moisuture   float64   `josn:"moisuture"`
	AirPressure float64   `json:"airPressure"`
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	err := tools.ReadJSON(w, r, &requestPayload)
	if err != nil {
		fmt.Println(err)
		tools.ErrorJSON(w, err)
		return
	}
	insertPayload(requestPayload.Surroundings)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "OK"}`))
	return
}

func insertPayload(payload []SurroundingsPalyload) {
	var token string
	var bucket string
	var org string
	var dbUrl string
	token = os.Getenv("INFLUXDB_TOKEN")
	bucket = os.Getenv("INFLUXDB_BUCKET")
	org = os.Getenv("INFLUXDB_ORG")
	dbUrl = os.Getenv("DB_URL")
	// fmt.Printf("connectingt to %s , bucket :%s ,org :%s ,token :%s\n", dbUrl, bucket, org, token)
	client := influxdb2.NewClient(dbUrl, token)
	writeAPI := client.WriteAPI(org, bucket)

	// Add this block to listen for errors from the writeAPI
	go func() {
		for err := range writeAPI.Errors() {
			fmt.Println("Error writing to InfluxDB:", err)
		}
	}()
	sort.Slice(payload, func(i, j int) bool {
		return payload[i].Number < payload[j].Number
	})


	for _, v := range payload {
		fmt.Println(v.Rssi)
		p := influxdb2.NewPointWithMeasurement("vuoy_surroundings").
			AddTag("user", "bar").
			AddField("Tempreture", v.Tempreture).
			AddField("Moisuture", v.Moisuture).
			AddField("AirPressure", v.AirPressure).
			AddField("Rssi", v.Rssi).
			SetTime(v.Timestamp)
		writeAPI.WritePoint(p)
		defer client.Close()
	}
	return
}
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := toolbox.JSONResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	fmt.Println("hit the broker")
	_ = tools.WriteJSON(w, http.StatusOK, payload)
}
