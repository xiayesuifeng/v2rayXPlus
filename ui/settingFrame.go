package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"strconv"
	"strings"
)

type SettingFrame struct {
	*widgets.QFrame
}

func NewSettingFrame(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *SettingFrame {
	frame := widgets.NewQFrame(parent, fo)

	settingFrame := &SettingFrame{QFrame: frame}
	settingFrame.init()

	return settingFrame
}

func (ptr *SettingFrame) init() {
	ptr.SetWindowTitle("设置")

	vboxLayout := widgets.NewQVBoxLayout2(ptr)

	baseLabel := widgets.NewQLabel2("基本", ptr, 0)

	scrollArea := widgets.NewQScrollArea(ptr)
	scrollLayout := widgets.NewQFormLayout(scrollArea)

	portEdit := widgets.NewQLineEdit2(strconv.FormatInt(int64(conf.Conf.ListerPort), 10), ptr)

	dnsEdit := widgets.NewQTextEdit2(strings.Join(conf.Conf.DnsServers, ","), ptr)
	dnsEdit.SetFixedHeight(64)

	scrollLayout.AddRow3("监听端口:", portEdit)
	scrollLayout.AddRow3("DNS服务器:", dnsEdit)

	vboxLayout.AddWidget(baseLabel, 0, core.Qt__AlignLeft)
	vboxLayout.AddWidget(scrollArea, 1, 0)

	ptr.SetLayout(vboxLayout)
}
