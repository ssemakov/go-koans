package main

// Create a map
// Implement frequencies. Count every word, including an empty string, and return an
// initialized map for empty input.

import (
	"maps"
	"testing"
)

func frequencies(words []string) map[string]int {
	return nil // TODO: initialize and populate the map.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		words []string
		want  map[string]int
	}{
		{nil, map[string]int{}}, {[]string{"go"}, map[string]int{"go": 1}},
		{[]string{"go", "code", "go"}, map[string]int{"go": 2, "code": 1}},
		{[]string{"", "go", ""}, map[string]int{"": 2, "go": 1}},
	}
	for _, tc := range cases {
		got := frequencies(tc.words)
		if got == nil {
			t.Error("return an initialized map")
			continue
		}
		if !maps.Equal(got, tc.want) {
			t.Errorf("frequencies(%q)=%v; want %v", tc.words, got, tc.want)
		}
	}
}
