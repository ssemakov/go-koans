package main

// Array assignment copies values
// Fill both expected arrays after the copy is modified. Check the whole values, not just
// one element.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	original := [3]int{1, 2, 3}
	copy := original
	copy[0] = 9

	wantOriginal := [3]int{-1, -1, -1} // TODO: replace this value.
	if got := original; got != wantOriginal {
		t.Errorf("Original: got %#v; want %#v", got, wantOriginal)
	}

	wantCopy := [3]int{-1, -1, -1} // TODO: replace this value.
	if got := copy; got != wantCopy {
		t.Errorf("Copy: got %#v; want %#v", got, wantCopy)
	}
}
