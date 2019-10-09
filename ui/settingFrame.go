package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"strconv"
	"strings"
	"unicode"
)

type SettingFrame struct {
	*widgets.QFrame

	themeComboBox *widgets.QComboBox

	portEdit *widgets.QLineEdit
	dnsEdit  *widgets.QPlainTextEdit

	saveButton   *widgets.QPushButton
	cancelButton *widgets.QPushButton
}

func NewSettingFrame(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *SettingFrame {
	frame := widgets.NewQFrame(parent, fo)

	settingFrame := &SettingFrame{QFrame: frame}
	settingFrame.init()
	settingFrame.initConnect()

	return settingFrame
}

func (ptr *SettingFrame) init() {
	ptr.SetWindowTitle("设置")

	vboxLayout := widgets.NewQVBoxLayout2(ptr)

	baseLabel := widgets.NewQLabel2("基本", ptr, 0)

	scrollArea := widgets.NewQScrollArea(ptr)
	scrollLayout := widgets.NewQFormLayout(scrollArea)

	ptr.themeComboBox = widgets.NewQComboBox(ptr)
	ptr.themeComboBox.AddItems([]string{"light", "dark"})

	ptr.portEdit = widgets.NewQLineEdit2(strconv.FormatInt(int64(conf.Conf.ListerPort), 10), ptr)

	ptr.dnsEdit = widgets.NewQPlainTextEdit2(strings.Join(conf.Conf.DnsServers, ",\n"), ptr)
	ptr.dnsEdit.SetFixedHeight(72)

	scrollLayout.AddRow3("主题:", ptr.themeComboBox)
	scrollLayout.AddRow3("监听端口:", ptr.portEdit)
	scrollLayout.AddRow3("DNS服务器:", ptr.dnsEdit)

	hboxLayout := widgets.NewQHBoxLayout2(ptr)

	ptr.saveButton = widgets.NewQPushButton2("保存", ptr)
	ptr.cancelButton = widgets.NewQPushButton2("取消", ptr)

	hboxLayout.AddStretch(1)
	hboxLayout.AddWidget(ptr.saveButton, 0, core.Qt__AlignRight)
	hboxLayout.AddWidget(ptr.cancelButton, 0, core.Qt__AlignRight)

	vboxLayout.AddWidget(baseLabel, 0, core.Qt__AlignLeft)
	vboxLayout.AddWidget(scrollArea, 1, 0)
	vboxLayout.AddLayout(hboxLayout, 0)

	ptr.SetLayout(vboxLayout)
}

func (ptr *SettingFrame) initConnect() {
	ptr.saveButton.ConnectClicked(ptr.saveButtonClicked)

	ptr.cancelButton.ConnectClicked(func(checked bool) {
		ptr.Close()
	})
}

func (ptr *SettingFrame) saveButtonClicked(checked bool) {
	conf.Conf.Theme = ptr.themeComboBox.CurrentText()
	conf.Conf.ListerPort, _ = strconv.Atoi(ptr.portEdit.Text())
	conf.Conf.DnsServers = strings.FieldsFunc(strings.ReplaceAll(ptr.dnsEdit.ToPlainText(), ",", ""), unicode.IsSpace)

	if err := conf.Conf.SaveConf(); err != nil {
		widgets.QMessageBox_Information(ptr, "错误", "配置文件保存失败，错误："+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	} else {
		ptr.Close()
	}
}
