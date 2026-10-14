package main

// Assert truth
// Replace the TODO value with true. Run the test to see how a failed check is reported.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	want := false // TODO: replace this value.
	got := true
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
