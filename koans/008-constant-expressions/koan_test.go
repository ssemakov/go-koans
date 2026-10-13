package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	const value = 1 << 3
	want := -999 // TODO: replace this value.
	got := value
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
