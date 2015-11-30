package main

import (
	"os"

	triegun "github.com/Maki-Daisuke/go-triegun"
)

const OUT_FILE = "ua_triegun.go"

var signatures = []string{
	"CFNetwork/",
	"iOS",
	"iPhone OS",
	"iPhone;",
	"iPad3,",
	"Mac OS X",
}

func main() {
	out, err := os.OpenFile(OUT_FILE, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	t := triegun.New()
	t.TagName = "UA"
	t.AddString(signatures...)
	err = t.Gen(out)
	out.Close()
	if err != nil {
		os.Remove(OUT_FILE)
		panic(err)
	}
}
