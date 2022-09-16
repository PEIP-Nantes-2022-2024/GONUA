package main

import "time"

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
