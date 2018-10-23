package widgets

import (
	"github.com/therecipe/qt/widgets"
)

type ConfigListItem struct {
	*widgets.QPushButton
}

func NewConfigListItem(text string, parent widgets.QWidget_ITF) *ConfigListItem {
	widget := widgets.NewQPushButton2(text, parent)

	configListItem := &ConfigListItem{QPushButton: widget}
	configListItem.init()

	return configListItem
}

func (ptr *ConfigListItem) init() {
	ptr.SetFixedSize2(300, 40)
	ptr.SetCheckable(true)
}
