package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	a, b := 2, 5
	a, b = b, a
	want := -999 // TODO: replace this value.
	got := a*10 + b
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
