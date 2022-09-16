package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func display(cours CoursItem) {
	if cours.StartAt.Before(time.Now()) && cours.EndAt.After(time.Now()) {
		color.New(color.FgCyan, color.Bold).Println(cours.Categories, ">", cours.StartAt.Local().Hour(), ":", cours.StartAt.Local().Minute(), " {", cours.RoomsForBlocks, "}")
	} else {
		fmt.Println(cours.StartAt.Local().Format("15:04") + " > " + cours.Categories + color.RedString(" {"+cours.RoomsForBlocks+"}"))
	}
}
