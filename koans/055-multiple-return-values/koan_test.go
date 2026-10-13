package main

import (
	"testing"
)

func pair() (int, string) { return 3, "go" }
func TestKoan(t *testing.T) {
	_, word := pair()
	want := "__" // TODO: replace this value.
	got := word
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
