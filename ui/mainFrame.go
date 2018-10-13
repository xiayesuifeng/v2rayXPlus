package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type MainFrame struct {
	*widgets.QFrame
}

func NewMainFrame(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *MainFrame {
	frame := &MainFrame{QFrame: widgets.NewQFrame(parent, fo)}

	frame.init()
	frame.initConnect()

	return frame
}

func (m *MainFrame) init() {

}

func (m *MainFrame) initConnect() {

}
