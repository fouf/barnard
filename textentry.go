package main

import (
	"github.com/jroimartin/gocui"
)

type TextboxEntry struct {
	Name         string
	UIBottomPane *BottomPane
	Title        string
	SendMessage  func(text string)
}

func (w *TextboxEntry) UpdateInputStatus(g *gocui.Gui, status string) {
	w.Title = status
	g.Execute(w.Layout)
}

func (w *TextboxEntry) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(w.Name, 0, maxY-3, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Editable = true
	v.Title = w.Title

	return nil
}

func (w *TextboxEntry) HandleSendMessageKey(g *gocui.Gui, v *gocui.View) error {
	w.SendMessage(v.Buffer())
	v.Clear()
	return nil
}
