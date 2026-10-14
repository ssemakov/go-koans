package main

// Type constraints
// Implement generic sum. The constraint permits addition for int and float64, and ~
// includes defined types with those underlying types.

import (
	"testing"
)

type Distance int
type Weight float64
type Number interface{ ~int | ~float64 }

func sum[T Number](values []T) T {
	var zero T
	return zero // TODO: accumulate using the constraint's allowed addition.
}

func TestKoan(t *testing.T) {
	if got := sum([]int{2, 3, -1}); got != 4 {
		t.Errorf("integer sum=%d; want 4", got)
	}
	if got := sum([]float64{1.5, 2.25}); got != 3.75 {
		t.Errorf("float sum=%v; want 3.75", got)
	}
	if got := sum([]Distance{2, 5}); got != Distance(7) {
		t.Errorf("defined int sum=%v; want 7", got)
	}
	if got := sum([]Weight{0.5, 1.25}); got != Weight(1.75) {
		t.Errorf("defined float sum=%v; want 1.75", got)
	}
	if got := sum([]int(nil)); got != 0 {
		t.Errorf("empty int sum=%d; want 0", got)
	}
	if got := sum([]Weight(nil)); got != 0 {
		t.Errorf("empty Weight sum=%v; want 0", got)
	}
}
