package conf

import (
	"encoding/json"
	"os"
)

type OutboundConfig struct {
	Protocol      string          `json:"protocol"`
	Tag           string          `json:"tag,omitempty"`
	Settings      json.RawMessage `json:"settings,omitempty"`
	StreamSetting *StreamConfig   `json:"streamSettings,omitempty"`
	MuxSettings   *MuxConfig      `json:"mux,omitempty"`
}

type InboundConfig struct {
	Port           *uint           `json:"port"`
	Protocol       string          `json:"protocol"`
	StreamSetting  *StreamConfig   `json:"streamSettings,omitempty"`
	Settings       json.RawMessage `json:"settings,omitempty"`
	Tag            string          `json:"tag,omitempty"`
	DomainOverride *[]string       `json:"domainOverride,omitempty"`
}

type SocketConfig struct {
	Mark   int32  `json:"mark"`
	TFO    *bool  `json:"tcpFastOpen,omitempty"`
	TProxy string `json:"tproxy,omitempty"`
}

type StreamConfig struct {
	Security       string        `json:"security,omitempty"`
	SocketSettings *SocketConfig `json:"sockopt,omitempty"`
}

type MuxConfig struct {
	Enabled     bool   `json:"enabled"`
	Concurrency uint16 `json:"concurrency"`
}

type V2rayConfig struct {
	Port            uint16           `json:"port"`
	RouterConfig    *RouterConfig    `json:"routing"`
	DNSConfig       *DnsConfig       `json:"dns"`
	InboundConfig   *InboundConfig   `json:"inbound"`
	OutboundConfig  *OutboundConfig  `json:"outbound"`
	InboundDetours  []InboundConfig  `json:"inboundDetour,omitempty"`
	OutboundDetours []OutboundConfig `json:"outboundDetour,omitempty"`
}

func ParseV2ray(conf string) (*V2rayConfig, error) {
	confJson, err := os.Open(conf)
	if err != nil {
		return nil, err
	}

	v2rayConfig := new(V2rayConfig)
	if err = json.NewDecoder(confJson).Decode(v2rayConfig); err != nil {
		return nil, err
	}

	return v2rayConfig, nil
}

func (v2ray *V2rayConfig) Save(conf string) error {
	confJson, err := os.OpenFile(conf, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	return json.NewEncoder(confJson).Encode(v2ray)
}
