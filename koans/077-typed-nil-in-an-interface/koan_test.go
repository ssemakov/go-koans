package main

// Typed nil in an interface
// Fill the nil and equality checks. An interface containing a nil *Failure still has a
// dynamic type; an unassigned error interface does not.

import (
	"testing"
)

type Failure struct{}

func (*Failure) Error() string { return "failure" }

func TestKoan(t *testing.T) {
	var pointer *Failure
	var boxed error = pointer
	var empty error

	wantPointerNil := false // TODO: replace this value.
	if got := pointer == nil; got != wantPointerNil {
		t.Errorf("PointerNil: got %#v; want %#v", got, wantPointerNil)
	}

	wantBoxedNil := true // TODO: replace this value.
	if got := boxed == nil; got != wantBoxedNil {
		t.Errorf("BoxedNil: got %#v; want %#v", got, wantBoxedNil)
	}

	wantEmptyNil := false // TODO: replace this value.
	if got := empty == nil; got != wantEmptyNil {
		t.Errorf("EmptyNil: got %#v; want %#v", got, wantEmptyNil)
	}

	wantSameTypedPointer := false // TODO: replace this value.
	if got := boxed == pointer; got != wantSameTypedPointer {
		t.Errorf("SameTypedPointer: got %#v; want %#v", got, wantSameTypedPointer)
	}
}
