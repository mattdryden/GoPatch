package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Quote struct {
	Quote     string `json:"q"`
	Author    string `json:"a"`
	Formatted string `json:"f"`
}

func getQuotes(config *Config) (string, error) {
	response, err := http.Get(config.Quotes.Server + config.Quotes.API)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	var quotes []Quote
	json.Unmarshal([]byte(body), &quotes)

	if response.StatusCode == 200 {
		return fmt.Sprintf("%s - %s", quotes[0].Quote, quotes[0].Author), nil
	} else {
		return "", err
	}

}
