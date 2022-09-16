package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func displayCours(cours CoursItem) {
	if cours.StartAt.Before(time.Now()) && cours.EndAt.After(time.Now()) {
		color.New(color.FgCyan, color.Bold).Println(cours.Categories, ">", cours.StartAt.Local().Hour(), ":", cours.StartAt.Local().Minute(), " {", cours.RoomsForBlocks, "}")
	} else {
		fmt.Println(cours.StartAt.Local().Format("15:04") + " > " + cours.Categories + color.RedString(" {"+cours.RoomsForBlocks+"}"))

func displayHours(current time.Time) {
	if current.Local().Hour() == time.Now().Hour() && current.Day() == time.Now().Day() {
		color.New(color.FgCyan, color.Bold).Println(current.Local().Format("15:04") + " > " + "Free")
	} else {
		fmt.Println(current.Local().Format("15:04") + " > " + "Free")
	}
}
