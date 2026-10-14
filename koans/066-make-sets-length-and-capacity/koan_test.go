package main

// Make sets length and capacity
// Complete the allocation: two zero-valued elements, with capacity for five. Append three
// values and check that an existing view observes them without changing its own length.

import (
	"slices"
	"testing"
)

func TestKoan(t *testing.T) {
	buffer := []int{} // TODO: allocate the required length and capacity.
	if len(buffer) != 2 || cap(buffer) != 5 {
		t.Fatalf("len=%d cap=%d; want len=2 cap=5", len(buffer), cap(buffer))
	}
	if !slices.Equal(buffer, []int{0, 0}) {
		t.Error("the initial elements must be zero")
	}
	old := buffer
	view := buffer[:cap(buffer)]
	buffer = append(buffer, 3, 4, 5)
	if len(old) != 2 {
		t.Error("append changed the old slice header")
	}
	if !slices.Equal(view, []int{0, 0, 3, 4, 5}) {
		t.Errorf("view=%v; want [0 0 3 4 5]", view)
	}
}
