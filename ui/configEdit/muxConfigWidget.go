package configEdit

import (
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"strconv"
)

type MuxConfigWidget struct {
	*widgets.QWidget

	enableCheckBox      *widgets.QCheckBox
	concurrencyLineEdit *widgets.QLineEdit

	serverOutboundConfig *conf.OutboundConfig
}

func NewMuxConfigWidget(parent widgets.QWidget_ITF) *MuxConfigWidget {
	widget := widgets.NewQWidget(parent, 0)

	muxConfigWidget := &MuxConfigWidget{QWidget: widget}
	muxConfigWidget.init()

	return muxConfigWidget
}

func (ptr *MuxConfigWidget) init() {
	formLayout := widgets.NewQFormLayout(ptr)

	ptr.enableCheckBox = widgets.NewQCheckBox2("启用", ptr)
	ptr.concurrencyLineEdit = widgets.NewQLineEdit(ptr)

	ptr.concurrencyLineEdit.SetPlaceholderText("8")

	formLayout.AddRow3("Mux", ptr.enableCheckBox)
	formLayout.AddRow3("最大并发连接数(concurrency)", ptr.concurrencyLineEdit)

	ptr.SetLayout(formLayout)
}

func (ptr *MuxConfigWidget) parseConfig(config *conf.OutboundConfig) {
	ptr.serverOutboundConfig = config

	if config.MuxSettings != nil {
		ptr.enableCheckBox.SetChecked(config.MuxSettings.Enabled)

		if config.MuxSettings.Concurrency != 0 {
			ptr.concurrencyLineEdit.SetText(strconv.Itoa(int(config.MuxSettings.Concurrency)))
		}
	}
}
