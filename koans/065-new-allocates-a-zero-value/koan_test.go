package main

// New allocates a zero value
// Compare a nil pointer, new(int), and an initialized address. Check each nil condition
// and value before the next assignment. Each new allocation starts at zero.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	var absent *int
	wantAbsentNil := false // TODO: replace this value.
	if got := absent == nil; got != wantAbsentNil {
		t.Errorf("AbsentNil: got %#v; want %#v", got, wantAbsentNil)
	}

	allocated := new(int)
	wantAllocatedNil := true // TODO: replace this value.
	if got := allocated == nil; got != wantAllocatedNil {
		t.Errorf("AllocatedNil: got %#v; want %#v", got, wantAllocatedNil)
	}

	wantAllocatedValue := -999 // TODO: replace this value.
	if got := *allocated; got != wantAllocatedValue {
		t.Errorf("AllocatedValue: got %#v; want %#v", got, wantAllocatedValue)
	}

	*allocated = 4
	wantAssignedValue := -999 // TODO: replace this value.
	if got := *allocated; got != wantAssignedValue {
		t.Errorf("AssignedValue: got %#v; want %#v", got, wantAssignedValue)
	}

	second := new(int)
	wantSecondValue := -999 // TODO: replace this value.
	if got := *second; got != wantSecondValue {
		t.Errorf("SecondValue: got %#v; want %#v", got, wantSecondValue)
	}

	initialized := 7
	address := &initialized
	wantInitializedValue := -999 // TODO: replace this value.
	if got := *address; got != wantInitializedValue {
		t.Errorf("InitializedValue: got %#v; want %#v", got, wantInitializedValue)
	}
}
