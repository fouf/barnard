package uiterm

import (
	"github.com/jroimartin/gocui"
)

type BottomPane struct {
	Name       string
	Title      string
	UILeftPane *LeftPane
}

func (w *BottomPane) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	v, err := g.SetView(w.Name, 0, int(w.UILeftPane.H*float32(maxY)), maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = w.Title

	return nil
}
