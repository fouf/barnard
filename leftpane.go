package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/layeh/gumble/gumble"
)

// Width and height between 0 and 1, a percentage of the maximum width / height
type LeftPane struct {
	Name          string
	Title         string
	X             int
	Y             int
	W             float32
	H             float32
	channels      *gumble.Channels
	client        **gumble.Client
	UIRightPane   **RightPane
	treeStructure []TreeItem
	speaking      bool
}

type TreeItem struct {
	isChannel bool
	isUser    bool
	Y         int
	c         *gumble.Channel
	u         *gumble.User
}

func (w *LeftPane) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	v, err := g.SetView(w.Name, w.X, w.Y, int(w.W*float32(maxX)), int(w.H*float32(maxY)))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = w.Title
	v.Highlight = true
	v.Wrap = true
	v.SelFgColor = gocui.AttrUnderline
	v.Clear()
	w.treeStructure = nil

	if w.channels == nil {
		return nil
	}
	space := "\t\t"
	w.DrawChannelTree(v, (*w.channels)[0], 0, 0, space)
	return nil
}
func (w *LeftPane) DrawChannelTree(v *gocui.View, c *gumble.Channel, offset int, height int, space string) {
	v.SelFgColor = gocui.ColorBlack | gocui.AttrBold
	v.SelBgColor = gocui.ColorWhite
	fmt.Fprintf(v, strings.Repeat(space, offset)+"\x1b[0;32m"+c.Name+"\n")
	fmt.Fprintf(v, "\x1b[0;39m")
	w.treeStructure = append(w.treeStructure, TreeItem{
		isChannel: true,
		c:         c,
	})
	height++
	users := c.Users
	var keys []int
	for k := range users {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	for _, val := range keys {
		u := c.Users[uint32(val)]
		fmt.Fprintf(v, strings.Repeat(space, offset+1)+u.Name)
		if u == (*w.client).Self {
			if w.speaking {
				fmt.Fprintf(v, " \x1b[0;32m Voice ")
			} else {
				fmt.Fprintf(v, " \x1b[0;31m Voice ")
			}

		}
		if u.SelfDeafened || u.Deafened {
			fmt.Fprintf(v, " \x1b[0;31mDeafened")
			fmt.Fprintf(v, " \x1b[0;31mMuted")
		} else if u.SelfMuted || u.Muted {
			fmt.Fprintf(v, " \x1b[0;31mM")
		}
		if u.Suppressed {
			fmt.Fprintf(v, " \x1b[0;31mSupressed")
		}
		fmt.Fprintf(v, "\x1b[0;39m\n")
		w.treeStructure = append(w.treeStructure, TreeItem{
			isUser: true,
			u:      c.Users[uint32(val)],
		})
		height++
	}
	keys = nil
	for k := range c.Children {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	offset++
	for _, val := range keys {
		w.DrawChannelTree(v, c.Children[uint32(val)], offset, height, space)
	}
}
func (w *LeftPane) setRightPaneDesc(g *gocui.Gui, v *gocui.View, y int) {
	_, cY := v.Cursor()
	if cY < 0 || cY >= len(w.treeStructure) {
		return
	}
	(*w.UIRightPane).contents = ""
	treeItem := w.treeStructure[cY]
	if treeItem.isChannel {
		if treeItem.c.Description == "" {
			treeItem.c.RequestDescription()
			//w.UIRightPane.contents = "Requesting channel description...\n"
		} else {
			(*w.UIRightPane).contents = treeItem.c.Description
		}
	} else if w.treeStructure[cY].isUser {
		if treeItem.u.Comment == "" {
			treeItem.u.RequestComment()
			//w.UIRightPane.contents = "Requesting user comment...\n"
		} else {
			(*w.UIRightPane).contents = treeItem.u.Comment
		}
	}
	g.Execute((*w.UIRightPane).Layout)
}
func (w *LeftPane) HandleUp(g *gocui.Gui, v *gocui.View) error {
	ox, oy := v.Origin()
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := v.SetOrigin(ox, oy-1); err != nil {
			return err
		}
	}
	w.setRightPaneDesc(g, v, cy)
	return nil
}
func (w *LeftPane) HandleDown(g *gocui.Gui, v *gocui.View) error {
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy+1); err != nil {
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+1); err != nil {
			return err
		}
	}
	w.setRightPaneDesc(g, v, cy)
	return nil
}
func (w *LeftPane) HandleAction(g *gocui.Gui, v *gocui.View) error {
	_, cY := v.Cursor()
	if cY < 0 || cY >= len(w.treeStructure) {
		return nil
	}
	if w.treeStructure[cY].isChannel {
		(*w.client).Self.Move(w.treeStructure[cY].c)
	}
	return nil
}
