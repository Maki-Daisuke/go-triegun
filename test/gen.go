package main

import triegun "github.com/Maki-Daisuke/go-triegun"

var signatures = []string{
	"CFNetwork/",
	"iOS",
	"iPhone OS",
	"iPhone;",
	"iPad3,",
	"Mac OS X",
}

func main() {
	t := triegun.New()
	t.TagName = "UA"
	t.AddString(signatures...)
	err := t.GenFile("ua_triegun.go")
	if err != nil {
		panic(err)
	}
}
