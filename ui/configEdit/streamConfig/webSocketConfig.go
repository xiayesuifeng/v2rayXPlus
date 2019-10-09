package streamConfig

import (
	"github.com/therecipe/qt/widgets"
)

type WebSocketConfig struct {
	*widgets.QWidget

	pathLineEdit *widgets.QLineEdit
}

func NewWebSocketConfig(parent widgets.QWidget_ITF) *WebSocketConfig {
	widget := widgets.NewQWidget(parent, 0)

	webSocketConfig := &WebSocketConfig{QWidget: widget}
	webSocketConfig.init()

	return webSocketConfig
}

func (ptr *WebSocketConfig) init() {
	formLayout := widgets.NewQFormLayout(ptr)
	formLayout.SetContentsMargins(0, 0, 0, 0)

	ptr.pathLineEdit = widgets.NewQLineEdit(ptr)

	ptr.pathLineEdit.SetPlaceholderText("/")

	formLayout.AddRow3("path", ptr.pathLineEdit)

	ptr.SetLayout(formLayout)
}
