package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	_ "time"

	"github.com/jroimartin/gocui"
	"github.com/kennygrant/sanitize"
	"github.com/layeh/gumble/gumble"
)

// TODO Sanitation ..?

func esc(str string) string {
	return sanitize.HTML(str)
}

func (b *Barnard) InitializeUI() error {
	var err error
	b.UI, err = gocui.NewGui()
	if err != nil {
		return err
	}
	b.UI.Highlight = true
	b.UILeftPane = &LeftPane{
		Name:        "LeftPane",
		Title:       "Barnard",
		X:           0,
		Y:           0,
		W:           0.7,
		H:           0.6,
		client:      &b.Client,
		UIRightPane: &b.UIRightPane,
		speaking:    false,
	}
	b.UIRightPane = &RightPane{
		Name:       "RightPane",
		Title:      "Information",
		UILeftPane: b.UILeftPane,
	}
	b.UIBottomPane = &BottomPane{
		Name:       "BottomPane",
		Title:      "log",
		UILeftPane: b.UILeftPane,
		Autoscroll: true,
	}
	b.UITextboxEntry = &TextboxEntry{
		Name:         "TextboxEntry",
		Title:        "",
		UIBottomPane: b.UIBottomPane,
		SendMessage:  b.HandleSendMessage,
	}
	b.ViewArray = []string{"TextboxEntry", "BottomPane", "LeftPane", "RightPane"}
	b.ViewActive = -1
	b.UI.SetManager(b.UILeftPane, b.UIRightPane, b.UIBottomPane, b.UITextboxEntry)
	b.SetupHotkeys()
	b.UI.Cursor = true
	b.UI.SelFgColor = gocui.ColorYellow
	b.start()
	b.UILeftPane.channels = &b.Client.Channels
	return nil
}
func (b *Barnard) socketListener() {
	var err error
	b.Listener, err = net.Listen("unix", "/tmp/barnard.sock")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	go func() {
		for {
			if b.Listener == nil {
				return
			}
			fd, err := b.Listener.Accept()
			if err != nil {
				//fmt.Fprintf(os.Stderr, "%s\n", err)
			}
			go b.socketServer(fd)
		}
	}()
}

func (b *Barnard) socketServer(c net.Conn) {
	buf := make([]byte, 512)
	if c == nil {
		return
	}
	nr, err := c.Read(buf)

	if err != nil {
		return
	}

	if nr > 0 {
		b.ToggleVoice()
	}
}

func (b *Barnard) SetupHotkeys() {
	var quitKey gocui.Key
	var actionKey gocui.Key
	var scrollDownKey gocui.Key
	var scrollUpKey gocui.Key
	var clearLogKey gocui.Key
	var nextViewKey gocui.Key
	var voiceKey gocui.Key
	var deafenKey gocui.Key
	var err error

	if quitKey, err = b.BarnardConfig.getKeyFromString(b.BarnardConfig.KeyQuit); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		quitKey = gocui.KeyCtrlC
	}
	if actionKey, err = b.BarnardConfig.getKeyFromString(b.BarnardConfig.KeyAction); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		actionKey = gocui.KeyEnter
	}
	if scrollDownKey, err = b.BarnardConfig.getKeyFromString(b.BarnardConfig.KeyScrollDown); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		scrollDownKey = gocui.KeyArrowDown
	}
	if scrollUpKey, err = b.BarnardConfig.getKeyFromString(b.BarnardConfig.KeyScrollUp); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		scrollUpKey = gocui.KeyArrowUp
	}
	if clearLogKey, err = b.BarnardConfig.getKeyFromString(b.BarnardConfig.KeyClearLog); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		clearLogKey = gocui.KeyCtrlL
	}
	if nextViewKey, err = b.BarnardConfig.getKeyFromString(b.BarnardConfig.KeyNextView); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		nextViewKey = gocui.KeyTab
	}
	if voiceKey, err = b.BarnardConfig.getKeyFromString(b.BarnardConfig.KeyVoice); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		voiceKey = gocui.KeyF1
	}
	if deafenKey, err = b.BarnardConfig.getKeyFromString(b.BarnardConfig.KeyDeafen); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		voiceKey = gocui.KeyCtrlM
	}

	if err := b.UI.SetKeybinding("", quitKey, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return gocui.ErrQuit }); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	if err := b.UI.SetKeybinding("", nextViewKey, gocui.ModNone, b.handleChangeView); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	if err := b.UI.SetKeybinding("TextboxEntry", actionKey, gocui.ModNone, b.UITextboxEntry.HandleSendMessageKey); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	if err := b.UI.SetKeybinding("", deafenKey, gocui.ModNone, b.HandleDeafen); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	if err := b.UI.SetKeybinding("BottomPane", scrollDownKey, gocui.ModNone, b.UIBottomPane.ScrollDown); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	if err := b.UI.SetKeybinding("BottomPane", clearLogKey, gocui.ModNone, b.UIBottomPane.Clear); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	if err := b.UI.SetKeybinding("BottomPane", scrollUpKey, gocui.ModNone, b.UIBottomPane.ScrollUp); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	if err := b.UI.SetKeybinding("LeftPane", scrollUpKey, gocui.ModNone, b.UILeftPane.HandleUp); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	if err := b.UI.SetKeybinding("LeftPane", scrollDownKey, gocui.ModNone, b.UILeftPane.HandleDown); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	if err := b.UI.SetKeybinding("LeftPane", actionKey, gocui.ModNone, b.UILeftPane.HandleAction); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	if err := b.UI.SetKeybinding("", voiceKey, gocui.ModNone, b.HandleToggleVoice); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	b.socketListener()
}

func (b *Barnard) ToggleVoice() {
	if b.UILeftPane == nil {
		return
	}
	if b.UILeftPane.speaking {
		b.Stream.StopSource()
		b.UILeftPane.speaking = false
	} else {
		b.Stream.StartSource()
		b.UILeftPane.speaking = true
	}
	b.UI.Execute(b.UILeftPane.Layout)
}
func (b *Barnard) HandleToggleVoice(g *gocui.Gui, v *gocui.View) error {
	b.ToggleVoice()
	return nil
}
func (b *Barnard) HandleSendMessage(text string) {
	if text == "" {
		return
	}
	if b.Client != nil && b.Client.Self != nil {
		b.Client.Self.Channel.Send(text, false)
		b.AddTextMessage(b.Client.Self, text)
	}
}

func (b *Barnard) handleChangeView(g *gocui.Gui, v *gocui.View) error {
	b.ViewActive = (b.ViewActive + 1) % len(b.ViewArray)

	if _, err := b.UI.SetCurrentView(b.ViewArray[b.ViewActive]); err != nil {
		return err
	}
	//if _, err := b.UI.SetViewOnTop(b.ViewArray[b.ViewActive]); err != nil {
	//	return err
	//}
	return nil
}

func (b *Barnard) UpdateInputStatus(status string) {

}
func (b *Barnard) AddTextMessage(sender *gumble.User, message string) {
	if sender == nil {
		b.UIBottomPane.AddLine(b.UI, message)
	} else {
		b.UIBottomPane.AddLine(b.UI, fmt.Sprintf("%s: %s\n", sender.Name, strings.TrimSpace(esc(message))))
	}
}

func (b *Barnard) HandleDeafen(g *gocui.Gui, v *gocui.View) error {
	d := b.Client.Self.SelfDeafened
	b.Client.Self.SetSelfDeafened(!d)
	if d {
		b.Client.Self.SetSelfMuted(!d)
	}
	g.Execute(b.UILeftPane.Layout)
	return nil
}
