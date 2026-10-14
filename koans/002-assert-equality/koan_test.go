package main

// Assert equality
// Fill both expected values. t.Errorf reports a failed check and continues, so one run can
// report both failures.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	wantSum := -999 // TODO: replace this value.
	if got := 1 + 1; got != wantSum {
		t.Errorf("Sum: got %#v; want %#v", got, wantSum)
	}

	wantDifferent := true // TODO: replace this value.
	if got := 2 == 3; got != wantDifferent {
		t.Errorf("Different: got %#v; want %#v", got, wantDifferent)
	}
}
