package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	const (
		first = iota
		second
		third
	)
	want := -999 // TODO: replace this value.
	got := third
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
