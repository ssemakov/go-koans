package main

import (
	"testing"
)

func increment(n *int) { (*n)++ }
func TestKoan(t *testing.T) {
	value := 3
	increment(&value)
	want := -999 // TODO: replace this value.
	got := value
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
