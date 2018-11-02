package configEdit

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type SocksConfig struct {
	*widgets.QFrame

	formLayout *widgets.QFormLayout

	userEdit  *widgets.QLineEdit
	passEdit  *widgets.QLineEdit
	levelEdit *widgets.QLineEdit
}

func NewSocksConfig(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *SocksConfig {
	frame := widgets.NewQFrame(parent, fo)

	socksConfig := &SocksConfig{QFrame: frame}
	socksConfig.init()

	return socksConfig
}

func (ptr *SocksConfig) init() {
	ptr.formLayout = widgets.NewQFormLayout(ptr)
	ptr.formLayout.SetContentsMargins(0, 0, 0, 0)

	ptr.userEdit = widgets.NewQLineEdit(ptr)
	ptr.passEdit = widgets.NewQLineEdit(ptr)
	ptr.levelEdit = widgets.NewQLineEdit2("0", ptr)

	ptr.userEdit.SetPlaceholderText("noauth留空")
	ptr.passEdit.SetPlaceholderText("noauth留空")

	ptr.formLayout.AddRow3("帐号", ptr.userEdit)
	ptr.formLayout.AddRow3("密码", ptr.passEdit)
	ptr.formLayout.AddRow3("level", ptr.levelEdit)

	ptr.SetLayout(ptr.formLayout)
}
