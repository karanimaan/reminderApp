package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	file := core.NewQFile2("listOfReminders.go")
	file.Open(core.QIODevice__ReadOnly)
	form := uitools.NewQUiLoader(nil).Load(file, nil)
	file.Close()


	addButton := widgets.NewQPushButtonFromPointer(form.FindChild("addButton", core.Qt__FindChildrenRecursively).Pointer())

	addButton.ConnectClicked(func(checked bool) {
		form.Close()
	})

	widgets.QApplication_Exec()
}
