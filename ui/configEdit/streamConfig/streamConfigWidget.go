package streamConfig

import (
	"github.com/therecipe/qt/widgets"
)

type StreamConfigWidget struct {
	*widgets.QWidget

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
