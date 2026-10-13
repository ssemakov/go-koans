package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	value := 3.75
	want := -999 // TODO: replace this value.
	got := int(value)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
