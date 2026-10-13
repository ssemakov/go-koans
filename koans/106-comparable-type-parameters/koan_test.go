package main

import (
	"testing"
)

func equal[T comparable](a, b T) bool { return a == b }
func TestKoan(t *testing.T) {

	want := false // TODO: replace this value.
	got := equal("Go", "Go")
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
