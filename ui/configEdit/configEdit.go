package configEdit

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"gitlab.com/xiayesuifeng/v2rayxplus/styles"
	"gitlab.com/xiayesuifeng/v2rayxplus/ui/configEdit/streamConfig"
	"log"
	"path"
)

type ConfigEdit struct {
	*widgets.QFrame

	vboxLayout *widgets.QVBoxLayout

	baseConfigWidget   *BaseConfigWidget
	streamConfigWidget *streamConfig.StreamConfigWidget
	muxConfigWidget    *MuxConfigWidget

	saveButton   *widgets.QPushButton
	cancelButton *widgets.QPushButton

	confName string
	conf     *conf.V2rayConfig
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

	ptr.baseConfigWidget = NewBaseConfigWidget(ptr)
	ptr.streamConfigWidget = streamConfig.NewStreamConfigWidget(ptr)
	ptr.muxConfigWidget = NewMuxConfigWidget(ptr)

	tabWidget := widgets.NewQTabWidget(ptr)
	tabWidget.SetMinimumSize2(660, 490)
	tabWidget.AddTab(ptr.baseConfigWidget, "基本")
	tabWidget.AddTab(ptr.streamConfigWidget, "传输配置")
	tabWidget.AddTab(ptr.muxConfigWidget, "Mux 配置")

	actionLayout := widgets.NewQHBoxLayout2(ptr)
	actionLayout.SetSpacing(20)

	ptr.saveButton = widgets.NewQPushButton(ptr)
	ptr.cancelButton = widgets.NewQPushButton(ptr)

	ptr.saveButton.SetFixedSize2(45, 45)
	ptr.cancelButton.SetFixedSize2(45, 45)

	ptr.saveButton.SetStyleSheet(styles.GetStyleSheet(styles.SaveButton))
	ptr.cancelButton.SetStyleSheet(styles.GetStyleSheet(styles.CancelButton))

	actionLayout.AddStretch(1)
	actionLayout.AddWidget(ptr.saveButton, 0, core.Qt__AlignCenter)
	actionLayout.AddWidget(ptr.cancelButton, 0, core.Qt__AlignCenter)
	actionLayout.AddStretch(1)

	ptr.vboxLayout.AddWidget(tabWidget, 10, core.Qt__AlignCenter)
	ptr.vboxLayout.AddLayout(actionLayout, 0)

	ptr.SetLayout(ptr.vboxLayout)
}

func (ptr *ConfigEdit) initConnect() {
	ptr.cancelButton.ConnectClicked(func(checked bool) {
		ptr.SetVisible(false)
		ptr.ParentWidget().ParentWidget().SetFixedSize2(350, 600)
	})

	ptr.saveButton.ConnectClicked(func(checked bool) {
		if err := ptr.saveConfig(); err != nil {
			widgets.QMessageBox_Warning(ptr, "错误", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		}
		ptr.SetVisible(false)
		ptr.ParentWidget().ParentWidget().SetFixedSize2(350, 600)
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

	var err error
	ptr.conf, err = conf.ParseV2ray(path.Join(conf.V2rayConfigPath, name+".json"))
	if err != nil {
		log.Println(err)
	}

	for _, config := range ptr.conf.OutboundConfigList {
		if config.Protocol != "freedom" {
			ptr.baseConfigWidget.parseConfig(config)
			ptr.streamConfigWidget.ParseConfig(config)
			ptr.muxConfigWidget.parseConfig(config)
			break
		}
	}
}

func (ptr *ConfigEdit) saveConfig() error {
	if err := ptr.baseConfigWidget.saveConfig(); err != nil {
		return err
	}

	if err := ptr.streamConfigWidget.SaveConfig(); err != nil {
		return err
	}

	return ptr.conf.Save(path.Join(conf.V2rayConfigPath, ptr.confName+".json"))
}
