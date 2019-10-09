package streamConfig

import (
	"github.com/therecipe/qt/widgets"
)

type HttpConfig struct {
	*widgets.QWidget

	hostTextEdit *widgets.QPlainTextEdit
	pathLineEdit *widgets.QLineEdit
}

func NewHttpConfig(parent widgets.QWidget_ITF) *HttpConfig {
	widget := widgets.NewQWidget(parent, 0)

	httpConfig := &HttpConfig{QWidget: widget}
	httpConfig.init()

	return httpConfig
}

func (ptr *HttpConfig) init() {
	formLayout := widgets.NewQFormLayout(ptr)
	formLayout.SetContentsMargins(0, 0, 0, 0)

	ptr.hostTextEdit = widgets.NewQPlainTextEdit(ptr)
	ptr.pathLineEdit = widgets.NewQLineEdit(ptr)

	ptr.hostTextEdit.SetPlaceholderText("多个域名用 ',' 分开,如 v2ray.com,v2fly.com")
	ptr.pathLineEdit.SetPlaceholderText("/")

	formLayout.AddRow3("host", ptr.hostTextEdit)
	formLayout.AddRow3("path", ptr.pathLineEdit)

	ptr.SetLayout(formLayout)
}
