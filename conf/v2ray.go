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
	Port               uint16            `json:"port,omitempty"`
	RouterConfig       *RouterConfig     `json:"routing"`
	DNSConfig          *DnsConfig        `json:"dns"`
	InboundConfigList  []*InboundConfig  `json:"inbounds"`
	OutboundConfigList []*OutboundConfig `json:"outbounds,omitempty"`
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

func NewV2rayConfig() *V2rayConfig {
	v2ray := &V2rayConfig{}

	port := uint(Conf.ListerPort)
	dokodemoConfig := &InboundConfig{Port: &port}
	dokodemoConfig.DomainOverride = &[]string{"tls", "http"}
	dokodemoConfig.Protocol = "dokodemo-door"
	dokodemoConfig.Settings = []byte(`{"network": "tcp,udp","followRedirect": true}`)

	freedomConfig := &OutboundConfig{}
	freedomConfig.Protocol = "freedom"
	freedomConfig.Tag = "direct"
	freedomConfig.StreamSetting = &StreamConfig{SocketSettings: &SocketConfig{Mark: 255}}

	fieldRule := []byte(`{"type":"field",
		"domain": ["geosite:cn"],
		"ip":["geoip:cn", "geoip:private"],
		"outboundTag":"direct"
	}`)

	routerConfig := &RouterConfig{}
	routerConfig.Settings = &RouterRulesConfig{DomainStrategy: "IPIfNonMatch", RuleList: []json.RawMessage{fieldRule}}

	serverOutboundConfig := &OutboundConfig{Settings: []byte("{}"), StreamSetting: &StreamConfig{SocketSettings: &SocketConfig{Mark: 255}}}

	v2ray.DNSConfig = &DnsConfig{Servers: Conf.DnsServers}
	v2ray.InboundConfigList = append(v2ray.InboundConfigList, dokodemoConfig)
	v2ray.OutboundConfigList = append(v2ray.OutboundConfigList, serverOutboundConfig)
	v2ray.OutboundConfigList = append(v2ray.OutboundConfigList, freedomConfig)
	v2ray.RouterConfig = routerConfig

	return v2ray
}
