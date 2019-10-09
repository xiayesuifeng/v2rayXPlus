package streamConfig

import (
	"github.com/therecipe/qt/widgets"
)

type TcpConfig struct {
	*widgets.QWidget

	tcpSettingsJsonEdit *widgets.QPlainTextEdit
}

func NewTcpConfig(parent widgets.QWidget_ITF) *TcpConfig {
	widget := widgets.NewQWidget(parent, 0)

	tcpConfig := &TcpConfig{QWidget: widget}
	tcpConfig.init()

	return tcpConfig
}

func (ptr *TcpConfig) init() {
	formLayout := widgets.NewQFormLayout(ptr)
	formLayout.SetContentsMargins(0, 0, 0, 0)

	ptr.tcpSettingsJsonEdit = widgets.NewQPlainTextEdit(ptr)
	ptr.tcpSettingsJsonEdit.SetPlaceholderText(`{
    "header": {
        "type": "none"
    }
}`)

	formLayout.AddRow3("tcpSettings", ptr.tcpSettingsJsonEdit)

	ptr.SetLayout(formLayout)
}
