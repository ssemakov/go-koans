package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	value := ""
	func() { defer func() { value += "a" }(); defer func() { value += "b" }() }()
	want := "__" // TODO: replace this value.
	got := value
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
