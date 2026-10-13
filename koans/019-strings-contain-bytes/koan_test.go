package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	value := "é"
	want := -999 // TODO: replace this value.
	got := len(value)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
