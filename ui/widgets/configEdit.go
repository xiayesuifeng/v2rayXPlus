package widgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/styles"
)

type ConfigEdit struct {
	*widgets.QFrame

	vboxLayout *widgets.QVBoxLayout

	serviceEdit *widgets.QLineEdit
	portEdit    *widgets.QLineEdit

	protocolComboBox *widgets.QComboBox

	saveButton   *widgets.QPushButton
	cancelButton *widgets.QPushButton

	confName string
}

func NewConfigEdit(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *ConfigEdit {
	frame := widgets.NewQFrame(parent, fo)

	configEdit := &ConfigEdit{QFrame: frame}
	configEdit.init()
	configEdit.initConnect()

	return configEdit
}

func (ptr *ConfigEdit) init() {
	ptr.SetMinimumSize2(700, 600)
	ptr.SetVisible(false)

	ptr.vboxLayout = widgets.NewQVBoxLayout2(ptr)
	ptr.vboxLayout.SetContentsMargins(20, 20, 20, 20)

	baseConfigLayout := widgets.NewQFormLayout(ptr)

	ptr.serviceEdit = widgets.NewQLineEdit(ptr)
	ptr.portEdit = widgets.NewQLineEdit(ptr)

	ptr.protocolComboBox = widgets.NewQComboBox(ptr)
	ptr.protocolComboBox.AddItems([]string{"shadowsocks", "vemss", "socks"})

	baseConfigLayout.AddRow3("服务器", ptr.serviceEdit)
	baseConfigLayout.AddRow3("端口", ptr.portEdit)
	baseConfigLayout.AddRow3("协议", ptr.protocolComboBox)

	actionLayout := widgets.NewQHBoxLayout2(ptr)
	actionLayout.SetSpacing(20)

	ptr.saveButton = widgets.NewQPushButton(ptr)
	ptr.cancelButton = widgets.NewQPushButton(ptr)

	ptr.saveButton.SetFixedSize2(45, 45)
	ptr.cancelButton.SetFixedSize2(45, 45)

	ptr.saveButton.SetStyleSheet(styles.SaveButton)
	ptr.cancelButton.SetStyleSheet(styles.CancelButton)

	actionLayout.AddStretch(1)
	actionLayout.AddWidget(ptr.saveButton, 0, core.Qt__AlignCenter)
	actionLayout.AddWidget(ptr.cancelButton, 0, core.Qt__AlignCenter)
	actionLayout.AddStretch(1)

	ptr.vboxLayout.AddLayout(baseConfigLayout, 1)
	ptr.vboxLayout.AddStretch(1)
	ptr.vboxLayout.AddLayout(actionLayout, 0)

	ptr.SetLayout(ptr.vboxLayout)
}

func (ptr *ConfigEdit) initConnect() {
	ptr.cancelButton.ConnectClicked(func(checked bool) {
		ptr.SetVisible(false)
		ptr.ParentWidget().ParentWidget().SetFixedSize2(350, 600)
	})

	ptr.protocolComboBox.ConnectCurrentIndexChanged2(func(text string) {

	})
}

func (ptr *ConfigEdit) ConfigChange(name string) {
	if ptr.IsVisible() {
		ptr.parseConfig(name)
	}
}

func (ptr *ConfigEdit) EditChange(name string) {
	ptr.SetVisible(true)
	ptr.parseConfig(name)
}

func (ptr *ConfigEdit) parseConfig(name string) {
	ptr.confName = name
}

func (ptr *ConfigEdit) saveConfig() {
	ptr.SetVisible(false)
	ptr.ParentWidget().ParentWidget().SetFixedSize2(350, 600)
}
