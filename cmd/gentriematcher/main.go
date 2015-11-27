package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/Maki-Daisuke/go-argvreader"
	"github.com/Maki-Daisuke/go-gentriematcher"
	"github.com/Maki-Daisuke/go-lines"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	PkgName string `short:"P" long:"package" default:"main" description:"package name"`
	TagName string `short:"T" long:"tag" default:"" description:"tag name included in the generated functions"`
}

var reId = regexp.MustCompile(`^[0-9a-zA-Z_]+$`)

func main() {
	args, err := flags.Parse(&opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
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

	signatures := []string{}

	reader := argvreader.NewReader(args)
	line_chan, err_chan := lines.LinesWithError(reader)
	for line := range line_chan {
		if line == "" {
			continue
		}
		signatures = append(signatures, line)
	}
	err = <-err_chan
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("package %s\n\n", opts.PkgName)

	err = triematcher.GenerateMatcher(os.Stdout, opts.TagName, signatures)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
