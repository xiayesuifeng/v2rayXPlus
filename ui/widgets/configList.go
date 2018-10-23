package widgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"io/ioutil"
	"strings"
)

type ConfigList struct {
	*widgets.QWidget

	vboxLayout *widgets.QVBoxLayout
}

func NewConfigList(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *ConfigList {
	widget := widgets.NewQWidget(parent, fo)

	configList := &ConfigList{QWidget: widget}
	configList.init()

	return configList
}

func (ptr *ConfigList) init() {
	ptr.vboxLayout = widgets.NewQVBoxLayout2(ptr)

	infos, err := ioutil.ReadDir(conf.ConfigPath)
	if err == nil {
		for _, info := range infos {
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
				name := strings.Split(info.Name(), ".json")[0]
				tmp := NewConfigListItem(name, ptr)
				ptr.vboxLayout.AddWidget(tmp, 0, core.Qt__AlignCenter)
			}
		}
	}

	ptr.SetLayout(ptr.vboxLayout)
}
