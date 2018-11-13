package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	core2 "gitlab.com/xiayesuifeng/v2rayxplus/core"
	"gitlab.com/xiayesuifeng/v2rayxplus/styles"
	"gitlab.com/xiayesuifeng/v2rayxplus/ui/configEdit"
	widgets2 "gitlab.com/xiayesuifeng/v2rayxplus/ui/widgets"
)

type MainFrame struct {
	*widgets.QFrame

	hboxLayout *widgets.QHBoxLayout

	configList *widgets2.ConfigList

	startButton *widgets.QPushButton

	configEdit *configEdit.ConfigEdit

	configName string
}

func NewMainFrame(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *MainFrame {
	frame := &MainFrame{QFrame: widgets.NewQFrame(parent, fo)}

	frame.init()
	frame.initConnect()

	return frame
}

func (m *MainFrame) init() {
	m.hboxLayout = widgets.NewQHBoxLayout2(m)
	m.hboxLayout.SetSpacing(0)
	m.hboxLayout.SetContentsMargins(0, 0, 0, 0)

	vboxLayout := widgets.NewQVBoxLayout2(m)

	m.configList = widgets2.NewConfigList(m, 0)

	m.startButton = widgets.NewQPushButton(m)
	m.startButton.SetFixedSize2(230, 230)
	if exited, _ := core2.StatusService("v2ray"); exited {
		m.startButton.SetStyleSheet(styles.GetStyleSheet(styles.StopButton))
		m.startButton.SetWindowTitle("on")
	} else {
		m.startButton.SetStyleSheet(styles.GetStyleSheet(styles.StartButton))
	}

	vboxLayout.AddSpacing(60)
	vboxLayout.AddWidget(m.startButton, 0, core.Qt__AlignHCenter)
	vboxLayout.AddSpacing(48)
	var versionLabel *widgets.QLabel
	if version, err := core2.GetVension(); err != nil {
		versionLabel = widgets.NewQLabel2("v2ray版本: 未安装", m, 0)
	} else {
		versionLabel = widgets.NewQLabel2("v2ray版本: "+version, m, 0)
	}

	vboxLayout.AddWidget(versionLabel, 0, core.Qt__AlignHCenter)
	vboxLayout.AddWidget(m.configList, 1, core.Qt__AlignBottom)

	vboxLayout.AddSpacing(30)

	m.configEdit = configEdit.NewConfigEdit(m, 0)

	m.hboxLayout.AddLayout(vboxLayout, 0)
	m.hboxLayout.AddWidget(m.configEdit, 0, core.Qt__AlignRight)

	m.SetLayout(m.hboxLayout)
}

func (m *MainFrame) initConnect() {
	m.startButton.ConnectClicked(func(checked bool) {
		if m.startButton.WindowTitle() == "on" {
			if core2.StopV2rayXPlusSerive(m.configList.ConfigName) {
				m.startButton.SetWindowTitle("off")
				m.startButton.SetStyleSheet(styles.GetStyleSheet(styles.StartButton))
			}
		} else {
			if core2.StartV2rayXPlusSerive(m.configList.ConfigName) {
				m.startButton.SetWindowTitle("on")
				m.startButton.SetStyleSheet(styles.GetStyleSheet(styles.StopButton))
			}
		}
	})

	m.configList.ConnectConfigChange(m.configEdit.ConfigChange)
	m.configList.ConnectConfigChange(func(name string) {
		if m.startButton.WindowTitle() == "on" {
			if name != m.configName {
				core2.StopV2rayXPlusSerive(m.configName)
				core2.StartV2rayXPlusSerive(name)
			}
			m.configName = name
		}
	})

	m.configList.ConnectEditConfig(m.configEdit.EditChange)

	m.configList.ConnectEditConfig(func(name string) {
		m.ParentWidget().SetFixedSize(core.NewQSize2(1050, 600))
	})
}
