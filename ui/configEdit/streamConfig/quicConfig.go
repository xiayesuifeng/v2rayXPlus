package streamConfig

import (
	"encoding/json"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
)

type QuicConfig struct {
	*widgets.QWidget

	key *widgets.QLineEdit

	securityComboBox *widgets.QComboBox
	typeComboBox     *widgets.QComboBox
}

func NewQuicConfig(parent widgets.QWidget_ITF) *QuicConfig {
	widget := widgets.NewQWidget(parent, 0)

	quicConfig := &QuicConfig{QWidget: widget}
	quicConfig.init()

	return quicConfig
}

func (ptr *QuicConfig) init() {
	formLayout := widgets.NewQFormLayout(ptr)
	formLayout.SetContentsMargins(0, 0, 0, 0)

	ptr.key = widgets.NewQLineEdit(ptr)
	ptr.key.SetPlaceholderText("可以是任意字符串。当security不为\"none\"时有效")

	ptr.securityComboBox = widgets.NewQComboBox(ptr)
	ptr.securityComboBox.AddItems([]string{"none", "aes-128-gcm", "chacha20-poly1305"})

	ptr.typeComboBox = widgets.NewQComboBox(ptr)
	ptr.typeComboBox.AddItems([]string{"none", "srtp", "utp", "wechat-video", "dtls", "wireguard"})

	formLayout.AddRow3("加密方式(security)", ptr.securityComboBox)
	formLayout.AddRow3("key(密钥)", ptr.key)
	formLayout.AddRow3("伪装类型(header type)", ptr.typeComboBox)

	ptr.SetLayout(formLayout)
}

func (ptr *QuicConfig) saveConfig() *conf.QuicConfig {
	quicConfig := conf.QuicConfig{}

	quicConfig.Key = ptr.key.Text()

	if ptr.securityComboBox.CurrentText() != "none" {
		quicConfig.Security = ptr.securityComboBox.CurrentText()
	}

	if ptr.typeComboBox.CurrentText() != "none" {
		quicConfig.Header = &conf.HeaderConfig{Type: ptr.typeComboBox.CurrentText()}
	}

	json, _ := json.Marshal(&quicConfig)
	if string(json) == "{}" {
		return nil
	}

	return &quicConfig
}
