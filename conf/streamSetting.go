package conf

type StreamConfig struct {
	Network        string              `json:"network,omitempty"`
	Security       string              `json:"security,omitempty"`
	SocketSettings *SocketConfig       `json:"sockopt,omitempty"`
	TcpSettings    string              `json:"tcpSettings,omitempty"`
	KcpSettings    *KcpConfig          `json:"kcpSettings,omitempty"`
	WsSettings     *WebSocketConfig    `json:"wsSettings,omitempty"`
	HttpSettings   *HttpConfig         `json:"httpSettings,omitempty"`
	DsSettings     *DomainSocketConfig `json:"dsSettings,omitempty"`
	QuicSettings   *QuicConfig         `json:"quicSettings,omitempty"`
}

type HeaderConfig struct {
	Type string `json:"type"`
}

type KcpConfig struct {
	Mtu              uint32        `json:"mtu,omitempty"`
	Tti              uint          `json:"tti,omitempty"`
	UplinkCapacity   uint          `json:"uplinkCapacity,omitempty"`
	DownlinkCapacity uint          `json:"downlinkCapacity,omitempty"`
	Congestion       bool          `json:"congestion,omitempty"`
	ReadBufferSize   uint          `json:"readBufferSize,omitempty"`
	WriteBufferSize  uint          `json:"writeBufferSize,omitempty"`
	Header           *HeaderConfig `json:"header,omitempty"`
}

type WebSocketConfig struct {
	Path    string            `json:"path,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

type HttpConfig struct {
	Path string   `json:"path,omitempty"`
	Host []string `json:"host,omitempty"`
}

type DomainSocketConfig struct {
	Path string `json:"path,omitempty"`
}

type QuicConfig struct {
	Security string        `json:"security,omitempty"`
	Key      string        `json:"key,omitempty"`
	Header   *HeaderConfig `json:"header,omitempty"`
}
