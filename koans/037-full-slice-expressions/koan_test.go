package main

// Full slice expressions
// Compare an unrestricted view with a full slice expression that caps capacity. Fill the
// capacity and original values after append.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	array := [4]int{1, 2, 3, 4}
	unrestricted := array[1:2]
	limited := array[1:2:2]
	limitedCapacity := cap(limited)
	unrestricted = append(unrestricted, 8)
	limited = append(limited, 9)
	limited[0] = 7

	wantUnrestrictedCapacity := -999 // TODO: replace this value.
	if got := cap(unrestricted); got != wantUnrestrictedCapacity {
		t.Errorf("UnrestrictedCapacity: got %#v; want %#v", got, wantUnrestrictedCapacity)
	}

	wantLimitedCapacity := -999 // TODO: replace this value.
	if got := limitedCapacity; got != wantLimitedCapacity {
		t.Errorf("LimitedCapacity: got %#v; want %#v", got, wantLimitedCapacity)
	}

	wantLimitedLength := -999 // TODO: replace this value.
	if got := len(limited); got != wantLimitedLength {
		t.Errorf("LimitedLength: got %#v; want %#v", got, wantLimitedLength)
	}

	wantOriginalArray := [4]int{-1, -1, -1, -1} // TODO: replace this value.
	if got := array; got != wantOriginalArray {
		t.Errorf("OriginalArray: got %#v; want %#v", got, wantOriginalArray)
	}

	wantLimitedFirst := -999 // TODO: replace this value.
	if got := limited[0]; got != wantLimitedFirst {
		t.Errorf("LimitedFirst: got %#v; want %#v", got, wantLimitedFirst)
	}

	wantLimitedSecond := -999 // TODO: replace this value.
	if got := limited[1]; got != wantLimitedSecond {
		t.Errorf("LimitedSecond: got %#v; want %#v", got, wantLimitedSecond)
	}
}
