package main

// Numeric conversion
// Fill the results for both signs. Compare conversion to int with math.Floor; they differ
// for negative fractions.

import (
	"math"
	"testing"
)

func TestKoan(t *testing.T) {
	positive, negative := 3.75, -3.75

	wantPositiveInt := -999 // TODO: replace this value.
	if got := int(positive); got != wantPositiveInt {
		t.Errorf("PositiveInt: got %#v; want %#v", got, wantPositiveInt)
	}

	wantNegativeInt := -999 // TODO: replace this value.
	if got := int(negative); got != wantNegativeInt {
		t.Errorf("NegativeInt: got %#v; want %#v", got, wantNegativeInt)
	}

	wantPositiveFloor := -999 // TODO: replace this value.
	if got := int(math.Floor(positive)); got != wantPositiveFloor {
		t.Errorf("PositiveFloor: got %#v; want %#v", got, wantPositiveFloor)
	}

	wantNegativeFloor := -999 // TODO: replace this value.
	if got := int(math.Floor(negative)); got != wantNegativeFloor {
		t.Errorf("NegativeFloor: got %#v; want %#v", got, wantNegativeFloor)
	}
}
