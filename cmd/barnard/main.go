package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"

	"github.com/jroimartin/gocui"
	"github.com/layeh/barnard"
	_ "github.com/layeh/barnard/uiterm"
	"github.com/layeh/gumble/gumble"
	_ "github.com/layeh/gumble/opus"
)

func main() {
	// Command line flags
	server := flag.String("server", "localhost:64738", "the server to connect to")
	username := flag.String("username", "", "the username of the client")
	password := flag.String("password", "", "the password of the server")
	insecure := flag.Bool("insecure", false, "skip server certificate verification")
	certificate := flag.String("certificate", "", "PEM encoded certificate and private key")

	flag.Parse()

	// Initialize
	b := barnard.Barnard{
		Config:  gumble.NewConfig(),
		Address: *server,
	}

	b.Config.Username = *username
	b.Config.Password = *password

	if *insecure {
		b.TLSConfig.InsecureSkipVerify = true
	}
	if *certificate != "" {
		cert, err := tls.LoadX509KeyPair(*certificate, *certificate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		b.TLSConfig.Certificates = append(b.TLSConfig.Certificates, cert)
	}

	//TODO Create UI

	err := b.InitializeUI()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer b.UI.Close()

	b.UI.SetManager(&b.UILeftPane, &b.UIRightPane, &b.UIBottomPane, &b.UITextboxEntry)
	b.SetupHotkeys()
	if err := b.UI.MainLoop(); err != nil && err != gocui.ErrQuit {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

}
