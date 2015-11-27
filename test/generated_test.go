package main

// Run "go generate" first to run test.
//go:generate go run genmatcher.go

import "testing"

func TestGenerate(t *testing.T) {
	for _, it := range signatures {
		if !MatchUAString(it) {
			t.Errorf(`should match against %q, but didn't`, it)
		}
	}

	if MatchUAString(`hogeFuga`) {
		t.Error(`should not match against "hogeFuga", but did`)
	}
	if !MatchUAString(`com.apple.geodd`) {
		t.Error(`should match against "com.apple.geodd", but didn't`)
	}
	if MatchUAString(`Apple`) {
		t.Error(`should not match against "Apple", but did`)
	}
	if !MatchUAString("Apple iPhone OS v3.1.3 AppleCoreMedia v1.0.0.7E18") {
		t.Error(`should match against "Apple iPhone OS v3.1.3 AppleCoreMedia v1.0.0.7E18", but didn't`)
	}
}
