package widgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/styles"
)

type ConfigListItem struct {
	widgets.QPushButton

	toolFrame *widgets.QFrame

	editButton   *widgets.QPushButton
	removeButton *widgets.QPushButton

	_ func() `constructor:"init"`

	_ func(name string) `signal:"editConfig"`
	_ func(name string) `signal:"removeConfig"`
}

func (ptr *ConfigListItem) init() {
	ptr.SetFixedSize2(300, 40)
	ptr.SetCheckable(true)
	ptr.SetStyleSheet(styles.GetStyleSheet(styles.ConfigListItem))
	ptr.SetFocusPolicy(core.Qt__NoFocus)

	ptr.toolFrame = widgets.NewQFrame(ptr, 0)
	ptr.toolFrame.SetVisible(false)
	ptr.toolFrame.SetGeometry2(0, 0, 300, 40)
	toolLayout := widgets.NewQHBoxLayout2(ptr.toolFrame)
	toolLayout.SetContentsMargins(6, 6, 6, 6)
	toolLayout.SetSpacing(6)

	ptr.editButton = widgets.NewQPushButton(ptr.toolFrame)
	ptr.editButton.SetFixedSize2(24, 24)
	ptr.editButton.SetStyleSheet(styles.GetStyleSheet(styles.EditButton))

	ptr.removeButton = widgets.NewQPushButton(ptr.toolFrame)
	ptr.removeButton.SetFixedSize2(24, 24)
	ptr.removeButton.SetStyleSheet(styles.GetStyleSheet(styles.RemoveButton))

	toolLayout.AddStretch(1)
	toolLayout.AddWidget(ptr.editButton, 0, core.Qt__AlignRight)
	toolLayout.AddWidget(ptr.removeButton, 0, core.Qt__AlignRight)

	ptr.initConnect()
}

func (ptr *ConfigListItem) initConnect() {
	ptr.ConnectEnterEvent(func(event *core.QEvent) {
		ptr.toolFrame.SetVisible(true)
	})

	ptr.ConnectLeaveEvent(func(event *core.QEvent) {
		ptr.toolFrame.SetVisible(false)
	})

	ptr.editButton.ConnectClicked(func(checked bool) {
		ptr.EditConfig(ptr.Text())
	})

	ptr.removeButton.ConnectClicked(func(checked bool) {
		ptr.RemoveConfig(ptr.Text())
	})
}
