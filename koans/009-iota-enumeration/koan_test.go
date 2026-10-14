package main

// Iota enumeration
// Fill the values for each constant block. Check a skipped entry, a reset, and bit flags
// generated with shifts.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	const (
		first = iota
		_
		third
	)
	const (
		reset = iota
		next
	)
	const (
		read = 1 << iota
		write
		execute
	)

	wantFirst := -999 // TODO: replace this value.
	if got := first; got != wantFirst {
		t.Errorf("First: got %#v; want %#v", got, wantFirst)
	}

	wantThird := -999 // TODO: replace this value.
	if got := third; got != wantThird {
		t.Errorf("Third: got %#v; want %#v", got, wantThird)
	}

	wantReset := -999 // TODO: replace this value.
	if got := reset; got != wantReset {
		t.Errorf("Reset: got %#v; want %#v", got, wantReset)
	}

	wantNext := -999 // TODO: replace this value.
	if got := next; got != wantNext {
		t.Errorf("Next: got %#v; want %#v", got, wantNext)
	}

	wantCombinedFlags := -999 // TODO: replace this value.
	if got := read | execute; got != wantCombinedFlags {
		t.Errorf("CombinedFlags: got %#v; want %#v", got, wantCombinedFlags)
	}
}
