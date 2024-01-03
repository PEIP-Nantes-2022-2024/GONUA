package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func preciseIfPolytech(cours CoursItem) CoursItem {
	cours.RoomsForBlocks = color.HiMagentaString(" { " + cours.RoomsForBlocks)
	if cours.StartAt.Weekday() == time.Thursday && cours.StartAt.Hour() > 12 {
		cours.RoomsForBlocks = cours.RoomsForBlocks + " - " + color.HiCyanString("Polytech")
	}
	cours.RoomsForBlocks = cours.RoomsForBlocks + color.HiMagentaString(" }")
	return cours
}

func displayCours(cours CoursItem) {
	cours = preciseIfPolytech(cours)
	if cours.StartAt.Before(time.Now()) && cours.EndAt.After(time.Now()) {
		color.New(color.FgCyan, color.Bold).Println(cours.StartAt.Local().Format("15:04") + " - " + cours.EndAt.Local().Format("15:04") + " > " + cours.Categories + cours.RoomsForBlocks)
	} else {
		fmt.Println(cours.StartAt.Local().Format("15:04") + " > " + cours.Categories + cours.RoomsForBlocks)
	}
}

func displayHours(current time.Time) {
	if current.Local().Hour() == time.Now().Hour() && current.Day() == time.Now().Day() {
		color.New(color.FgCyan, color.Bold).Println(current.Local().Format("15:04") + " > " + "Free")
	} else {
		fmt.Println(current.Local().Format("15:04") + " > " + "Free")
	}
}

func displayDay(current time.Time, cours Cours) {
	if len(cours) == 0 {
		fmt.Println("No classes found")
		return
	}
	reader := 0
	for reader < len(cours) {
		if current.Local().Hour() == cours[reader].StartAt.Local().Hour() {
			if reader+1 < len(cours) {
				if cours[reader+1].StartAt.Local().Hour() == current.Local().Hour() {
					print(color.RedString("- "))
					displayCours(cours[reader])
					reader++
					print(color.RedString("- "))
				}
			}
			displayCours(cours[reader])
			current = current.Add(time.Duration(cours[reader].EndAt.Sub(cours[reader].StartAt)) + 10*time.Minute)
			reader++
		} else {
			if current.Local().Minute() == 0 {
				displayHours(current)
			}
			current = current.Add(time.Hour / 2)
		}
	}
}
