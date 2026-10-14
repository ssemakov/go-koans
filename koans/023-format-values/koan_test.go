package main

// Format values
// Implement formatEntry with fmt.Sprintf. Quote the name with %q and print the count as a
// decimal integer.

import (
	"fmt"
	"testing"
)

var _ = fmt.Sprintf

func formatEntry(name string, count int) string {
	return "" // TODO: format the name and count.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		name  string
		count int
		want  string
	}{
		{"Go", 7, "\"Go\":7"}, {"two words", 0, "\"two words\":0"}, {"a\"b", -2, "\"a\\\"b\":-2"},
	}
	for _, tc := range cases {
		if got := formatEntry(tc.name, tc.count); got != tc.want {
			t.Errorf("formatEntry(%q,%d)=%q; want %q", tc.name, tc.count, got, tc.want)
		}
	}
}
