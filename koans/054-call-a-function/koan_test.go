package main

// Call a function
// Complete the two calls to remaining. Pass capacity and booked in that order, then
// use the first result with held in the second call. Keep the helper and checks
// unchanged.

import "testing"

func remaining(total, used int) int { return total - used }

func TestKoan(t *testing.T) {
	cases := []struct{ capacity, booked, held, available, afterHold int }{
		{12, 5, 2, 7, 5},
		{4, 0, 4, 4, 0},
		{10, 10, 0, 0, 0},
	}
	for _, tc := range cases {
		available := 0 // TODO: call remaining with tc.capacity and tc.booked.
		afterHold := 0 // TODO: call remaining with available and tc.held.
		if available != tc.available || afterHold != tc.afterHold {
			t.Errorf("capacity=%d booked=%d held=%d: available=%d afterHold=%d; want %d and %d",
				tc.capacity, tc.booked, tc.held, available, afterHold, tc.available, tc.afterHold)
		}
	}
}
