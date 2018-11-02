package main

import (
	"errors"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"gitlab.com/xiayesuifeng/v2rayxplus/ui"
	"log"
	"os"
	"path/filepath"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetApplicationVersion("0.1.0")
	app.SetApplicationName("v2rayXPlus")
	app.SetWindowIcon(gui.NewQIcon5(":/resources/v2rayXPlus-64px.svg"))

	ui.NewMainWindow().Show()

	os.Exit(app.Exec())
}

func init() {
	var err error
	conf.ConfigPath, err = getConfPath()
	if err != nil {
		log.Panicln(err)
	}

	if _, err := os.Stat(conf.ConfigPath); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(conf.ConfigPath, 0755); err != nil {
				log.Panicln("config dir create failure")
			}
		}
	}
}

func getConfPath() (string, error) {
	path := os.Getenv("HOME")
	if path == "" {
		return "", errors.New("get home failure")
	}

	return filepath.Join(path, ".config/V2rayXPlus"), nil
}
