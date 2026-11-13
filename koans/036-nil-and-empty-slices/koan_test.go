package main

// Nil and empty slices
// Complete the appends to the nil slice and the non-nil empty slice. Both should contain 7
// afterward. Keep the checks that show their initial difference and their contents after
// append.

import (
	"slices"
	"testing"
)

func TestKoan(t *testing.T) {
	var absent []int
	if absent != nil || len(absent) != 0 {
		t.Fatal("a nil slice starts nil with length zero")
	}
	// TODO: append 7 to absent and store the returned slice in absent.
	if absent == nil || !slices.Equal(absent, []int{7}) {
		t.Errorf("after append, absent=%v; want a non-nil slice containing [7]", absent)
	}

	empty := []int{}
	if empty == nil || len(empty) != 0 {
		t.Fatal("an empty slice literal starts non-nil with length zero")
	}
	// TODO: append 7 to empty and store the returned slice in empty.
	if empty == nil || !slices.Equal(empty, []int{7}) {
		t.Errorf("after append, empty=%v; want a non-nil slice containing [7]", empty)
	}
}
