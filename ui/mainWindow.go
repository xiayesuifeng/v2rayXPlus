package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type MainWindow struct {
	*widgets.QMainWindow
}

func NewMainWindow() *MainWindow {
	w := widgets.NewQMainWindow(nil, 0)

	m := &MainWindow{w}
	m.init()

	return m
}

func (m *MainWindow) init() {
	m.SetWindowTitle("v2rayXPlus")
	m.SetFixedSize(core.NewQSize2(350, 600))

	frame := NewMainFrame(m, 0)
	m.SetCentralWidget(frame)
}
