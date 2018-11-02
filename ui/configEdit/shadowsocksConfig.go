package configEdit

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type ShadowsocsConfig struct {
	*widgets.QFrame

	formLayout *widgets.QFormLayout

	emailEdit    *widgets.QLineEdit
	passwordEdit *widgets.QLineEdit

	methodEdit *widgets.QComboBox

	otaCheckBox *widgets.QCheckBox
	udpCheckBox *widgets.QCheckBox
}

func NewShadowsocsConfig(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *ShadowsocsConfig {
	frame := widgets.NewQFrame(parent, fo)

	shadowsocsConfig := &ShadowsocsConfig{QFrame: frame}
	shadowsocsConfig.init()

	return shadowsocsConfig
}

func (ptr *ShadowsocsConfig) init() {
	ptr.formLayout = widgets.NewQFormLayout(ptr)
	ptr.formLayout.SetContentsMargins(0, 0, 0, 0)

	ptr.emailEdit = widgets.NewQLineEdit(ptr)
	ptr.passwordEdit = widgets.NewQLineEdit(ptr)

	ptr.emailEdit.SetPlaceholderText("可选")

	ptr.methodEdit = widgets.NewQComboBox(ptr)
	ptr.otaCheckBox = widgets.NewQCheckBox2("ota", ptr)
	ptr.udpCheckBox = widgets.NewQCheckBox2("udp转发", ptr)

	ptr.methodEdit.AddItems([]string{"aes-256-cfb", "aes-128-cfb", "chacha20", "chacha20-ietf", "aes-256-gcm", "aes-128-gcm", "chacha20-poly1305"})
	ptr.methodEdit.SetCurrentIndexDefault(5)

	ptr.formLayout.AddRow3("邮箱", ptr.emailEdit)
	ptr.formLayout.AddRow3("加密协议", ptr.methodEdit)
	ptr.formLayout.AddRow3("密码", ptr.passwordEdit)
	ptr.formLayout.AddRow5(ptr.otaCheckBox)
	ptr.formLayout.AddRow5(ptr.udpCheckBox)

	ptr.SetLayout(ptr.formLayout)
}
