package main

// Match a pattern
// Write the regexp for a complete ID: lowercase go followed by one or more ASCII digits.
// Reject prefixes, suffixes, spaces, and non-ASCII digits.

import (
	"regexp"
	"testing"
)

var idPattern = regexp.MustCompile("TODO") // TODO: supply an anchored pattern.

func TestKoan(t *testing.T) {
	cases := []struct {
		text string
		want bool
	}{
		{"go0", true}, {"go123", true}, {"go", false}, {"Go1", false}, {"xgo1", false},
		{"go1x", false}, {"go-1", false}, {"go1\n", false}, {"go١", false}, {"", false},
	}
	for _, tc := range cases {
		if got := idPattern.MatchString(tc.text); got != tc.want {
			t.Errorf("MatchString(%q)=%v; want %v", tc.text, got, tc.want)
		}
	}
}
