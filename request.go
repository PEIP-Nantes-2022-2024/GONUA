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

type CoursItem struct {
	ID                              int         `json:"id"`
	CelcatID                        string      `json:"celcat_id"`
	Categories                      string      `json:"categories"`
	StartAt                         time.Time   `json:"start_at"`
	EndAt                           time.Time   `json:"end_at"`
	Notes                           string      `json:"notes"`
	Custom1                         interface{} `json:"custom1"`
	Custom2                         interface{} `json:"custom2"`
	Custom3                         interface{} `json:"custom3"`
	Color                           string      `json:"color"`
	PlaceID                         interface{} `json:"place_id"`
	RoomsForBlocks                  string      `json:"rooms_for_blocks"`
	RoomsForItemDetails             string      `json:"rooms_for_item_details"`
	TeachersForBlocks               string      `json:"teachers_for_blocks"`
	TeachersForItemDetails          string      `json:"teachers_for_item_details"`
	EducationalGroupsForBlocks      string      `json:"educational_groups_for_blocks"`
	EducationalGroupsForItemDetails string      `json:"educational_groups_for_item_details"`
	ModulesForBlocks                string      `json:"modules_for_blocks"`
	ModulesForItemDetails           string      `json:"modules_for_item_details"`
}

type Cours []CoursItem

func request(start string, end string, classe string) Cours {
	url := "https://edt-v2.univ-nantes.fr/events?start=" + start + "&end=" + end + "&timetables%5B0%5D=" + classe

	cours := Cours{}

	getJson(url, &cours)

	return cours
}
