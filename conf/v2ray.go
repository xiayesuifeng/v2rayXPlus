package conf

import (
	"encoding/json"
	"os"
)

type OutboundConfig struct {
	Protocol      string          `json:"protocol"`
	Tag           string          `json:"tag"`
	Settings      json.RawMessage `json:"settings"`
	StreamSetting *StreamConfig   `json:"streamSettings"`
	MuxSettings   *MuxConfig      `json:"mux"`
}

type InboundConfig struct {
	Port           *uint           `json:"port"`
	Protocol       string          `json:"protocol"`
	StreamSetting  *StreamConfig   `json:"streamSettings"`
	Settings       json.RawMessage `json:"settings"`
	Tag            string          `json:"tag"`
	DomainOverride *[]string       `json:"domainOverride"`
}

type SocketConfig struct {
	Mark   int32  `json:"mark"`
	TFO    *bool  `json:"tcpFastOpen"`
	TProxy string `json:"tproxy"`
}

type StreamConfig struct {
	Security       string        `json:"security"`
	SocketSettings *SocketConfig `json:"sockopt"`
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
	InboundDetours  []InboundConfig  `json:"inboundDetour"`
	OutboundDetours []OutboundConfig `json:"outboundDetour"`
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
