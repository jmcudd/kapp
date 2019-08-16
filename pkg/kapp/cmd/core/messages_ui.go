package core

import (
	"time"

	"github.com/cppforlife/go-cli-ui/ui"
)

type MessagesUI interface {
	NotifySection(msg string, args ...interface{})
	Notify(msgs []string)
}

type PlainMessagesUI struct {
	ui ui.UI
}

var _ MessagesUI = &PlainMessagesUI{}

func NewPlainMessagesUI(ui ui.UI) *PlainMessagesUI {
	return &PlainMessagesUI{ui: ui}
}

func (ui *PlainMessagesUI) NotifySection(msg string, args ...interface{}) {
	ui.notify("---- "+msg+" ----", args...)
}

func (ui *PlainMessagesUI) Notify(msgs []string) {
	for _, msg := range msgs {
		ui.notify("%s", msg)
	}
}

func (ui *PlainMessagesUI) notify(msg string, args ...interface{}) {
	ui.ui.BeginLinef(time.Now().Format("3:04:05PM")+": "+msg+"\n", args...)
}
