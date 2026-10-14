package main

// Map assignment shares data
// Fill the map contents and lengths after modifying an assigned map. Assignment shares the
// entries; it does not clone the map.

import (
	"maps"
	"testing"
)

func TestKoan(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2}
	alias := original
	alias["a"] = 9
	delete(alias, "b")
	original["c"] = 3
	wantOriginal := map[string]int{} // TODO: all entries.
	wantAlias := map[string]int{}    // TODO: all entries.
	if !maps.Equal(original, wantOriginal) {
		t.Errorf("original=%v; want %v", original, wantOriginal)
	}
	if !maps.Equal(alias, wantAlias) {
		t.Errorf("alias=%v; want %v", alias, wantAlias)
	}
}
