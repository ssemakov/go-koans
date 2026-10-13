package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	values := map[string]int{"a": 2}
	want := -999 // TODO: replace this value.
	got := values["missing"]
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
