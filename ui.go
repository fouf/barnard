package barnard

import (
	"fmt"
	"os"
	_ "strings"
	_ "time"

	"github.com/jroimartin/gocui"
	"github.com/kennygrant/sanitize"
	"github.com/layeh/barnard/uiterm"
)

// TODO Sanitation ..?

func esc(str string) string {
	return sanitize.HTML(str)
}

func (b *Barnard) InitializeUI() error {
	var err error
	b.UI, err = gocui.NewGui()
	if err != nil {
		return err
	}

	b.UILeftPane = uiterm.LeftPane{
		Name:  "LeftPane",
		Title: "Barnard",
		X:     0,
		Y:     0,
		W:     0.7,
		H:     0.7,
	}
	b.UIRightPane = uiterm.RightPane{
		Name:       "RightPane",
		Title:      "Tx",
		UILeftPane: &b.UILeftPane,
	}
	b.UIBottomPane = uiterm.BottomPane{
		Name:       "BottomPane",
		Title:      "log",
		UILeftPane: &b.UILeftPane,
	}
	b.UITextboxView = uiterm.TextboxView{
		Name: "TextboxView",
		X:    0,
		Y:    3,
		H:    4,
	}
	b.UITextboxEntry = uiterm.TextboxEntry{
		Name:         "TextboxEntry",
		Title:        "",
		UIBottomPane: &b.UIBottomPane,
	}
	b.UI.SelFgColor = gocui.ColorGreen
	return nil
}

func (b *Barnard) SetupHotkeys() {
	// Todo, move to own file / customizability!
	if err := b.UI.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return gocui.ErrQuit }); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	if err := b.UI.SetKeybinding("", gocui.KeyCtrlA, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		b.UI.SetCurrentView("TextboxEntry")
		b.UI.SetViewOnTop("TextboxEntry")
		b.UI.Cursor = true
		return nil
	}); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

}

func (b *Barnard) UpdateInputStatus(status string) {

}
