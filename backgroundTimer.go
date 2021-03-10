package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
	"time"
)
func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	type reminder struct {
		time time.Time
	}
	reminders := []reminder{}
	for {
		for _, reminder := range reminders {
			if reminder.time == time.Now() {
				//
			}
		}
	}
	app.Exec()
}
