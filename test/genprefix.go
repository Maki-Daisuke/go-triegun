// +build ignore

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
	out, err := os.OpenFile("ua_genprefix.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(out, "package main")
	err = gentriematcher.GenerateHasPrefix(out, "UA", signatures)
	if err != nil {
		panic(err)
	}
}
