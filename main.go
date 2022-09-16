package main

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

const classe_171 = "72431"

func main() {
	fmt.Println("Hello, welcome to the calendar CLI")
	prompt := promptui.Select{
		Label: "Select Action",
		Items: []string{"See today classes", "See next day classes", "Full week", "Exit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	var start_date time.Time
	var end_date time.Time
	var text string

	switch result {
	case "See next day classes":
		start_date = time.Now().AddDate(0, 0, 1)
		end_date = time.Now().AddDate(0, 0, 1)
		for start_date.Weekday() > time.Friday || start_date.Weekday() < time.Monday {
			start_date = start_date.AddDate(0, 0, 1)
			end_date = end_date.AddDate(0, 0, 1)
		}
		text = "Classes for the " + color.CyanString(start_date.Format("02 01 2006"))
	case "See today classes":
		start_date = time.Now()
		end_date = time.Now()
		text = "Classes for the " + color.CyanString(start_date.Format("02 01 2006"))
	case "Full week":
		start_date = time.Now()
		end_date = time.Now().AddDate(0, 0, 6)
		text = "Classes for the week of " + color.CyanString(start_date.Format("02 01 2006"))
	case "Exit":
		color.New(color.FgRed).Println("Bye bye")
		return
	}

	fmt.Println(text)
	fmt.Println("")
	s := spinner.New(spinner.CharSets[39], 250*time.Millisecond)
	s.Start()
	cours := request(start_date.Format("2006-01-02"), end_date.Format("2006-01-02"), classe_171)

	s.Stop()

	if len(cours) == 0 {
		fmt.Println("No classes found")
		return
	}
	if result != "Full week" {
		if cours[0].StartAt.After(time.Now()) {
			color.New(color.FgCyan, color.Bold).Println("Now")
		}
		reader := 0
		current := time.Date(start_date.Year(), start_date.Month(), start_date.Day(), 8, 0, 0, 0, time.Local)
		for reader < len(cours) {
			if current.Local().Hour() == cours[reader].StartAt.Local().Hour() {
				display(cours[reader])
				current = current.Add(time.Duration(cours[reader].EndAt.Sub(cours[reader].StartAt)) + 10*time.Minute)
				reader++
			} else {
				if current.Local().Minute() == 0 {
					displayHours(current)
				}
				current = current.Add(time.Hour / 2)
			}
		}
		if cours[len(cours)-1].EndAt.Before(time.Now()) {
			color.New(color.FgCyan, color.Bold).Println("Now")
		}
	}
}
