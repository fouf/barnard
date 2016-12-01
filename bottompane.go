package main

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

type BottomPane struct {
	Name       string
	Title      string
	UILeftPane *LeftPane
	lines      []string
	Autoscroll bool
}

func (w *BottomPane) AddLine(g *gocui.Gui, line string) {
	w.lines = append(w.lines, line)
	g.Execute(w.Layout)
}
func (w *BottomPane) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(w.Name, 0, int(w.UILeftPane.H*float32(maxY)), maxX-1, maxY-3)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = w.Title
	v.Wrap = true
	if w.Autoscroll {
		v.Autoscroll = true
		w.Autoscroll = false
	}
	v.Clear()
	for _, line := range w.lines {
		fmt.Fprintf(v, "%s", line)
	}
	return nil
}
func (w *BottomPane) Scroll(v *gocui.View, delta int) error {
	_, sY := v.Size()
	oX, oY := v.Origin()
	if oY+delta > strings.Count(v.ViewBuffer(), "\n")-sY-1 {
		v.Autoscroll = true
	} else if oY+delta >= 0 {
		v.Autoscroll = false
		if err := v.SetOrigin(oX, oY+delta); err != nil {
			return err
		}
	}
	return nil
}
func (w *BottomPane) ScrollUp(g *gocui.Gui, v *gocui.View) error {
	if err := w.Scroll(v, -1); err != nil {
		return err
	}
	return nil
}
func (w *BottomPane) ScrollDown(g *gocui.Gui, v *gocui.View) error {
	if err := w.Scroll(v, 1); err != nil {
		return err
	}
	return nil
}
func (w *BottomPane) Clear(g *gocui.Gui, v *gocui.View) error {
	w.lines = w.lines[:0]
	w.Autoscroll = true
	if err := v.SetOrigin(0, 0); err != nil {
		return err
	}
	v.Clear()
	return nil
}
