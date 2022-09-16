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
}
