package main

// Capture a variable
// Compare a closure reading an outer variable with a function using a copied argument.
// Check the closure before and after changing the outer variable, then check the copied
// argument and the caller's value.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	value := 1
	captured := func() int { return value }
	copied := func(n int) int { n++; return n }

	wantBefore := -999 // TODO: replace this value.
	if got := captured(); got != wantBefore {
		t.Errorf("Before: got %#v; want %#v", got, wantBefore)
	}

	value = 5

	wantAfter := -999 // TODO: replace this value.
	if got := captured(); got != wantAfter {
		t.Errorf("After: got %#v; want %#v", got, wantAfter)
	}

	wantArgumentResult := -999 // TODO: replace this value.
	if got := copied(value); got != wantArgumentResult {
		t.Errorf("ArgumentResult: got %#v; want %#v", got, wantArgumentResult)
	}

	wantOuterValue := -999 // TODO: replace this value.
	if got := value; got != wantOuterValue {
		t.Errorf("OuterValue: got %#v; want %#v", got, wantOuterValue)
	}
}
