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
