package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	values := map[string]int{"a": 2}
	delete(values, "a")
	want := -999 // TODO: replace this value.
	got := len(values)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
