package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
	"time"
)

type Reminder struct {
	name string
	time time.Time
}

var reminders []Reminder

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	type reminder struct {
		name string
		time time.Time
	}
	showListOfReminders()
	for _, reminder := range reminders {
		if reminder.time == time.Now() {
			showReminder(reminder.name)
		}
	}
	app.Exec()
}
