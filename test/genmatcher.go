package main

import (
	"fmt"
	"os"

	gentriematcher "github.com/Maki-Daisuke/go-gentriematcher"
)

var signatures = []string{
	"AppleCoreMedia", // iOS native quick time player
	"Apple-PubSub",   // RSS reader for screen saver?
	"cloudd/",        // iCloud daemon
	"itunesstored",
	"gamed/", // Probably, Apple Game Center
	"com.apple.geod",
	"com.apple.invitation-registration",
	"com.apple.Maps", // Map app?
	"ocspd/",         // Mac OS X's ocspd, verifying certificate validity
}

func main() {
	out, err := os.OpenFile("ua_gentriematcher.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintln(out, "package main")
	err = gentriematcher.GenerateMatcher(out, "UA", signatures)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
