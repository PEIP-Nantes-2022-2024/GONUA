package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var RequestClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := RequestClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func request(start string, end string, classe string) Cours {
	url := "https://edt-v2.univ-nantes.fr/events?start=" + start + "&end=" + end + "&timetables%5B0%5D=" + classe

	cours := Cours{}

	getJson(url, &cours)

	return cours
}
