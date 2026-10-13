package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	a := map[string]int{"x": 1}
	b := a
	b["x"] = 5
	want := -999 // TODO: replace this value.
	got := a["x"]
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
