package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/Maki-Daisuke/go-argvreader"
	"github.com/Maki-Daisuke/go-lines"
	"github.com/Maki-Daisuke/go-triegun"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	PkgName          string `short:"p" long:"package" default:"main" description:"package name"`
	TagName          string `short:"t" long:"tag" default:"" description:"tag name included in the generated functions"`
	DisableIsIn      bool   `short:"I" long:"disable-isin" default:"false" description:"Suppress generating code for IsIn* functions"`
	DisableMatch     bool   `short:"M" long:"disable-match" default:"false" description:"Suppress generating code for Match* functions"`
	DisableHasPrefix bool   `short:"P" long:"disable-hasprefix" default:"false" description:"Suppress generating code for HasPrefix* functions"`
}

var reId = regexp.MustCompile(`^[0-9a-zA-Z_]+$`)

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Usage = "[OPTIONS] [FILES...]"
	args, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}
	if !reId.MatchString(opts.PkgName) {
		fmt.Fprintf(os.Stderr, "Package name must be an identifier, but %q is not\n", opts.PkgName)
		os.Exit(1)
	}
	if opts.TagName != "" && !reId.MatchString(opts.TagName) {
		fmt.Fprintf(os.Stderr, "Tag name must be an identifier, but %q is not\n", opts.TagName)
		os.Exit(1)
	}

	t := triegun.New()
	t.PkgName = opts.PkgName
	t.TagName = opts.TagName
	t.DisableIsIn = opts.DisableIsIn
	t.DisableMatch = opts.DisableMatch
	t.DisableHasPrefix = opts.DisableHasPrefix

	reader := argvreader.NewReader(args)
	line_chan, err_chan := lines.LinesWithError(reader)
	for line := range line_chan {
		if line == "" {
			continue
		}
		t.AddString(line)
	}
	err = <-err_chan
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = t.Gen(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
