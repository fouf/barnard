package uiterm

import (
	"github.com/jroimartin/gocui"
)

type RightPane struct {
	Name       string
	Title      string
	UILeftPane *LeftPane
}

func (w *RightPane) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	v, err := g.SetView(w.Name, w.UILeftPane.X+int(w.UILeftPane.W*float32(maxX)), 0, maxX-1, int(w.UILeftPane.H*float32(maxY)))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = w.Title

	return nil
}
