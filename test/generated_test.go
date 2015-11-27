package main

// Run "go generate" first to run test.
//go:generate go run genmatcher.go

import (
	"regexp"
	"strings"
	"testing"
)

var userAgents = []string{
	"Mozilla/5.0 (Linux; U; Android 4.0.4; ja-jp; SonyEricssonSO-03D Build/6.1.F.0.106) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
	"Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; rv:11.0) like Gecko/20100101 Firefox/12.0",
	"Mozilla/5.0 (Linux; U; Android 4.2.2; ja-jp; SonySO-02F Build/14.1.H.2.119) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 8_4_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12H321 Safari/600.1.4",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
	"iPad3,4/7.0.2 (11A501)",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 8_4_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12H321",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 8_4_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12H321 Safari/600.1.4",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_0_2 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13A452 Safari/601.1",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; InfoPath.2)",
	"Mozilla/5.0 (Linux; U; Android 4.4.2; ja-jp; LGV31 Build/KVT49L.LGV3110f) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/34.0.1847.118 Mobile Safari/537.36",
	"iPad3,4/7.0.2 (11A501)",
	"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.93 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2528.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10240",
}

var answers = []bool{
	false,
	false,
	false,
	true,
	false,
	true,
	true,
	true,
	true,
	false,
	false,
	true,
	false,
	false,
	false,
	false,
}

func TestGenerate(t *testing.T) {
	for i, it := range userAgents {
		if MatchUAString(it) != answers[i] {
			if answers[i] {
				t.Errorf(`should match against %q, but didn't`, it)
			} else {
				t.Errorf(`should not match against %q, but did`, it)
			}
		}
	}
}

func BenchmarkRegexp(b *testing.B) {
	for i := range signatures {
		signatures[i] = regexp.QuoteMeta(signatures[i])
	}
	reApple := regexp.MustCompile(strings.Join(signatures, "|"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ua := range userAgents {
			reApple.MatchString(ua)
		}
	}
}

func BenchmarkGeneraetd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, ua := range userAgents {
			MatchUAString(ua)
		}
	}
}
