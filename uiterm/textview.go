package uiterm

import (
	_ "strings"
	_ "unicode/utf8"

	"github.com/jroimartin/gocui"
)

type TextboxView struct {
	Name string
	X    int
	Y    int
	W    int
	H    int
	text string
}

func (t *TextboxView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(t.Name, t.X, maxY-t.Y, maxX, maxY+1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Editable = true
	v.Clear()
	return nil
}
