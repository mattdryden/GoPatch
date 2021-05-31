package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type Notification struct {
	message string
	user    string
}

func sendNotification(notification *Notification) bool {

	config := LoadConfig()

	data := url.Values{}

	data.Set("message", notification.message)
	data.Set("user", notification.user)
	data.Set("token", config.Pushover.Token)

	response, err := http.PostForm(config.Pushover.Server+config.Pushover.API, data)

	if err != nil {
		fmt.Printf("Error: %q", err)
	}

	defer response.Body.Close()

	return response.StatusCode == 200

}
