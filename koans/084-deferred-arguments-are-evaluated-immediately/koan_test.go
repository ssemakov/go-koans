package main

// Deferred arguments are evaluated immediately
// Compare a deferred argument with a deferred closure reading the variable later. Then
// check how defer changes a named result before it reaches the caller.

import (
	"testing"
)

func adjusted() (result int) { defer func() { result++ }(); return 4 }

func TestKoan(t *testing.T) {
	value := 1
	argument, closure := 0, 0
	func() {
		defer func(n int) { argument = n }(value)
		defer func() { closure = value }()
		value = 2
	}()
	result := adjusted()

	wantArgument := -999 // TODO: replace this value.
	if got := argument; got != wantArgument {
		t.Errorf("Argument: got %#v; want %#v", got, wantArgument)
	}

	wantClosure := -999 // TODO: replace this value.
	if got := closure; got != wantClosure {
		t.Errorf("Closure: got %#v; want %#v", got, wantClosure)
	}

	wantNamedResult := -999 // TODO: replace this value.
	if got := result; got != wantNamedResult {
		t.Errorf("NamedResult: got %#v; want %#v", got, wantNamedResult)
	}
}
