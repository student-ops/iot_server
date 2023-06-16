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

type SlackResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

func main() {

	slackBotOAuth := os.Getenv("SLACK_BOT_OAUTH")

	json_data := map[string]interface{}{
		"channel": "#url_manage",
		"text":    "*test*",
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
