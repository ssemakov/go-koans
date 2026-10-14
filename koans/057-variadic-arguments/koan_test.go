package main

// Variadic arguments
// Implement total as a variadic function. Zero arguments return zero; positive and
// negative values are included.

import (
	"testing"
)

func total(values ...int) int {
	return 0 // TODO: sum the variadic values.
}

func TestKoan(t *testing.T) {
	if got := total(); got != 0 {
		t.Errorf("total()=%d; want 0", got)
	}
	if got := total(7); got != 7 {
		t.Errorf("total(7)=%d; want 7", got)
	}
	if got := total(1, 2, 3); got != 6 {
		t.Errorf("total(1,2,3)=%d; want 6", got)
	}
	if got := total(-4, 2); got != -2 {
		t.Errorf("total(-4,2)=%d; want -2", got)
	}
}
