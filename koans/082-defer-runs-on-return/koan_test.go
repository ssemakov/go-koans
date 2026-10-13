package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	value := 0
	func() { defer func() { value++ }() }()
	want := -999 // TODO: replace this value.
	got := value
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
