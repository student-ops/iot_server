package main

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"time"

	// influxdb2 "github.com/influxdata/influxdb-client-go/v2"
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
	Rssi        int       `json:"rssi"`
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

	err = InsertPayload(requestPayload.Surroundings)
	if err != nil {
		fmt.Println("Failed to insert payload:", err)
		tools.ErrorJSON(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "OK"}`))
	return
}
func InsertPayload(payload []SurroundingsPalyload) error {
	clientOptions := influxdb2.DefaultOptions()
	client := influxdb2.NewClientWithOptions(dbUrl, token, clientOptions)
	defer client.Close()

	writeAPI := client.WriteAPIBlocking(org, bucket)

	sort.Slice(payload, func(i, j int) bool {
		return payload[i].Number < payload[j].Number
	})

	for _, v := range payload {
		p := influxdb2.NewPointWithMeasurement("vuoy_surroundings").
			AddTag("user", "1").
			AddField("Tempreture", v.Tempreture).
			AddField("Moisuture", v.Moisuture).
			AddField("AirPressure", v.AirPressure).
			AddField("Rssi", v.Rssi).
			SetTime(time.Now())
		// 同期的に書き込み
		if err := writeAPI.WritePoint(context.Background(), p); err != nil {
			return err // 書き込み中にエラーが発生した場合、エラーを直接返します
		}
	}

	return nil
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := toolbox.JSONResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	fmt.Println("hit the broker")
	_ = tools.WriteJSON(w, http.StatusOK, payload)
}
