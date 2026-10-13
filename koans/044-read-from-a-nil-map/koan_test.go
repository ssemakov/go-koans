package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	var values map[string]int
	want := -999 // TODO: replace this value.
	got := values["x"]
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
