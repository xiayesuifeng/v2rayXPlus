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
  background-color: rgba(255, 255, 255, 0.3);
}

QPushBUtton::pressed {
  border-image: url(":/resources/outline-edit-24px.svg");
  background-color: rgba(255, 255, 255, 0.1);
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
  background-color: rgba(255, 255, 255, 0.3);
}

QPushBUtton::pressed {
  border-image: url(":/resources/outline-close-24px.svg");
  background-color: rgba(255, 255, 255, 0.1);
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
  background-image: url(":/resources/outline-done-24px.svg") 2 2 2 2 repeat repeat;
  background-color: rgb(255, 255, 255);
}

QPushButton::hover {
  background-image: url(":/resources/outline-done-24px.svg");
  background-color: rgba(255, 255, 255, 0.3);
}

QPushBUtton::pressed {
  border-image: url(":/resources/outline-done-24px.svg");
  background-image: rgba(255, 255, 255, 0.5);
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
  background-color: rgba(255, 255, 255, 0.3);
}

QPushBUtton::pressed {
  background-image: url(":/resources/outline-close-24px.svg");
  background-color: rgba(255, 255, 255, 0.5);
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
  background-color: rgba(255, 255, 255, 0.5);
}

QPushBUtton::pressed {
  background-color: rgba(255, 255, 255, 0.3);
}`
