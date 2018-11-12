package conf

var ConfigPath string
var V2rayConfigPath string

var Conf *Config

type Config struct {
	Theme      string   `json:"theme"`
	ListerPort int      `json:"lister_port"`
	DnsServers []string `json:"dnsServers"`
}
