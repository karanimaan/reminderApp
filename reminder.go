package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func main() {

	window := widgets.NewQMainWindow(nil, 0)

	centralWidget := widgets.NewQWidget(nil, 0)
	window.SetCentralWidget(centralWidget)

	layout := widgets.NewQVBoxLayout()
	centralWidget.SetLayout(layout)

	label := widgets.NewQLabel2("test", nil, 0)
	label.SetFont(gui.NewQFont2("sans", 50, 0, false))
	layout.AddWidget(label, 0, core.Qt__AlignHCenter)

	button := widgets.NewQPushButton2("close", nil)
	button.SetFont(gui.NewQFont2("sans", 20, 0, false))
	button.SetFixedSize2(100, 50)
	button.ConnectClicked(func(checked bool) {
		window.Close()
	})
	layout.AddWidget(button, 0, core.Qt__AlignHCenter)

	window.ShowMaximized()

}
