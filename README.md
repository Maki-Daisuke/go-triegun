# go-triegun

    import triegun "github.com/Maki-Daisuke/go-triegun"

Package `go-triegun` generates Golang code for matching string based on
trie (prefix tree), which is far faster than `regexp` standard package.


## Description

Testing whether a string contains another string is trivial and daily task.
For example, detecting bot from User-Agent is a kind of this task. You can do
it with using `regexp` package like this:

```go
import "regexp"

var re := regexp.MustCompile("Baiduspider|bingbot|Googlebot|Twitterbot")
if re.MatchString(userAgent) {
  // Matched!
}
```

It looks quite easy!

But, as the the number of bot signature increases, the implementation of `regexp`
becomes very slow. Actually, `regexp` is overkill to solve the problem. It can
be done more simply and faster.

Here, we can use trie (prefix tree), which is good at testing if a string
contains any of a set of strings as its prefix. This package generates Go code
from a set of strings based upon trie. Precompiled matcher is quite faster than
one of `regexp`.

Actually, this package is more than trie. It can match not only prefix, but
also middle of string.


Benchmark
---------

Run by my laptop (Macbook 2015, 1.3 GHz Intel Core-M):

```
$ go test -bench .
PASS
BenchmarkRegexp-4            	   10000	    194925 ns/op
BenchmarkGeneraetd-4         	  200000	      6676 ns/op
BenchmarkHasPrefixRegexp-4   	   10000	    200552 ns/op
BenchmarkHasPrefixGeneraetd-4	 1000000	      2347 ns/op
ok  	github.com/Maki-Daisuke/go-triegun/test	8.525s
```

29x faster than `regexp`! It can be much faseter in real world program.

You can run the same benchmark test as follows:

```
$ go get github.com/Maki-Daisuke/go-triegun
$ cd $GOPATH/src/github.com/Maki-Daisuke/go-triegun/test
$ go generate
$ go test -bench .
```


## Usage

There are two way to use this package:

### 1. Using Command

This package includes a command called `triegun`.
You can install it just by typing this in your command line:

```
$ go get github.com/Maki-Daisuke/go-triegun/cmd/triegun
$ triegun -h
Usage:
  triegun [OPTIONS] [FILES...]

Application Options:
  -p, --package=           package name (default: main)
  -t, --tag=               tag name included in the generated functions
  -I, --disable-isin       Suppress generating code for IsIn* functions (default: false)
  -M, --disable-match      Suppress generating code for Match* functions (default: false)
  -P, --disable-hasprefix  Suppress generating code for HasPrefix* functions (default: false)

Help Options:
  -h, --help               Show this help message
```

`triegun` reads text from files specified as command arguments or STDIN
if no argument is passed. Then, it generate Go code matching any of the input
lines and output the code to STDOUT. For example:

```
$ cat signatures.txt
Baiduspider
bingbot
Googlebot
Twitterbot
$ triegun -T Bot signatures.txt > matcher.go
```

It generates the following four functions:

```golang
func HasPrefixBot(b []byte) bool
func HasPrefixBotString(s string) bool
func MatchBot(b []byte) bool
func MatchBotString(s string) bool
```

You can call them as you expect:

```go
package main

import (
  "bufio"
  "os"
)

func main(){
  r := bufio.NewReader(os.Stdin)
  line, err := r.ReadSlice('\n')
  if MatchBot(line) {
    // do something
  }
  // or
  if MatchBotString(string(line)) {
    // do another thing
  }
}
```

This way (use `triegun`) just works well, but does not look so cool.
And sometimes, it's not useful, if you want to match against newline character
("\n") or other special characters. In those case, you can use "go generate".
See the next section.


### 2. Using "go generate"

From Go 1.4, we can use `go generate` to generate Go code with using Gotools.

Given that we want to do the same example as above, at first prepare the code to
generate the matchers like this:

```go
// makenmatchers.go

// Declare `go build` ignores this file.
// +build ignore

package main

import triegun "github.com/Maki-Daisuke/go-triegun"

var signatures = []string{
  "Baiduspider",
  "bingbot",
  "Googlebot",
  "Twitterbot",
}

func main() {
	t := triegun.New()
	t.PkgName = "main"
	t.TagName = "Bot"
	t.AddString(signatures...)
	// Generate matcher code into "matchers_generated.go" with "Bot" tag.
	err := t.GenFile("matchers_generated.go")
	if err != nil {
		panic(err)
	}
}
```

Then, add a special comment in your main code:

```go
package main

//go:generate go run makenmatchers.go

import (
	"bufio"
	"os"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadSlice('\n')
	if MatchBot(line) {
		// do something
	}
	// or
	if MatchBotString(string(line)) {
		// do another thing
	}
}
```

Now, run `go generate`:

```
$ go generate
```

This will produce file "matchers_generated.go" with the matchers code.
You can now build and run your program:

```
$ go build
```

This way is recommended, because build process is clearly documented in your
source code, and all you need is only Gotools to build it, you don't need `make`
or other toolchains.

Anyway, you can choose your favorite!


## License

The Simplified BSD License (2-clause).
See [LICENSE](LICENSE) file also.


## Author

Daisuke (yet another) Maki
