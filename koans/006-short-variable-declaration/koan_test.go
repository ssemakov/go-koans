package main

// Short variable declaration
// Repair the assignment inside the block. Update the existing total instead of declaring a
// shadowing variable. Keep the checks.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	for _, delta := range []int{1, 4, -2} {
		total := 10
		{
			total := total + delta // TODO: update the outer variable.
			_ = total
		}
		if total != 10+delta {
			t.Errorf("delta=%d: total=%d; want %d", delta, total, 10+delta)
		}
	}
}
