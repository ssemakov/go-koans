package main

import (
	"testing"
)

func add(a, b int) int { return a + b }
func TestKoan(t *testing.T) {

	want := -999 // TODO: replace this value.
	got := add(3, 4)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
