package conf

import (
	"encoding/json"
	"os"
)

var ConfigPath string
var V2rayConfigPath string

var Conf *Config

type Config struct {
	Theme      string   `json:"theme"`
	ListerPort int      `json:"lister_port"`
	DnsServers []string `json:"dnsServers"`
}

func (c *Config) SaveConf() error {
	file, err := os.OpenFile(ConfigPath+"/config.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(file).Encode(c); err != nil {
		return err
	}

	return nil
}
