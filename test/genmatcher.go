package main

import (
	"fmt"
	"os"

	gentriematcher "github.com/Maki-Daisuke/go-gentriematcher"
)

var signatures = []string{
	"CFNetwork/",
	"iOS",
	"iPhone OS",
	"iPhone;",
	"iPad3,",
	"Mac OS X",
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
