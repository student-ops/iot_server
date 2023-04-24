package sample

import (
	"fmt"
	"math/rand"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type ThermostatSetting struct {
	user string
	max  float64 //temperature
	avg  float64 //temperature
}

func main() {
	// You can generate a Token from the "Tokens Tab" in the UI
	const token = "F-QFQpmCL9UkR3qyoXnLkzWj03s6m4eCvYgDl1ePfHBf9ph7yxaSgQ6WN0i9giNgRTfONwVMK1f977r_g71oNQ=="
	const bucket = "test"
	const org = "iot"

	client := influxdb2.NewClient("http://localhost:8086", token)
	// always close client at the end
	// get non-blocking write client
	writeAPI := client.WriteAPI(org, bucket)

	// write line protocol
	for {
		rand.Seed(time.Now().UnixNano())

		avg := float64(20 + rand.Float64()*20)
		max := float64(30 + rand.Float64()*20)

		t := ThermostatSetting{user: "foo", avg: float64(avg), max: float64(max)}
		fmt.Println(fmt.Printf("user: %s, avg: %f, max: %f", t.user, t.avg, t.max))
		p := influxdb2.NewPointWithMeasurement("thermostat").
			AddTag("unit", "temperature").
			AddTag("user", t.user).
			AddField("avg", t.avg).
			AddField("max", t.max).
			SetTime(time.Now())
		writeAPI.WritePoint(p)
		// Flush writes
		writeAPI.Flush()

		writeAPI.Flush()
		defer client.Close()
		time.Sleep(3 * time.Second)
	}
	// Flush writes
}
