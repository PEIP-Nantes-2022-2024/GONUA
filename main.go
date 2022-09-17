package main

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func completer(d prompt.Document) []prompt.Suggest {
	classes := loadJson()

	k := []prompt.Suggest{}

	for t, d := range classes {
		k = append(k, prompt.Suggest{Text: t, Description: d})
	}
	return prompt.FilterHasPrefix(k, d.GetWordBeforeCursor(), true)
}

func main() {
	fmt.Println("Hello, welcome to the calendar CLI")
	promptUI := promptui.Select{
		Label: "Select Action",
		Items: []string{"See today classes", "See next day classes", "Full week", "Exit"},
	}

	_, result, err := promptUI.Run()

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
		for start_date.Weekday() > time.Friday || start_date.Weekday() < time.Monday {
			start_date = start_date.AddDate(0, 0, 1)
		}
		end_date = start_date
		text = "Classes for the " + color.CyanString(start_date.Format("Mon 02 Jan 2006"))
	case "See today classes":
		start_date = time.Now()
		end_date = start_date
		text = "Classes for the " + color.CyanString(start_date.Format("Mon 02 Jan 2006"))
	case "Full week":
		start_date = time.Now().AddDate(0, 0, int(-time.Now().Weekday()+1))
		end_date = start_date.AddDate(0, 0, 4)
		text = "Classes for the week from " + color.CyanString(start_date.Format("Mon 02 Jan 2006")) + " to " + color.CyanString(end_date.Format("Mon 02 Jan 2006"))
	case "Exit":
		color.New(color.FgRed).Println("Bye bye")
		return
	}

	promptUI = promptui.Select{
		Label: "Select a class",
		Items: []string{
			"Default class (mine so 171PEIP)",
			"I want another class",
			"Exit (yes you can cancel now too, I'm so generous with cancel state)",
		},
	}

	_, resultc, err := promptUI.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	var classe string
	switch resultc {
	case "Default class (mine so 171PEIP)":
		classe = "171PEIP"
	case "I want another class":
		fmt.Println("Enter group name (ex: 171PEIP)")
		classe = prompt.Input("> ", completer)
	case "Exit (yes you can cancel now too, I'm so generous with cancel state)":
		color.New(color.FgRed).Println("Bye bye")
		return
	}

	classes := loadJson()

	fmt.Println(text)
	fmt.Println("")
	s := spinner.New(spinner.CharSets[39], 250*time.Millisecond)
	s.Start()
	cours := request(start_date.Format("2006-01-02"), end_date.Format("2006-01-02"), classes[classe])
	s.Stop()

	if result != "Full week" {
		current := time.Date(start_date.Year(), start_date.Month(), start_date.Day(), 8, 0, 0, 0, time.Local)
		displayDay(current, cours)
	} else {
		for start_date.Before(end_date) {
			day_classes := Cours{}
			for _, c := range cours {
				if c.StartAt.Weekday() == start_date.Weekday() {
					day_classes = append(day_classes, c)
				}
			}
			if start_date.YearDay() == time.Now().YearDay() && start_date.Year() == time.Now().Year() {
				fmt.Println("====== " + color.HiCyanString(start_date.Format("Mon 02 Jan 2006")) + " ======")
			} else {
				fmt.Println("====== " + color.BlueString(start_date.Format("Mon 02 Jan 2006")) + " ======")
			}
			displayDay(start_date, day_classes)
			fmt.Println("")
			start_date = time.Date(start_date.Year(), start_date.Month(), start_date.Day()+1, 8, 0, 0, 0, time.Local)
		}
	}
}
