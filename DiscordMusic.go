// This file provides a basic "quick start" example of using the Discordgo
// package to connect to Discord using the New() helper function.
package main

import (
	"fmt"
  "io/ioutil"
	"time"
  "os/exec"

	"github.com/bwmarrin/discordgo"
  "github.com/go-ini/ini"
)

func main() {
	ver := "1"

  fmt.Println("  ___  _                   _ __  __         _    ")
  fmt.Println(" |   \\(_)___ __ ___ _ _ __| |  \\/  |_  _ __(_)__ ")
  fmt.Println(" | |) | (_-</ _/ _ \\ '_/ _` | |\\/| | || (_-< / _|")
  fmt.Println(" |___/|_/__/\\__\\___/_| \\__,_|_|  |_|\\_,_/__/_\\__|")
  fmt.Println("Version " + ver)
	fmt.Println("")
  fmt.Println("___MUSIC___")

  cmd := exec.Command("snip/Snip.exe")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

  cfg, err := ini.Load("config.ini")
  if err != nil {
    fmt.Println(err)
    return
  }

  email := cfg.Section("Credentials").Key("email").String()
  password := cfg.Section("Credentials").Key("password").String()

	// Call the helper function New() passing username and password command
	// line arguments. This returns a new Discord session, authenticates,
	// connects to the Discord data websocket, and listens for events.
	dg, err := discordgo.New(email, password)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Open the websocket and begin listening.
	dg.Open()

  file, err := ioutil.ReadFile("snip/Snip.txt")
  if err != nil {
    panic(err)
  }

	if string(file) != "" {
		dg.UpdateStatus(0, "🎧 " + string(file))
		fmt.Println("Now Playing: " + string(file))
	} else {
		dg.UpdateStatus(0, "🎧 Nothing")
		fmt.Println("Now Playing: Nothing")
	}

  for {
    new, err := ioutil.ReadFile("snip/Snip.txt")
    if err != nil {
      panic(err)
    }
    if string(file) != string(new) {
			if string(new) != "" {
	      dg.UpdateStatus(0, "🎧 " + string(new))
	      fmt.Println("Now Playing: " + string(new))
			} else {
				dg.UpdateStatus(0, "🎧 Nothing")
	      fmt.Println("Now Playing: Nothing")
			}
    }
    file = new
    time.Sleep(5000000000)
  }
}
