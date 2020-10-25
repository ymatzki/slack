package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// get webhook url
	url, exists := os.LookupEnv("SLACK_WEBHOOK_URL")
	if !exists {
		panic("SLACK_WEBHOOK_URL environment is not set")
	}

	// post message
	message := message{Text: "Hello, World!"}
	err := post(url, message)
	if err != nil {
		panic(err)
	}
}

type message struct {
	Text string `json:"text"`
}

func post(url string, message message) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fmt.Printf("%#v", res)
	return nil
}
