package styles

import (
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	_ "gitlab.com/xiayesuifeng/v2rayxplus/resources"
	"strings"
)

const darkColor string = "255,255,255"
const lightColor string = "100,100,100"

const SettingButton string = `
QPushButton[flat="true"] {
  border: 2px;
  outline: none;
  color: #eaeaea;
}

QPushButton {
  border-image: url(":/resources/{theme}/outline-settings-24px.svg");
}

QPushButton::hover {
  border-image: url(":/resources/{theme}/outline-settings-24px.svg");
  background-color: rgba(255, 255, 255, 0.5);
}

QPushButton::pressed {
  border-image: url(":/resources/{theme}/outline-settings-24px.svg");
  background-color: rgba(255, 255, 255, 0.3);
}`

const EditButton string = `
QPushButton[flat="true"] {
  border: 2px;
  outline: none;
  color: #eaeaea;
}

QPushButton {
  border-image: url(":/resources/{theme}/outline-edit-24px.svg");
}

QPushButton::hover {
  border-image: url(":/resources/{theme}/outline-edit-24px.svg");
  background-color: rgba(255, 255, 255, 0.5);
}

QPushButton::pressed {
  border-image: url(":/resources/{theme}/outline-edit-24px.svg");
  background-color: rgba(255, 255, 255, 0.3);
}`

const RemoveButton string = `
QPushButton[flat="true"] {
  border: 2px;
  outline: none;
  color: #eaeaea;
}

QPushButton {
  border-image: url(":/resources/{theme}/outline-close-24px.svg");
}

QPushButton::hover {
  border-image: url(":/resources/{theme}/outline-close-24px.svg");
  background-color: rgba(255, 255, 255, 0.5);
}

QPushButton::pressed {
  border-image: url(":/resources/{theme}/outline-close-24px.svg");
  background-color: rgba(255, 255, 255, 0.3);
}`

const SaveButton string = `
QPushButton {
  border: 2px;
  border-radius: 22px;
  margins: 10px;
  outline: none;
  color: #eaeaea;
  background-position: center;
  background-repeat: no-repeat;
}

QPushButton {
  background-image: url(":/resources/{theme}/outline-done-24px.svg");
  background-color: rgb({color});
}

QPushButton::hover {
  background-image: url(":/resources/{theme}/outline-done-24px.svg");
  background-color: rgba({color}, 0.5);
}

QPushButton::pressed {
  background-image: url(":/resources/{theme}/outline-done-24px.svg");
  background-image: rgba({color}, 0.3);
}`

const CancelButton string = `
QPushButton {
  border: 2px;
  border-radius: 22px;
  outline: none;
  color: #eaeaea;
  background-position: center;
  background-repeat: no-repeat;
}

QPushButton {
  background-image: url(":/resources/{theme}/outline-close-24px.svg");
  background-color: rgb({color});
}

QPushButton::hover {
  background-image: url(":/resources/{theme}/outline-close-24px.svg");
  background-color: rgba({color}, 0.5);
}

QPushButton::pressed {
  background-image: url(":/resources/{theme}/outline-close-24px.svg");
  background-color: rgba({color}, 0.3);
}`

const ConfigListItem string = `
QPushButton {
  border: none;
  outline: none;
  color: #fffffff;
  background-position: left;
  background-repeat: no-repeat;
  border-radius:5px;
}

QPushButton::checked {
  background-image: url(":/resources/{theme}/outline-done-24px.svg");
}

QPushButton::hover {
  background-color: rgba(255, 255, 255, 0.3);
}

QPushButton::pressed {
  background-color: rgba(255, 255, 255, 0.5);
}`

const StartButton string = `
QPushButton {
  border: 2px;
  outline: none;
  color: #eaeaea;
  border-radius: 115px;
  background-color: rgba({color}, 0.5);
}

QPushButton {
  border-image: url(":/resources/rochet_start.png");
}

QPushButton::hover {
  border-image: url(":/resources/rochet_stop.png");
  background-color: #0828f1a6;
}

QPushButton::pressed {
  border-image: url(":/resources/rochet_stop.png");
  background-color: rgba(255, 255, 255, 0.3);
}`

const StopButton string = `
QPushButton {
  border: 2px;
  outline: none;
  color: #eaeaea;
  border-radius: 115px;
  background-color: rgba({color}, 0.5);
}

QPushButton {
  border-image: url(":/resources/rochet_stop.png");
}

QPushButton::hover {
  border-image: url(":/resources/rochet_start.png");
  background-color: #0860d8ff;
}

QPushButton::pressed {
  border-image: url(":/resources/rochet_start.png");
  background-color: rgba(255, 255, 255, 0.3);
}`

const AddButton string = `
QPushButton {
  border: 2px;
  border-radius: 22px;
  outline: none;
  color: #eaeaea;
  background-position: center;
  background-repeat: no-repeat;
}

QPushButton {
  background-image: url(":/resources/{theme}/outline-add-24px.svg");
  background-color: rgb({color});
}

QPushButton::hover {
  background-image: url(":/resources/{theme}/outline-add-24px.svg");
  background-color: rgba({color}, 0.5);
}

QPushButton::pressed {
  background-image: url(":/resources/{theme}/outline-add-24px.svg");
  background-color: rgba({color}, 0.3);
}`

func GetStyleSheet(qss string) string {
	if conf.Conf.Theme == "light" {
		qss = strings.Replace(qss, "{theme}", "dark", -1)
		qss = strings.Replace(qss, "{color}", darkColor, -1)
	} else {
		qss = strings.Replace(qss, "{theme}", "light", -1)
		qss = strings.Replace(qss, "{color}", lightColor, -1)
	}
	return qss
}
