package main

// Slice an array
// Fill the length, capacity, and original array after a slice is modified. Capacity starts
// at the slice's first element and extends to the end of the array.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	part := array[1:3]
	part[0] = 9

	wantLength := -999 // TODO: replace this value.
	if got := len(part); got != wantLength {
		t.Errorf("Length: got %#v; want %#v", got, wantLength)
	}

	wantCapacity := -999 // TODO: replace this value.
	if got := cap(part); got != wantCapacity {
		t.Errorf("Capacity: got %#v; want %#v", got, wantCapacity)
	}

	wantArray := [5]int{-1, -1, -1, -1, -1} // TODO: replace this value.
	if got := array; got != wantArray {
		t.Errorf("Array: got %#v; want %#v", got, wantArray)
	}
}
