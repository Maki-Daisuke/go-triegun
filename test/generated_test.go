package main

import "testing"

func TestGenerate(t *testing.T) {
	for _, it := range []string{
		"AppleCoreMedia", // iOS native quick time player
		"Apple-PubSub",   // RSS reader for screen saver?
		"cloudd/",        // iCloud daemon
		"itunesstored",
		"gamed/", // Probably, Apple Game Center
		"com.apple.geod",
		"com.apple.invitation-registration",
		"com.apple.Maps", // Map app?
		"ocspd/",         // Mac OS X's ocspd, verifying certificate validity
	} {
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
