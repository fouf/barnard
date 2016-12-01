package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/layeh/gumble/gumble"
	_ "github.com/layeh/gumble/opus"
)

func main() {
	// Configuration
	var c Config
	c.LoadConfig()
	// Command line flags
	time.Sleep(time.Second * 2)
	server := flag.String("server", c.DefaultServer, "the server to connect to")
	username := flag.String("username", c.DefaultUsername, "the username of the client")
	password := flag.String("password", c.DefaultPassword, "the password of the server")
	insecure := flag.Bool("insecure", false, "skip server certificate verification")
	certificate := flag.String("certificate", "", "PEM encoded certificate and private key")

	flag.Parse()

	// Initialize
	b := Barnard{
		Config:        gumble.NewConfig(),
		Address:       *server,
		BarnardConfig: c,
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
	defer b.Listener.Close()
	if err := b.UI.MainLoop(); err != nil && err != gocui.ErrQuit {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
