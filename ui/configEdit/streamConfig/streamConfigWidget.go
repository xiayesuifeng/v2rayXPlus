package streamConfig

import (
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"strconv"
	"strings"
)

type StreamConfigWidget struct {
	*widgets.QWidget

	streamConfig *conf.StreamConfig

	networkComboBox  *widgets.QComboBox
	securityComboBox *widgets.QComboBox

	stackedLayout *widgets.QStackedLayout

	tcpConfig          *TcpConfig
	kcpConfig          *KcpConfig
	webSocketConfig    *WebSocketConfig
	httpConfig         *HttpConfig
	domainSocketConfig *DomainSocketConfig
	quicConfig         *QuicConfig
}

func NewStreamConfigWidget(parent widgets.QWidget_ITF) *StreamConfigWidget {
	widget := widgets.NewQWidget(parent, 0)

	transportConfigWidget := &StreamConfigWidget{QWidget: widget}
	transportConfigWidget.init()
	transportConfigWidget.initConnect()

	return transportConfigWidget
}

func (ptr *StreamConfigWidget) init() {
	vboxLayout := widgets.NewQVBoxLayout2(ptr)

	formLayout := widgets.NewQFormLayout(ptr)

	ptr.networkComboBox = widgets.NewQComboBox(ptr)
	ptr.networkComboBox.AddItems([]string{"tcp", "kcp", "ws", "http", "domainsocket", "quic"})

	ptr.securityComboBox = widgets.NewQComboBox(ptr)
	ptr.securityComboBox.AddItems([]string{"none", "tls"})

	ptr.tcpConfig = NewTcpConfig(ptr)
	ptr.kcpConfig = NewKcpConfig(ptr)
	ptr.webSocketConfig = NewWebSocketConfig(ptr)
	ptr.httpConfig = NewHttpConfig(ptr)
	ptr.domainSocketConfig = NewDomainSocketConfig(ptr)
	ptr.quicConfig = NewQuicConfig(ptr)

	ptr.stackedLayout = widgets.NewQStackedLayout2(ptr)
	ptr.stackedLayout.AddWidget(ptr.tcpConfig)
	ptr.stackedLayout.AddWidget(ptr.kcpConfig)
	ptr.stackedLayout.AddWidget(ptr.webSocketConfig)
	ptr.stackedLayout.AddWidget(ptr.httpConfig)
	ptr.stackedLayout.AddWidget(ptr.domainSocketConfig)
	ptr.stackedLayout.AddWidget(ptr.quicConfig)

	formLayout.AddRow3("network:", ptr.networkComboBox)
	formLayout.AddRow3("security:", ptr.securityComboBox)

	vboxLayout.AddLayout(formLayout, 0)
	vboxLayout.AddLayout(ptr.stackedLayout, 0)

	ptr.SetLayout(vboxLayout)
}

func (ptr *StreamConfigWidget) initConnect() {
	ptr.networkComboBox.ConnectCurrentIndexChanged(ptr.stackedLayout.SetCurrentIndex)
}

func (ptr *StreamConfigWidget) ParseConfig(config *conf.OutboundConfig) {
	ptr.streamConfig = config.StreamSetting

	ptr.networkComboBox.SetCurrentText(ptr.streamConfig.Network)
	switch ptr.streamConfig.Network {
	case "tcp":
		ptr.tcpConfig.tcpSettingsJsonEdit.SetPlainText(ptr.streamConfig.TcpSettings)
	case "kcp":
		if ptr.streamConfig.KcpSettings != nil {
			ptr.kcpConfig.mtuLineEdit.SetText(strconv.Itoa(int(ptr.streamConfig.KcpSettings.Mtu)))
			ptr.kcpConfig.ttiLineEdit.SetText(strconv.Itoa(int(ptr.streamConfig.KcpSettings.Tti)))
			ptr.kcpConfig.uplinkCapacityLineEdit.SetText(strconv.Itoa(int(ptr.streamConfig.KcpSettings.UplinkCapacity)))
			ptr.kcpConfig.downlinkCapacityLineEdit.SetText(strconv.Itoa(int(ptr.streamConfig.KcpSettings.DownlinkCapacity)))
			ptr.kcpConfig.readBufferSizeLineEdit.SetText(strconv.Itoa(int(ptr.streamConfig.KcpSettings.ReadBufferSize)))
			ptr.kcpConfig.writeBufferSizeLineEdit.SetText(strconv.Itoa(int(ptr.streamConfig.KcpSettings.WriteBufferSize)))
			ptr.kcpConfig.congestionCheckBox.SetChecked(ptr.streamConfig.KcpSettings.Congestion)
			if ptr.streamConfig.KcpSettings.Header != nil {
				ptr.kcpConfig.typeComboBox.SetCurrentText(ptr.streamConfig.KcpSettings.Header.Type)
			}
		}
	case "ws":
		if ptr.streamConfig.WsSettings != nil {
			ptr.webSocketConfig.pathLineEdit.SetText(ptr.streamConfig.WsSettings.Path)
		}
	case "http":
		if ptr.streamConfig.HttpSettings != nil {
			ptr.httpConfig.pathLineEdit.SetText(ptr.streamConfig.HttpSettings.Path)
			ptr.httpConfig.hostTextEdit.SetPlainText(strings.Join(ptr.streamConfig.HttpSettings.Host, ","))
		}
	case "domainsocket":
		if ptr.streamConfig.DsSettings != nil {
			ptr.domainSocketConfig.pathLineEdit.SetText(ptr.streamConfig.DsSettings.Path)
		}
	case "quic":
		if ptr.streamConfig.QuicSettings != nil {
			ptr.quicConfig.securityComboBox.SetCurrentText(ptr.streamConfig.QuicSettings.Security)
			ptr.quicConfig.key.SetText(ptr.streamConfig.QuicSettings.Key)
			if ptr.streamConfig.QuicSettings.Header != nil {
				ptr.quicConfig.typeComboBox.SetCurrentText(ptr.streamConfig.QuicSettings.Header.Type)
			}
		}
	}
}
