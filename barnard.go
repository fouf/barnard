package barnard

import (
	"crypto/tls"

	"github.com/jroimartin/gocui"
	"github.com/layeh/barnard/uiterm"
	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleopenal"
)

type Barnard struct {
	Config    *gumble.Config
	Client    *gumble.Client
	Address   string
	TLSConfig tls.Config

	Stream *gumbleopenal.Stream

	UI             *gocui.Gui
	UILeftPane     uiterm.LeftPane
	UIRightPane    uiterm.RightPane
	UIBottomPane   uiterm.BottomPane
	UITextboxView  uiterm.TextboxView
	UITextboxEntry uiterm.TextboxEntry
}
