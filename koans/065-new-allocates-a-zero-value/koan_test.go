package main

// New allocates a zero value
// Compare a nil pointer, new(int), and an initialized address. Fill both nil checks and
// values; new returns allocated zero-initialized storage.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	var absent *int
	allocated := new(int)
	initialized := 7
	address := &initialized
	second := new(int)
	*allocated = 4

	wantAbsentNil := false // TODO: replace this value.
	if got := absent == nil; got != wantAbsentNil {
		t.Errorf("AbsentNil: got %#v; want %#v", got, wantAbsentNil)
	}

	wantAllocatedNil := true // TODO: replace this value.
	if got := allocated == nil; got != wantAllocatedNil {
		t.Errorf("AllocatedNil: got %#v; want %#v", got, wantAllocatedNil)
	}

	wantAllocatedValue := -999 // TODO: replace this value.
	if got := *allocated; got != wantAllocatedValue {
		t.Errorf("AllocatedValue: got %#v; want %#v", got, wantAllocatedValue)
	}

	wantSecondValue := -999 // TODO: replace this value.
	if got := *second; got != wantSecondValue {
		t.Errorf("SecondValue: got %#v; want %#v", got, wantSecondValue)
	}

	wantInitializedValue := -999 // TODO: replace this value.
	if got := *address; got != wantInitializedValue {
		t.Errorf("InitializedValue: got %#v; want %#v", got, wantInitializedValue)
	}
}
