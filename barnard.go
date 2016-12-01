package main

import (
	"crypto/tls"
	"net"

	"github.com/jroimartin/gocui"
	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleopenal"
)

type Barnard struct {
	Config        *gumble.Config
	Client        *gumble.Client
	Address       string
	TLSConfig     tls.Config
	BarnardConfig Config
	Stream        *gumbleopenal.Stream

	UI             *gocui.Gui
	UILeftPane     *LeftPane
	UIRightPane    *RightPane
	UIBottomPane   *BottomPane
	UITextboxEntry *TextboxEntry
	ViewArray      []string
	ViewActive     int
	Listener       net.Listener
}
