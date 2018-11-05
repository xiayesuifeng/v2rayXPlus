package conf

import "encoding/json"

type VMessAccount struct {
	ID       string `json:"id"`
	AlterIds uint16 `json:"alterId"`
	Security string `json:"security,omitempty"`
}

type VMessOutboundTarget struct {
	Address string         `json:"address"`
	Port    uint16         `json:"port"`
	Users   []VMessAccount `json:"users"`
}

type VMessOutboundConfig struct {
	Receivers []*VMessOutboundTarget `json:"vnext"`
}

func NewVMessOutboundConfig(jsonData json.RawMessage) (*VMessOutboundConfig, error) {
	conf := &VMessOutboundConfig{}
	return conf, json.Unmarshal(jsonData, conf)
}
