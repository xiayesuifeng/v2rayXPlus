package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	core2 "gitlab.com/xiayesuifeng/v2rayxplus/core"
	widgets2 "gitlab.com/xiayesuifeng/v2rayxplus/ui/widgets"
)

type MainFrame struct {
	*widgets.QFrame

	vboxLayout *widgets.QVBoxLayout

	configList *widgets2.ConfigList

	startButton *widgets.QPushButton
}

func NewMainFrame(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *MainFrame {
	frame := &MainFrame{QFrame: widgets.NewQFrame(parent, fo)}

	frame.init()
	frame.initConnect()

	return frame
}

func (m *MainFrame) init() {
	m.vboxLayout = widgets.NewQVBoxLayout2(m)

	m.configList = widgets2.NewConfigList(m, 0)

	m.startButton = widgets.NewQPushButton2("启动", m)
	m.startButton.SetFixedSize2(230, 230)

	m.vboxLayout.AddSpacing(60)
	m.vboxLayout.AddWidget(m.startButton, 0, core.Qt__AlignHCenter)
	m.vboxLayout.AddSpacing(48)
	var versionLabel *widgets.QLabel
	if version, err := core2.GetVension(); err != nil {
		versionLabel = widgets.NewQLabel2("v2ray版本: 未安装", m, 0)
	} else {
		versionLabel = widgets.NewQLabel2("v2ray版本: "+version, m, 0)
	}

	m.vboxLayout.AddWidget(versionLabel, 0, core.Qt__AlignHCenter)
	m.vboxLayout.AddWidget(m.configList, 1, core.Qt__AlignBottom)

	m.vboxLayout.AddSpacing(60)

	m.SetLayout(m.vboxLayout)
}

func (m *MainFrame) initConnect() {
	m.startButton.ConnectClicked(func(checked bool) {
		if m.startButton.Text() == "启动" {
			m.startButton.SetText("关闭")
		} else if m.startButton.Text() == "关闭" {
			m.startButton.SetText("启动")
		}
	})
}
