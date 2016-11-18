package uiterm

import (
	"github.com/jroimartin/gocui"
)

func badLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if _, err := g.SetView("right_pane", int(0.7*float32(maxX)), 0, maxX-1, int(0.7*float32(maxY))); err != nil && err != gocui.ErrUnknownView {
		return err
	}

	if _, err := g.SetView("bottom_pane", -1, int(0.7*float32(maxY)), maxX, maxY); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	return nil
}
