package main

// Delete a key
// Implement removeKeys. Delete the requested keys in place, preserve other entries, and
// accept absent keys or a nil map.

import (
	"maps"
	"testing"
)

func removeKeys(values map[string]int, keys []string) { // TODO: delete only the requested keys.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		values map[string]int
		keys   []string
		want   map[string]int
	}{
		{map[string]int{"a": 1, "b": 2}, []string{"a"}, map[string]int{"b": 2}},
		{map[string]int{"a": 1}, []string{"missing", "a", "a"}, map[string]int{}},
		{map[string]int{"a": 1}, nil, map[string]int{"a": 1}},
		{nil, []string{"a"}, nil},
	}
	for _, tc := range cases {
		removeKeys(tc.values, tc.keys)
		if !maps.Equal(tc.values, tc.want) {
			t.Errorf("removeKeys(%q): got %v; want %v", tc.keys, tc.values, tc.want)
		}
	}
}
