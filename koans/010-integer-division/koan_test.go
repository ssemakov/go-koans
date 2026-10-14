package main

// Integer division
// Fill the quotient and remainder values for both signs, then compare floating-point
// division.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	positive, negative, divisor := 7, -7, 2
	floating := float64(positive) / float64(divisor)

	wantPositive := -999 // TODO: replace this value.
	if got := positive / divisor; got != wantPositive {
		t.Errorf("Positive: got %#v; want %#v", got, wantPositive)
	}

	wantNegative := -999 // TODO: replace this value.
	if got := negative / divisor; got != wantNegative {
		t.Errorf("Negative: got %#v; want %#v", got, wantNegative)
	}

	wantNegativeRemainder := -999 // TODO: replace this value.
	if got := negative % divisor; got != wantNegativeRemainder {
		t.Errorf("NegativeRemainder: got %#v; want %#v", got, wantNegativeRemainder)
	}

	wantFloating := -999.0 // TODO: replace this value.
	if got := floating; got != wantFloating {
		t.Errorf("Floating: got %#v; want %#v", got, wantFloating)
	}
}
