package main

// Capture a variable
// Compare a closure reading an outer variable with a function using a copied argument.
// Fill the state after the outer variable changes.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	value := 1
	captured := func() int { return value }
	copied := func(n int) int { n++; return n }
	before := captured()
	value = 5
	after := captured()
	argumentResult := copied(value)

	wantBefore := -999 // TODO: replace this value.
	if got := before; got != wantBefore {
		t.Errorf("Before: got %#v; want %#v", got, wantBefore)
	}

	wantAfter := -999 // TODO: replace this value.
	if got := after; got != wantAfter {
		t.Errorf("After: got %#v; want %#v", got, wantAfter)
	}

	wantArgumentResult := -999 // TODO: replace this value.
	if got := argumentResult; got != wantArgumentResult {
		t.Errorf("ArgumentResult: got %#v; want %#v", got, wantArgumentResult)
	}

	wantOuterValue := -999 // TODO: replace this value.
	if got := value; got != wantOuterValue {
		t.Errorf("OuterValue: got %#v; want %#v", got, wantOuterValue)
	}
}
