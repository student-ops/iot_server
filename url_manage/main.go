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

func main() {
	resp, err := http.Get("http://localhost:4041/api/tunnels")
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

	slackBotOAuth := os.Getenv("SLACK_BOT_OAUTH")

	json_data := map[string]interface{}{
		"channel": "#url_manage",
		"text":    "*Grafana:* \n"+url,
	}

	payload, err := json.Marshal(json_data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	req, err := http.NewRequest("POST",  "https://slack.com/api/chat.postMessage", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+slackBotOAuth)

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response:", resp.Status)
}
