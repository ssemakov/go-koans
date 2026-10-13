package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	type Count int
	value := Count(4)
	want := -999 // TODO: replace this value.
	got := int(value)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
