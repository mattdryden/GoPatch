package main

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

type Notification struct {
	message string
	user    string
	token   string
}

func (n Notification) Encode() (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(n)
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		values.Set(typ.Field(i).Name, fmt.Sprint(iVal.Field(i)))
	}
	return
}

func (n Notification) Send(config *Config) (bool, error) {
	response, err := http.PostForm(config.Pushover.Server+config.Pushover.API, n.Encode())

	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	return response.StatusCode == 200, nil
}
