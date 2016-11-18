package uiterm

import (
	"github.com/jroimartin/gocui"
)

// Width and height between 0 and 1, a percentage of the maximum width / height
type LeftPane struct {
	Name  string
	Title string
	X     int
	Y     int
	W     float32
	H     float32
}

func (w *LeftPane) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	v, err := g.SetView(w.Name, w.X, w.Y, int(w.W*float32(maxX)), int(w.H*float32(maxY)))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = w.Title

	return nil
}
