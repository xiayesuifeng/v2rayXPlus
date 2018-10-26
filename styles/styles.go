package styles

import _ "gitlab.com/xiayesuifeng/v2rayxplus/resources"

const EditButton string = `
QPushButton[flat="true"] {
  border: 2px;
  outline: none;
  color: #eaeaea;
}

QPushButton {
  border-image: url(":/resources/outline-edit-24px.svg");
}

QPushButton::hover {
  border-image: url(":/resources/outline-edit-24px.svg");
  background-color: rgba(255, 255, 255, 0.5);
}

QPushButton::pressed {
  border-image: url(":/resources/outline-edit-24px.svg");
  background-color: rgba(255, 255, 255, 0.3);
}`

const RemoveButton string = `
QPushButton[flat="true"] {
  border: 2px;
  outline: none;
  color: #eaeaea;
}

QPushButton {
  border-image: url(":/resources/outline-close-24px.svg");
}

QPushButton::hover {
  border-image: url(":/resources/outline-close-24px.svg");
  background-color: rgba(255, 255, 255, 0.5);
}

QPushButton::pressed {
  border-image: url(":/resources/outline-close-24px.svg");
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
  background-image: url(":/resources/outline-done-24px.svg");
  background-color: rgb(255, 255, 255);
}

QPushButton::hover {
  background-image: url(":/resources/outline-done-24px.svg");
  background-color: rgba(255, 255, 255, 0.5);
}

QPushButton::pressed {
  background-image: url(":/resources/outline-done-24px.svg");
  background-image: rgba(255, 255, 255, 0.3);
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
  background-image: url(":/resources/outline-close-24px.svg");
  background-color: rgb(255, 255, 255);
}

QPushButton::hover {
  background-image: url(":/resources/outline-close-24px.svg");
  background-color: rgba(255, 255, 255, 0.5);
}

QPushButton::pressed {
  background-image: url(":/resources/outline-close-24px.svg");
  background-color: rgba(255, 255, 255, 0.3);
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
  background-image: url(":/resources/outline-done-24px.svg");
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
  background-color: rgba(255, 255, 255, 0.5);
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
  background-color: rgba(255, 255, 255, 0.5);
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
