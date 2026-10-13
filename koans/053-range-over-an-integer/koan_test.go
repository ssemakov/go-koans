package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	total := 0
	for i := range 4 {
		total += i
	}
	want := -999 // TODO: replace this value.
	got := total
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
