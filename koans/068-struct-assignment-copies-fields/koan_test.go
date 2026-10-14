package main

// Struct assignment copies fields
// Fill the expected fields after struct assignment. Scalar fields are copied, while copied
// slices and pointers still refer to shared storage.

import (
	"slices"
	"testing"
)

type Record struct {
	Count   int
	Values  []int
	Pointer *int
}

func TestKoan(t *testing.T) {
	number := 4
	original := Record{Count: 1, Values: []int{2, 3}, Pointer: &number}
	copied := original
	copied.Count = 9
	copied.Values[0] = 7
	*copied.Pointer = 8
	wantCount := -999         // TODO: original scalar.
	wantValues := []int{-999} // TODO: original slice.
	wantPointed := -999       // TODO: original pointed-to value.
	if original.Count != wantCount {
		t.Errorf("original.Count=%d; want %d", original.Count, wantCount)
	}
	if !slices.Equal(original.Values, wantValues) {
		t.Errorf("original.Values=%v; want %v", original.Values, wantValues)
	}
	if *original.Pointer != wantPointed {
		t.Errorf("*original.Pointer=%d; want %d", *original.Pointer, wantPointed)
	}
	if copied.Count != 9 {
		t.Error("the copied scalar was not updated")
	}
}
