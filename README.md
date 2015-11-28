# go-gentriematcher

    import "github.com/Maki-Daisuke/go-gentriematcher"

Package `go-gentriematcher` generates Golang code for matching string based on
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
BenchmarkRegexp-4   	   10000	    219568 ns/op
BenchmarkGeneraetd-4	  200000	      7679 ns/op
ok  	github.com/Maki-Daisuke/go-gentriematcher/test	3.842s
```

28x faster than `regexp`! It can be much faseter in real world program.

You can run the same benchmark test as follows:

```
$ go get github.com/Maki-Daisuke/go-gentriematcher
$ cd $GOPATH/src/github.com/Maki-Daisuke/go-gentriematcher/test
$ go generate
$ go test -bench .
```


## Usage

There are two way to use this package:

### 1. Using Command

This package includes a command called `gentriematcher`.
You can install it just by typing this in your command line:

```
$ go get github.com/Maki-Daisuke/go-gentriematcher/cmd/gentriematcher
$ gentriematcher -h
Usage:
  gentriematcher [OPTIONS] [FILES]

Application Options:
  -P, --package= package name (default: main)
  -T, --tag=     tag name included in the generated functions

Help Options:
  -h, --help     Show this help message
```

`gentriematcher` reads text from files specified as command arguments or STDIN
if no argument is passed. Then, it generate Go code matching any of the input
lines and output the code to STDOUT. For example:

```
$ cat signatures.txt
Baiduspider
bingbot
Googlebot
Twitterbot
$ gentriematcher signatures.txt > matcher.go
```

It generates the following two functions:

```golang
func Match<TAG>(b []byte) bool
func Match<TAG>String(s string) bool
```

Here, `<TAG>` is a string specified by `-P` option (see below).
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
  if Match(line) {
    // do something
  }
  // or
  if MatchString(string(line)) {
    // do another thing
  }
}
```

#### Options

- `-P` | `--package`
  - Package name used in the generated file.
  - Default: `"main"`
- `-T` | `--tag`
  - Tag that is icluded in the generated functions
  - For example, if you specify `"UA"`, it generates functions `MatchUA` and `MatchUAString`

This way (use `gentriematcher`) just works well, but does not look so cool.
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

import (
	"fmt"
	"os"

	gentriematcher "github.com/Maki-Daisuke/go-gentriematcher"
)

var signatures = []string{
  "Baiduspider",
  "bingbot",
  "Googlebot",
  "Twitterbot",
}

func main() {
	out, err := os.OpenFile("matchers_generated.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
    panic(err)
	}
	fmt.Fprintln(out, "package main")
  // Generate matcher code into "matchers_generated.go" with empty tag ("").
	err = gentriematcher.GenerateMatcher(out, "", signatures)
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
  if Match(line) {
    // do something
  }
  // or
  if MatchString(string(line)) {
    // do another thing
  }
}
```

Now, run "go generate":

```
$ go generate
```

This will produce file "matchers_generated.go" with the matchers code.
You can now build and run your program:

```
$ go run main.go
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
