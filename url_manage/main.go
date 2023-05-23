package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	if len(data.Tunnels) > 0 {
		fmt.Println(data.Tunnels[0].PublicURL)
	} else {
		fmt.Println("No tunnels found")
	}
}
