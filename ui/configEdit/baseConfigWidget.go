package configEdit

import (
	"encoding/json"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"log"
	"strconv"
)

type BaseConfigWidget struct {
	*widgets.QWidget

	protocolLayout *widgets.QStackedLayout

	serviceEdit *widgets.QLineEdit
	portEdit    *widgets.QLineEdit

	protocolComboBox *widgets.QComboBox

	serverOutboundConfig *conf.OutboundConfig

	shadowsocsConfig *ShadowsocsConfig
	vmessConfig      *VmessConfig
	socksConfig      *SocksConfig
}

func NewBaseConfigWidget(parent widgets.QWidget_ITF) *BaseConfigWidget {
	widget := widgets.NewQWidget(parent, 0)

	baseConfigFrame := &BaseConfigWidget{QWidget: widget}
	baseConfigFrame.init()
	baseConfigFrame.initConnect()

	return baseConfigFrame
}

func (ptr *BaseConfigWidget) init() {
	vboxLayout := widgets.NewQVBoxLayout2(ptr)

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

	vboxLayout.AddLayout(baseConfigLayout, 1)
	vboxLayout.AddLayout(ptr.protocolLayout, 1)

	ptr.SetLayout(vboxLayout)
}

func (ptr *BaseConfigWidget) initConnect() {
	ptr.protocolComboBox.ConnectCurrentIndexChanged(ptr.protocolLayout.SetCurrentIndex)
}

func (ptr *BaseConfigWidget) parseConfig(config *conf.OutboundConfig) {
	ptr.serverOutboundConfig = config

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

func (ptr *BaseConfigWidget) saveConfig() error {
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
	}
	return nil
}
