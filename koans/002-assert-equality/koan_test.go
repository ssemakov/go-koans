package main

import (
	"testing"
)

func TestKoan(t *testing.T) {

	want := -999 // TODO: replace this value.
	got := 1 + 1
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
