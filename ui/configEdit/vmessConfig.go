package configEdit

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type VmessConfig struct {
	*widgets.QFrame

	formLayout *widgets.QFormLayout

	uuidLineEdit    *widgets.QLineEdit
	alterIdLineEdit *widgets.QLineEdit
}

func NewVmessConfig(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *VmessConfig {
	frame := widgets.NewQFrame(parent, fo)

	vmessConfig := &VmessConfig{QFrame: frame}
	vmessConfig.init()

	return vmessConfig
}

func (ptr *VmessConfig) init() {
	ptr.formLayout = widgets.NewQFormLayout(ptr)

	ptr.uuidLineEdit = widgets.NewQLineEdit(ptr)
	ptr.alterIdLineEdit = widgets.NewQLineEdit2("64", ptr)

	ptr.formLayout.AddRow3("用户uuid", ptr.uuidLineEdit)
	ptr.formLayout.AddRow3("alterId", ptr.alterIdLineEdit)

	ptr.SetLayout(ptr.formLayout)
}
