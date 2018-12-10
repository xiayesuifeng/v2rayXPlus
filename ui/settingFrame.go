package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type SettingFrame struct {
	*widgets.QFrame
}

func NewSettingFrame(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *SettingFrame {
	frame := widgets.NewQFrame(parent, fo)

	settingFrame := &SettingFrame{QFrame: frame}
	settingFrame.init()

	return settingFrame
}

func (ptr *SettingFrame) init() {
	ptr.SetWindowTitle("设置")
}
