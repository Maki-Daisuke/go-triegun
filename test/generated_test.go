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
		if !MatchUA(it) {
			t.Errorf(`should match against %q, but didn't`, it)
		}
	}

	if MatchUA(`hogeFuga`) {
		t.Error(`should not match against "hogeFuga", but did`)
	}
	if MatchUA(`com.apple.geodd`) {
		t.Error(`should not match against "com.apple.geodd", but did`)
	}
	if MatchUA(`Apple`) {
		t.Error(`should not match against "Apple", but did`)
	}
}
