package main

// Value receivers use a copy
// Fill the caller state and method result. A value receiver copies the struct, but its
// slice field still shares an array.

import (
	"testing"
)

type Counter struct {
	N      int
	Values []int
}

func (c Counter) Change() int { c.N++; c.Values[0]++; return c.N }

func TestKoan(t *testing.T) {
	counter := Counter{N: 2, Values: []int{4}}
	returned := counter.Change()

	wantCallerNumber := -999 // TODO: replace this value.
	if got := counter.N; got != wantCallerNumber {
		t.Errorf("CallerNumber: got %#v; want %#v", got, wantCallerNumber)
	}

	wantReturnedNumber := -999 // TODO: replace this value.
	if got := returned; got != wantReturnedNumber {
		t.Errorf("ReturnedNumber: got %#v; want %#v", got, wantReturnedNumber)
	}

	wantSharedElement := -999 // TODO: replace this value.
	if got := counter.Values[0]; got != wantSharedElement {
		t.Errorf("SharedElement: got %#v; want %#v", got, wantSharedElement)
	}
}
