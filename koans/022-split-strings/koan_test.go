package main

// Split strings
// Implement splitFields with strings.Split. The separator is an exact string, not a set of
// delimiter characters. Preserve empty fields.

import (
	"slices"
	"strings"
	"testing"
)

var _ = strings.Split

func splitFields(text, sep string) []string {
	return nil // TODO: use strings.Split.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		text, sep string
		want      []string
	}{
		{"a,b,c", ",", []string{"a", "b", "c"}},
		{"a,,b,", ",", []string{"a", "", "b", ""}},
		{"a,|;b;c", ",|;", []string{"a", "b;c"}},
		{"", ",", []string{""}},
		{"abc", ",", []string{"abc"}},
	}
	for _, tc := range cases {
		if got := splitFields(tc.text, tc.sep); !slices.Equal(got, tc.want) {
			t.Errorf("splitFields(%q,%q)=%q; want %q", tc.text, tc.sep, got, tc.want)
		}
	}
}
