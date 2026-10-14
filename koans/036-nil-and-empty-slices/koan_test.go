package main

// Nil and empty slices
// Fill the lengths and nil checks before and after append. Nil and non-nil empty slices
// can both be appended to.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	var absent []int
	empty := []int{}
	appendedAbsent := append(absent, 7)
	appendedEmpty := append(empty, 7)

	wantAbsentLength := -999 // TODO: replace this value.
	if got := len(absent); got != wantAbsentLength {
		t.Errorf("AbsentLength: got %#v; want %#v", got, wantAbsentLength)
	}

	wantEmptyLength := -999 // TODO: replace this value.
	if got := len(empty); got != wantEmptyLength {
		t.Errorf("EmptyLength: got %#v; want %#v", got, wantEmptyLength)
	}

	wantAbsentNil := false // TODO: replace this value.
	if got := absent == nil; got != wantAbsentNil {
		t.Errorf("AbsentNil: got %#v; want %#v", got, wantAbsentNil)
	}

	wantEmptyNil := true // TODO: replace this value.
	if got := empty == nil; got != wantEmptyNil {
		t.Errorf("EmptyNil: got %#v; want %#v", got, wantEmptyNil)
	}

	wantAppendedAbsentNil := true // TODO: replace this value.
	if got := appendedAbsent == nil; got != wantAppendedAbsentNil {
		t.Errorf("AppendedAbsentNil: got %#v; want %#v", got, wantAppendedAbsentNil)
	}

	wantAppendedEmptyLength := -999 // TODO: replace this value.
	if got := len(appendedEmpty); got != wantAppendedEmptyLength {
		t.Errorf("AppendedEmptyLength: got %#v; want %#v", got, wantAppendedEmptyLength)
	}
}
