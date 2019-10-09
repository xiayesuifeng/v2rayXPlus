package configEdit

import (
	"encoding/json"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"gitlab.com/xiayesuifeng/v2rayxplus/styles"
	"log"
	"path"
	"strconv"
)

type ConfigEdit struct {
	*widgets.QFrame

	vboxLayout     *widgets.QVBoxLayout
	protocolLayout *widgets.QStackedLayout

	serviceEdit *widgets.QLineEdit
	portEdit    *widgets.QLineEdit

	protocolComboBox *widgets.QComboBox

	saveButton   *widgets.QPushButton
	cancelButton *widgets.QPushButton

	serverOutboundConfig *conf.OutboundConfig

	shadowsocsConfig *ShadowsocsConfig
	vmessConfig      *VmessConfig
	socksConfig      *SocksConfig

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

	baseConfigLayout := widgets.NewQFormLayout(ptr)

	ptr.serviceEdit = widgets.NewQLineEdit(ptr)
	ptr.portEdit = widgets.NewQLineEdit(ptr)

	ptr.protocolComboBox = widgets.NewQComboBox(ptr)
	ptr.protocolComboBox.AddItems([]string{"shadowsocks", "vmess", "socks"})

	baseConfigLayout.AddRow3("服务器", ptr.serviceEdit)
	baseConfigLayout.AddRow3("端口", ptr.portEdit)
	baseConfigLayout.AddRow3("协议", ptr.protocolComboBox)

	ptr.protocolLayout = widgets.NewQStackedLayout2(ptr)

	ptr.shadowsocsConfig = NewShadowsocsConfig(ptr, 0)
	ptr.vmessConfig = NewVmessConfig(ptr, 0)
	ptr.socksConfig = NewSocksConfig(ptr, 0)

	ptr.protocolLayout.AddWidget(ptr.shadowsocsConfig)
	ptr.protocolLayout.AddWidget(ptr.vmessConfig)
	ptr.protocolLayout.AddWidget(ptr.socksConfig)

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

	ptr.vboxLayout.AddLayout(baseConfigLayout, 1)
	ptr.vboxLayout.AddLayout(ptr.protocolLayout, 1)
	ptr.vboxLayout.AddStretch(1)
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

	ptr.protocolComboBox.ConnectCurrentIndexChanged(ptr.protocolLayout.SetCurrentIndex)
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
			ptr.serverOutboundConfig = config
			break
		}
	}

	protocol := ptr.serverOutboundConfig.Protocol
	ptr.protocolComboBox.SetCurrentText(protocol)
	switch protocol {
	case "shadowsocks":
		shadowsocksConf, err := conf.NewShadowsocksServer(ptr.serverOutboundConfig.Settings)
		if err != nil {
			log.Println(err)
		}
		if len(shadowsocksConf.Servers) > 0 {
			ptr.serviceEdit.SetText(shadowsocksConf.Servers[0].Address)
			port := strconv.FormatUint(uint64(shadowsocksConf.Servers[0].Port), 10)
			ptr.portEdit.SetText(port)
			ptr.shadowsocsConfig.ParseConf(shadowsocksConf.Servers[0])
		}
	case "vmess":
		vmessConf, err := conf.NewVMessOutboundConfig(ptr.serverOutboundConfig.Settings)
		if err != nil {
			log.Println(err)
		}
		if len(vmessConf.Receivers) > 0 {
			ptr.serviceEdit.SetText(vmessConf.Receivers[0].Address)
			port := strconv.FormatUint(uint64(vmessConf.Receivers[0].Port), 10)
			ptr.portEdit.SetText(port)
			ptr.vmessConfig.ParseConf(vmessConf.Receivers[0])
		}

	case "socks":
		socksConfig, err := conf.NewSocksClientConfig(ptr.serverOutboundConfig.Settings)
		if err != nil {
			log.Println(err)
		}
		if len(socksConfig.Servers) > 0 {
			ptr.serviceEdit.SetText(socksConfig.Servers[0].Address)
			port := strconv.FormatUint(uint64(socksConfig.Servers[0].Port), 10)
			ptr.portEdit.SetText(port)
			ptr.socksConfig.ParseConf(socksConfig.Servers[0])
		}
	}
}

func (ptr *ConfigEdit) saveConfig() error {
	switch ptr.protocolLayout.CurrentIndex() {
	case 0:
		ptr.serverOutboundConfig.Protocol = "shadowsocks"
		shadowsocksConf, err := conf.NewShadowsocksServer(ptr.serverOutboundConfig.Settings)
		if err != nil {
			return err
		}
		if len(shadowsocksConf.Servers) == 0 {
			shadowsocksConf.Servers = append(shadowsocksConf.Servers, conf.ShadowsocksServer{})
		}
		shadowsocksConf.Servers[0].Address = ptr.serviceEdit.Text()
		port, err := strconv.ParseUint(ptr.portEdit.Text(), 10, 0)
		if err != nil {
			return err
		}
		shadowsocksConf.Servers[0].Port = uint16(port)
		ptr.shadowsocsConfig.SaveConf(&shadowsocksConf.Servers[0])
		settings, err := json.Marshal(shadowsocksConf)
		if err != nil {
			return err
		}
		ptr.serverOutboundConfig.Settings = settings
		return ptr.conf.Save(path.Join(conf.V2rayConfigPath, ptr.confName+".json"))
	case 1:
		ptr.serverOutboundConfig.Protocol = "vmess"
		vmessConf, err := conf.NewVMessOutboundConfig(ptr.serverOutboundConfig.Settings)
		if err != nil {
			return err
		}
		if len(vmessConf.Receivers) == 0 {
			vmessConf.Receivers = append(vmessConf.Receivers, &conf.VMessOutboundTarget{})
		}
		vmessConf.Receivers[0].Address = ptr.serviceEdit.Text()
		port, err := strconv.ParseUint(ptr.portEdit.Text(), 10, 0)
		if err != nil {
			return err
		}
		vmessConf.Receivers[0].Port = uint16(port)
		ptr.vmessConfig.SaveConf(vmessConf.Receivers[0])
		settings, err := json.Marshal(vmessConf)
		if err != nil {
			return err
		}
		ptr.serverOutboundConfig.Settings = settings
		return ptr.conf.Save(path.Join(conf.V2rayConfigPath, ptr.confName+".json"))
	case 2:
		ptr.serverOutboundConfig.Protocol = "socks"
		socksConfig, err := conf.NewSocksClientConfig(ptr.serverOutboundConfig.Settings)
		if err != nil {
			return err
		}
		if len(socksConfig.Servers) == 0 {
			socksConfig.Servers = append(socksConfig.Servers, &conf.SocksRemoteConfig{})
		}
		socksConfig.Servers[0].Address = ptr.serviceEdit.Text()
		port, err := strconv.ParseUint(ptr.portEdit.Text(), 10, 0)
		if err != nil {
			return err
		}
		socksConfig.Servers[0].Port = uint16(port)
		ptr.socksConfig.SaveConf(socksConfig.Servers[0])
		settings, err := json.Marshal(socksConfig)
		if err != nil {
			return err
		}
		ptr.serverOutboundConfig.Settings = settings
		return ptr.conf.Save(path.Join(conf.V2rayConfigPath, ptr.confName+".json"))
	}
	return nil
}
