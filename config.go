package main

import (
	"errors"
	"fmt"
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
	"github.com/jroimartin/gocui"
)

type Config struct {
	KeyNextView     string
	KeyAction       string
	KeyScrollDown   string
	KeyScrollUp     string
	KeyClearLog     string
	KeyQuit         string
	KeyVoice        string
	KeyDeafen       string
	DefaultServer   string
	DefaultUsername string
	DefaultPassword string
}

// LoadConfig will load the config located at .config/barnard/config.ini in the home directory of the user.
func (c *Config) LoadConfig() {
	usr, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	if _, err := toml.DecodeFile(usr.HomeDir+"/.config/barnard/config.ini", &c); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

// This function will return a uint16 key value from a string. Used to parse configuration.
func (c *Config) getKeyFromString(k string) (gocui.Key, error) {
	var ret gocui.Key
	switch k {
	case "KeyF1":
		ret = gocui.KeyF1
	case "KeyF2":
		ret = gocui.KeyF2
	case "KeyF3":
		ret = gocui.KeyF3
	case "KeyF4":
		ret = gocui.KeyF4
	case "KeyF5":
		ret = gocui.KeyF5
	case "KeyF6":
		ret = gocui.KeyF6
	case "KeyF7":
		ret = gocui.KeyF7
	case "KeyF8":
		ret = gocui.KeyF8
	case "KeyF9":
		ret = gocui.KeyF9
	case "KeyF10":
		ret = gocui.KeyF10
	case "KeyF11":
		ret = gocui.KeyF11
	case "KeyF12":
		ret = gocui.KeyF12
	case "KeyInsert":
		ret = gocui.KeyInsert
	case "KeyDelete":
		ret = gocui.KeyDelete
	case "KeyHome":
		ret = gocui.KeyHome
	case "KeyEnd":
		ret = gocui.KeyEnd
	case "KeyPgup":
		ret = gocui.KeyPgup
	case "KeyPgdn":
		ret = gocui.KeyPgdn
	case "KeyArrowUp":
		ret = gocui.KeyArrowUp
	case "KeyArrowDown":
		ret = gocui.KeyArrowDown
	case "KeyArrowLeft":
		ret = gocui.KeyArrowLeft
	case "KeyArrowRight":
		ret = gocui.KeyArrowRight
	case "MouseLeft":
		ret = gocui.MouseLeft
	case "MouseRight":
		ret = gocui.MouseRight
	case "MouseMiddle":
		ret = gocui.MouseMiddle
	case "KeyCtrlTidle":
		ret = gocui.KeyCtrlTilde
	case "KeyCtrl2":
		ret = gocui.KeyCtrl2
	case "KeyCtrlSpace":
		ret = gocui.KeyCtrlSpace
	case "KeyCtrlA":
		ret = gocui.KeyCtrlA
	case "KeyCtrlB":
		ret = gocui.KeyCtrlB
	case "KeyCtrlC":
		ret = gocui.KeyCtrlC
	case "KeyCtrlD":
		ret = gocui.KeyCtrlD
	case "KeyCtrlE":
		ret = gocui.KeyCtrlE
	case "KeyCtrlF":
		ret = gocui.KeyCtrlF
	case "KeyCtrlG":
		ret = gocui.KeyCtrlG
	case "KeyCtrlH":
		ret = gocui.KeyCtrlH
	case "KeyTab":
		ret = gocui.KeyTab
	case "KeyCtrlI":
		ret = gocui.KeyCtrlI
	case "KeyCtrlJ":
		ret = gocui.KeyCtrlJ
	case "KeyCtrlK":
		ret = gocui.KeyCtrlK
	case "KeyCtrlL":
		ret = gocui.KeyCtrlL
	case "KeyEnter":
		ret = gocui.KeyEnter
	case "KeyCtrlM":
		ret = gocui.KeyCtrlM
	case "KeyCtrlN":
		ret = gocui.KeyCtrlN
	case "KeyCtrlO":
		ret = gocui.KeyCtrlO
	case "KeyCtrlP":
		ret = gocui.KeyCtrlP
	case "KeyCtrlQ":
		ret = gocui.KeyCtrlQ
	case "KeyCtrlR":
		ret = gocui.KeyCtrlR
	case "KeyCtrlS":
		ret = gocui.KeyCtrlS
	case "KeyCtrlT":
		ret = gocui.KeyCtrlT
	case "KeyCtrlU":
		ret = gocui.KeyCtrlU
	case "KeyCtrlV":
		ret = gocui.KeyCtrlV
	case "KeyCtrlW":
		ret = gocui.KeyCtrlW
	case "KeyCtrlX":
		ret = gocui.KeyCtrlX
	case "KeyCtrlY":
		ret = gocui.KeyCtrlY
	case "KeyCtrlZ":
		ret = gocui.KeyCtrlZ
	case "KeyCtrlLsqBracket":
		ret = gocui.KeyCtrlLsqBracket
	case "KeyEsc":
		ret = gocui.KeyEsc
	case "KeyCtrl3":
		ret = gocui.KeyCtrl3
	case "KeyCtrl4":
		ret = gocui.KeyCtrl4
	case "KeyCtrl5":
		ret = gocui.KeyCtrl5
	case "KeyCtrl6":
		ret = gocui.KeyCtrl6
	case "KeyCtrl7":
		ret = gocui.KeyCtrl7
	case "KeyCtrl8":
		ret = gocui.KeyCtrl8
	case "KeyCtrlBackslash":
		ret = gocui.KeyCtrlBackslash
	case "KeyCtrlRsqBracket":
		ret = gocui.KeyCtrlRsqBracket
	case "KeyCtrlSlash":
		ret = gocui.KeyCtrlSlash
	case "KeyCtrlUnderscore":
		ret = gocui.KeyCtrlUnderscore
	case "KeySpace":
		ret = gocui.KeySpace
	case "KeyBackspace2":
		ret = gocui.KeyBackspace2
	default:
		return 0, errors.New("Unknown key")
	}
	return ret, nil
}
