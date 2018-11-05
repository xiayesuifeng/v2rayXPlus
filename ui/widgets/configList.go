package widgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"io/ioutil"
	"strings"
)

type ConfigList struct {
	widgets.QWidget

	vboxLayout *widgets.QVBoxLayout

	buttonGroup *widgets.QButtonGroup

	_ func() `constructor:"init"`

	_ func(name string) `signal:"configChange"`
	_ func(name string) `signal:"editConfig"`
	_ func(name string) `signal:"removeConfig"`
}

func (ptr *ConfigList) init() {
	ptr.vboxLayout = widgets.NewQVBoxLayout2(ptr)
	ptr.vboxLayout.SetSpacing(0)

	ptr.buttonGroup = widgets.NewQButtonGroup(ptr)

	infos, err := ioutil.ReadDir(conf.V2rayConfigPath)
	if err == nil {
		for _, info := range infos {
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
				name := strings.Split(info.Name(), ".json")[0]
				tmp := NewConfigListItem2(name, ptr)
				tmp.ConnectEditConfig(ptr.EditConfig)
				tmp.ConnectRemoveConfig(ptr.RemoveConfig)
				ptr.vboxLayout.AddWidget(tmp, 0, core.Qt__AlignCenter)
				ptr.buttonGroup.AddButton(tmp, 0)
			}
		}
	}

	ptr.buttonGroup.Buttons()[0].SetChecked(true)

	ptr.SetLayout(ptr.vboxLayout)

	ptr.initConnect()
}

func (ptr *ConfigList) initConnect() {
	ptr.buttonGroup.ConnectButtonClicked(func(button *widgets.QAbstractButton) {
		ptr.ConfigChange(button.Text())
	})
}
