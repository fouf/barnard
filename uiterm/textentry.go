package uiterm

import (
	_ "strings"
	_ "unicode/utf8"

	"github.com/jroimartin/gocui"
)

type TextboxEntry struct {
	Name         string
	UIBottomPane *BottomPane
	Title        string
	text         string
}

func (w *TextboxEntry) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(w.Name, 0, maxY-3, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Editable = true
	v.Title = w.Title
	v.Clear()
	return nil
}
