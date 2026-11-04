package main

// Variadic arguments
// Predict the results of calling total with zero, one, and several arguments. Inside
// total, values is a slice of all supplied integers. Change only the expected
// values; keep the variadic declaration, calls, and checks.

import "testing"

func total(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func TestKoan(t *testing.T) {
	wantEmpty := -999 // TODO: predict total().
	if got := total(); got != wantEmpty {
		t.Errorf("total()=%d; want %d", got, wantEmpty)
	}
	wantOne := -999 // TODO: predict total(7).
	if got := total(7); got != wantOne {
		t.Errorf("total(7)=%d; want %d", got, wantOne)
	}
	wantMany := -999 // TODO: predict total(1, 2, 3).
	if got := total(1, 2, 3); got != wantMany {
		t.Errorf("total(1,2,3)=%d; want %d", got, wantMany)
	}
	wantSigned := -999 // TODO: predict total(-4, 2).
	if got := total(-4, 2); got != wantSigned {
		t.Errorf("total(-4,2)=%d; want %d", got, wantSigned)
	}
}
