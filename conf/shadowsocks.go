package conf

type ShadowsocksServer struct {
	Address  string `json:"address"`
	Port     uint16 `json:"port"`
	Cipher   string `json:"method"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Ota      bool   `json:"ota"`
	Level    byte   `json:"level"`
}

type ShadowsocksConfig struct {
	Servers []ShadowsocksServer `json:"servers"`
}
