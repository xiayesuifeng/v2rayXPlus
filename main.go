package main

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"gitlab.com/xiayesuifeng/v2rayxplus/core"
	"gitlab.com/xiayesuifeng/v2rayxplus/ui"
	"log"
	"os"
	"path/filepath"
)

var (
	help   = flag.Bool("h", false, "help")
	config = flag.String("c", "", "config file")
)

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *config != "" {
		if err := core.StartV2ray(*config); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

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
	conf.V2rayConfigPath = conf.ConfigPath + "/v2ray"

	if _, err := os.Stat(conf.V2rayConfigPath); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(conf.V2rayConfigPath, 0755); err != nil {
				log.Panicln("config dir create failure")
			}
		}
	}

	if err := parseConfig(); err != nil {
		log.Panicln(err)
	}
}

func getConfPath() (string, error) {
	path := os.Getenv("HOME")
	if path == "" {
		return "", errors.New("get home failure")
	}

	return filepath.Join(path, ".config/V2rayXPlus"), nil
}

func parseConfig() error {
	file, err := os.OpenFile(conf.ConfigPath+"/config.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	conf.Conf = &conf.Config{}
	err = json.NewDecoder(file).Decode(conf.Conf)
	if err != nil {
		if err.Error() == "EOF" {
			conf.Conf.Theme = "light"
			conf.Conf.ListerPort = 8102
			conf.Conf.DnsServers = []string{"8.8.8.8", "8.8.4.4", "localhost"}
			if err := json.NewEncoder(file).Encode(conf.Conf); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
