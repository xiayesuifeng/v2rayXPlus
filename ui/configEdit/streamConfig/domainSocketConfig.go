package streamConfig

import (
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
)

type DomainSocketConfig struct {
	*widgets.QWidget

	pathLineEdit *widgets.QLineEdit
}

func NewDomainSocketConfig(parent widgets.QWidget_ITF) *DomainSocketConfig {
	widget := widgets.NewQWidget(parent, 0)

	domainSocketConfig := &DomainSocketConfig{QWidget: widget}
	domainSocketConfig.init()

	return domainSocketConfig
}

func (ptr *DomainSocketConfig) init() {
	formLayout := widgets.NewQFormLayout(ptr)
	formLayout.SetContentsMargins(0, 0, 0, 0)

	ptr.pathLineEdit = widgets.NewQLineEdit(ptr)

	ptr.pathLineEdit.SetPlaceholderText("/path/to/ds/file")

	formLayout.AddRow3("path", ptr.pathLineEdit)

	ptr.SetLayout(formLayout)
}

func (ptr *DomainSocketConfig) saveConfig() *conf.DomainSocketConfig {
	path := ptr.pathLineEdit.Text()

	if path != "" {
		return &conf.DomainSocketConfig{Path: path}
	} else {
		return nil
	}
}
