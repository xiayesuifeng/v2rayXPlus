package main

import (
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/ui"
	"os"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetApplicationVersion("0.1.0")
	app.SetApplicationName("v2rayXPlus")

	ui.NewMainWindow().Show()

	os.Exit(app.Exec())
}
