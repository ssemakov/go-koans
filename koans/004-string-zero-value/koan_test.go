package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	var value string
	want := "__" // TODO: replace this value.
	got := value
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
