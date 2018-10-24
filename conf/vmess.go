package conf

type VMessAccount struct {
	ID       string `json:"id"`
	AlterIds uint16 `json:"alterId"`
	Security string `json:"security"`
}

type VMessOutboundTarget struct {
	Address string         `json:"address"`
	Port    uint16         `json:"port"`
	Users   []VMessAccount `json:"users"`
}

type VMessOutboundConfig struct {
	Receivers []*VMessOutboundTarget `json:"vnext"`
}
