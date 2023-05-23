package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_insertdb(t *testing.T) {

	payload := []SurroundingsPalyload{
		{
			Number:      1,
			Timestamp:   time.Now(),
			Rssi:        -3,
			Tempreture:  10.5,
			Moisuture:   0.5,
			AirPressure: 1012.2,
		},
		{
			Number:      2,
			Timestamp:   time.Now(),
			Rssi:        -4,
			Tempreture:  15.3,
			Moisuture:   0.6,
			AirPressure: 1011.9,
		},
	}
	test_payload := payload
	timestampStr := "2023-05-23T10:00:00Z"
	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		t.Errorf("Failed to parse timestamp: %s", err)
		return
	}
	test_payload[0].Timestamp = timestamp
	test_payload[1].Timestamp = timestamp
	fmt.Println(time.Now())

	InsertPayload(payload)
	InsertPayload(test_payload)
}
