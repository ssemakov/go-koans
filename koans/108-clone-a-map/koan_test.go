package main

import (
	"maps"
	"testing"
)

func TestKoan(t *testing.T) {
	original := map[string]int{"a": 1}
	clone := maps.Clone(original)
	clone["a"] = 9
	want := -999 // TODO: replace this value.
	got := original["a"]
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
