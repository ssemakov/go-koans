package main

// Simultaneous assignment
// Repair the swap using one simultaneous assignment. Both right-hand values must come from
// the original pair.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	for _, pair := range [][2]int{{2, 5}, {0, 7}, {-3, 4}, {8, 8}} {
		a, b := pair[0], pair[1]
		// TODO: replace the two assignments with a simultaneous swap.
		a = b
		b = a
		if a != pair[1] || b != pair[0] {
			t.Errorf("swap(%v) = [%d %d]; want [%d %d]", pair, a, b, pair[1], pair[0])
		}
	}
}
