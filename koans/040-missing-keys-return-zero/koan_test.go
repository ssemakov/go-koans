package main

// Missing keys return zero
// Fill both values and presence flags. A missing key and a present key storing zero have
// the same value but different presence.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	counts := map[string]int{"zero": 0, "go": 2}
	zero, hasZero := counts["zero"]
	missing, hasMissing := counts["missing"]

	wantZeroValue := -999 // TODO: replace this value.
	if got := zero; got != wantZeroValue {
		t.Errorf("ZeroValue: got %#v; want %#v", got, wantZeroValue)
	}

	wantMissingValue := -999 // TODO: replace this value.
	if got := missing; got != wantMissingValue {
		t.Errorf("MissingValue: got %#v; want %#v", got, wantMissingValue)
	}

	wantZeroPresent := false // TODO: replace this value.
	if got := hasZero; got != wantZeroPresent {
		t.Errorf("ZeroPresent: got %#v; want %#v", got, wantZeroPresent)
	}

	wantMissingPresent := true // TODO: replace this value.
	if got := hasMissing; got != wantMissingPresent {
		t.Errorf("MissingPresent: got %#v; want %#v", got, wantMissingPresent)
	}
}
