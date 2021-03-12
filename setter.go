package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"strconv"
	"strings"
	"time"
)

func showSetter() {
	window := widgets.NewQMainWindow(nil, 0)

	centralWidget := widgets.NewQWidget(nil, 0)
	window.SetCentralWidget(centralWidget)

	layout := widgets.NewQVBoxLayout()
	centralWidget.SetLayout(layout)

	nameEdit := widgets.NewQLineEdit2("<name>", nil)
	layout.AddWidget(nameEdit, 0, 4)

	timeEdit := widgets.NewQLineEdit2("0 min, 0 sec", nil)
	layout.AddWidget(timeEdit, 0, core.Qt__AlignHCenter)

	button := widgets.NewQPushButton2("close", nil)
	button.SetFont(gui.NewQFont2("sans", 20, 0, false))
	button.SetFixedSize2(100, 50)
	button.ConnectClicked(func(checked bool) {
		strList := strings.Split(timeEdit.Text(), " ")
		min, _ := strconv.Atoi(strList[0])
		sec, _ := strconv.Atoi(strList[2])
		duration := time.Minute*time.Duration(min) + time.Second*time.Duration(sec)
		reminderTime := time.Now().Add(duration)
		reminders = append(reminders, Reminder{nameEdit.Text(), reminderTime})
		window.Close()
	})
	layout.AddWidget(button, 0, core.Qt__AlignHCenter)

	window.ShowMaximized()
}
