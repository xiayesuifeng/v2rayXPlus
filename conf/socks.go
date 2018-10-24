package conf

type SocksAccount struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}

type SocksServerConfig struct {
	AuthMethod string          `json:"auth"`
	Accounts   []*SocksAccount `json:"accounts"`
	UDP        bool            `json:"udp"`
	Host       string          `json:"ip"`
	Timeout    uint32          `json:"timeout"`
	UserLevel  uint32          `json:"userLevel"`
}
