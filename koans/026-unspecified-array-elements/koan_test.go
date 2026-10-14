package main

// Unspecified array elements
// Fill the complete expected arrays. Omitted and explicitly indexed elements leave other
// positions at their zero values.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	partial := [4]int{8, 3}
	indexed := [4]int{2: 9}

	wantPartial := [4]int{-1, -1, -1, -1} // TODO: replace this value.
	if got := partial; got != wantPartial {
		t.Errorf("Partial: got %#v; want %#v", got, wantPartial)
	}

	wantIndexed := [4]int{-1, -1, -1, -1} // TODO: replace this value.
	if got := indexed; got != wantIndexed {
		t.Errorf("Indexed: got %#v; want %#v", got, wantIndexed)
	}
}
