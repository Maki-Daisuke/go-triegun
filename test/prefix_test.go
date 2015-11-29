package main

// Run "go generate" first to run test.
//go:generate go run genprefix.go

import (
	"regexp"
	"strings"
	"testing"
)

var rePrefix *regexp.Regexp

func init() {
	s := make([]string, len(signatures))
	for i := range signatures {
		s[i] = regexp.QuoteMeta(signatures[i])
	}
	rePrefix = regexp.MustCompile("^(?:" + strings.Join(s, "|") + ")")
}

func TestHasPrefix(t *testing.T) {
	for _, it := range userAgents {
		if HasPrefixUAString(it) != rePrefix.MatchString(it) {
			if rePrefix.MatchString(it) {
				t.Errorf(`should match against %q, but didn't`, it)
			} else {
				t.Errorf(`should not match against %q, but did`, it)
			}
		}
	}
}

func BenchmarkHasPrefixRegexp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ua := range userAgents {
			reApple.MatchString(ua)
		}
	}
}

func BenchmarkHasPrefixGeneraetd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, ua := range userAgents {
			HasPrefixUAString(ua)
		}
	}
}
