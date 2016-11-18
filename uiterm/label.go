package uiterm

import (
	"github.com/jroimartin/gocui"
)

type LabelWidget struct {
	name string
	text string
	x    int
	y    int
	w    int
	h    int
}

func (l *LabelWidget) Layout(g *gocui.Gui) error {

	return nil
}
