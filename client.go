package main

import (
	"fmt"
	"net"
	"os"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleopenal"
	"github.com/layeh/gumble/gumbleutil"
)

func (b *Barnard) start() {
	b.Config.Attach(gumbleutil.AutoBitrate)
	b.Config.Attach(b)

	var err error
	_, err = gumble.DialWithDialer(new(net.Dialer), b.Address, b.Config, &b.TLSConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Audio
	if os.Getenv("ALSOFT_LOGLEVEL") == "" {
		os.Setenv("ALSOFT_LOGLEVEL", "0")
	}
	if stream, err := gumbleopenal.New(b.Client); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	} else {
		b.Stream = stream
	}
}

// OnConnect logs in the chat that we have connected, and the welcome message.
func (b *Barnard) OnConnect(e *gumble.ConnectEvent) {

	b.Client = e.Client

	b.UITextboxEntry.UpdateInputStatus(b.UI, fmt.Sprintf("To: %s", e.Client.Self.Channel.Name))
	b.UIBottomPane.AddLine(b.UI, fmt.Sprintf("Connected to %s\n", b.Client.Conn.RemoteAddr()))
	if e.WelcomeMessage != nil {
		b.UIBottomPane.AddLine(b.UI, (fmt.Sprintf("Welcome message: %s", esc(*e.WelcomeMessage))))
	}
}

// OnDisconnect disconnects gumble and prints the reason in the chat log.
func (b *Barnard) OnDisconnect(e *gumble.DisconnectEvent) {
	var reason string
	switch e.Type {
	case gumble.DisconnectError:
		reason = "connection error"
	}
	if reason == "" {
		b.UIBottomPane.AddLine(b.UI, "Disconnected\n")
	} else {
		b.UIBottomPane.AddLine(b.UI, "Disconnected: "+reason+"\n")
	}
}

// OnTextMessage  calls AddTextMessage to display the chat message.
func (b *Barnard) OnTextMessage(e *gumble.TextMessageEvent) {
	b.AddTextMessage(e.Sender, e.Message)
}

// OnUserChange will handle when a user changes channels, including ourselves.
func (b *Barnard) OnUserChange(e *gumble.UserChangeEvent) {
	if e.Type.Has(gumble.UserChangeChannel) && e.User == b.Client.Self {
		b.UITextboxEntry.UpdateInputStatus(b.UI, fmt.Sprintf("To: %s", e.User.Channel.Name))
	}
	b.UI.Execute(b.UILeftPane.Layout)
}

func (b *Barnard) OnChannelChange(e *gumble.ChannelChangeEvent) {
	b.UI.Execute(b.UILeftPane.Layout)
}

// OnPermissionDenied handles a permission event, printing the
// error to the chat log.
func (b *Barnard) OnPermissionDenied(e *gumble.PermissionDeniedEvent) {
	var info string
	switch e.Type {
	case gumble.PermissionDeniedOther:
		info = e.String
	case gumble.PermissionDeniedPermission:
		info = "insufficient permissions"
	case gumble.PermissionDeniedSuperUser:
		info = "cannot modify SuperUser"
	case gumble.PermissionDeniedInvalidChannelName:
		info = "invalid channel name"
	case gumble.PermissionDeniedTextTooLong:
		info = "text too long"
	case gumble.PermissionDeniedTemporaryChannel:
		info = "temporary channel"
	case gumble.PermissionDeniedMissingCertificate:
		info = "missing certificate"
	case gumble.PermissionDeniedInvalidUserName:
		info = "invalid user name"
	case gumble.PermissionDeniedChannelFull:
		info = "channel full"
	case gumble.PermissionDeniedNestingLimit:
		info = "nesting limit"
	}
	b.UIBottomPane.AddLine(b.UI, fmt.Sprintf("Permission denied: %s\n", info))
}

func (b *Barnard) OnUserList(e *gumble.UserListEvent) {
}

func (b *Barnard) OnACL(e *gumble.ACLEvent) {
}

func (b *Barnard) OnBanList(e *gumble.BanListEvent) {
}

func (b *Barnard) OnContextActionChange(e *gumble.ContextActionChangeEvent) {
}

func (b *Barnard) OnServerConfig(e *gumble.ServerConfigEvent) {
}
