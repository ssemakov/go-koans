package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	double := func(n int) int { return n * 2 }
	want := -999 // TODO: replace this value.
	got := double(5)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
