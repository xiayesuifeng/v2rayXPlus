package widgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/styles"
)

type ConfigListItem struct {
	*widgets.QPushButton

	toolFrame *widgets.QFrame

	EditButton   *widgets.QPushButton
	RemoveButton *widgets.QPushButton
}

func NewConfigListItem(text string, parent widgets.QWidget_ITF) *ConfigListItem {
	widget := widgets.NewQPushButton2(text, parent)

	configListItem := &ConfigListItem{QPushButton: widget}
	configListItem.init()
	configListItem.initConnect()

	return configListItem
}

func (ptr *ConfigListItem) init() {
	ptr.SetFixedSize2(300, 40)
	ptr.SetCheckable(true)

	ptr.toolFrame = widgets.NewQFrame(ptr, 0)
	ptr.toolFrame.SetVisible(false)
	ptr.toolFrame.SetGeometry2(0, 0, 300, 40)
	toolLayout := widgets.NewQHBoxLayout2(ptr.toolFrame)
	toolLayout.SetContentsMargins(6, 6, 6, 6)
	toolLayout.SetSpacing(6)
	toolLayout.SetAlign(core.Qt__AlignRight)

	ptr.EditButton = widgets.NewQPushButton(ptr.toolFrame)
	ptr.EditButton.SetFixedSize2(24, 24)
	ptr.EditButton.SetStyleSheet(styles.EditButton)

	ptr.RemoveButton = widgets.NewQPushButton(ptr.toolFrame)
	ptr.RemoveButton.SetFixedSize2(24, 24)
	ptr.RemoveButton.SetStyleSheet(styles.RemoveButton)

	toolLayout.AddWidget(ptr.EditButton, 0, core.Qt__AlignRight)
	toolLayout.AddWidget(ptr.RemoveButton, 0, core.Qt__AlignRight)
}

func (ptr *ConfigListItem) initConnect() {
	ptr.ConnectEnterEvent(func(event *core.QEvent) {
		ptr.toolFrame.SetVisible(true)
	})

	ptr.ConnectLeaveEvent(func(event *core.QEvent) {
		ptr.toolFrame.SetVisible(false)
	})
}
