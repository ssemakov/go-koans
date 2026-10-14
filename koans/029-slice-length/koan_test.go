package main

// Slice length
// Track the two slice headers. Appending within capacity changes the returned slice's
// length, not the old header's length.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	values := make([]int, 2, 4)
	old := values
	values = append(values, 7)

	wantOldLength := -999 // TODO: replace this value.
	if got := len(old); got != wantOldLength {
		t.Errorf("OldLength: got %#v; want %#v", got, wantOldLength)
	}

	wantNewLength := -999 // TODO: replace this value.
	if got := len(values); got != wantNewLength {
		t.Errorf("NewLength: got %#v; want %#v", got, wantNewLength)
	}

	wantOldCapacity := -999 // TODO: replace this value.
	if got := cap(old); got != wantOldCapacity {
		t.Errorf("OldCapacity: got %#v; want %#v", got, wantOldCapacity)
	}

	wantNewCapacity := -999 // TODO: replace this value.
	if got := cap(values); got != wantNewCapacity {
		t.Errorf("NewCapacity: got %#v; want %#v", got, wantNewCapacity)
	}

	wantAppendedValue := -999 // TODO: replace this value.
	if got := old[:3][2]; got != wantAppendedValue {
		t.Errorf("AppendedValue: got %#v; want %#v", got, wantAppendedValue)
	}
}
