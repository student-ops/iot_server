package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Tunnel struct {
	PublicURL string `json:"public_url"`
}

type TunnelResponse struct {
	Tunnels []Tunnel `json:"tunnels"`
}

func fetch_ngrok_url(host_port string) string {
	resp, err := http.Get("http://" + host_port + "/api/tunnels")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data TunnelResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	var url string
	if len(data.Tunnels) > 0 {
		fmt.Println(data.Tunnels[0].PublicURL)
		url = data.Tunnels[0].PublicURL
	} else {
		fmt.Println("No tunnels found")
	}
	return url
}

type SlackResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

func main() {

	slackBotOAuth := os.Getenv("SLACK_BOT_OAUTH")
	// slackBotOAuth := ""
	ngrok_app := fetch_ngrok_url("ngrok_app:4040")
	ngrok_influx := fetch_ngrok_url("ngrok_influxdb:4040")

	json_data := map[string]interface{}{
		"channel": "#url_manage",
		"text":    "*App:* \n" + ngrok_app + "\n*Grafana:* \n" + ngrok_influx,
	}

	payload, err := json.Marshal(json_data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+slackBotOAuth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var slackResponse SlackResponse
	err = json.Unmarshal(body, &slackResponse)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
		return
	}

	if !slackResponse.OK {
		log.Fatal("Error posting message:", slackResponse.Error)
		return
	}

	fmt.Println("Response body:", string(body))
}
