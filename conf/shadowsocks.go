package conf

import (
	"encoding/json"
)

type ShadowsocksServer struct {
	Address  string `json:"address"`
	Port     uint16 `json:"port"`
	Cipher   string `json:"method"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
	Ota      bool   `json:"ota"`
	Level    byte   `json:"level"`
}

type ShadowsocksConfig struct {
	Servers []ShadowsocksServer `json:"servers"`
}

func NewShadowsocksServer(jsonData json.RawMessage) (*ShadowsocksConfig, error) {
	conf := &ShadowsocksConfig{}
	return conf, json.Unmarshal(jsonData, conf)
}
