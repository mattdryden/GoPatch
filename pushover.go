package main

import (
	"net/http"
	"net/url"
)

type Notification struct {
	message string
	user    string
	token   string
}

func sendNotification(notification *Notification, config *Config) (bool, error) {

	data := url.Values{}

	data.Set("message", notification.message)
	data.Set("user", notification.user)
	data.Set("token", notification.token)

	response, err := http.PostForm(config.Pushover.Server+config.Pushover.API, data)

	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	return response.StatusCode == 200, nil

}
