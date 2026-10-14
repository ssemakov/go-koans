package main

// Dereference a pointer
// Implement swap through the two non-nil pointers. It must also work when both pointers
// refer to the same variable.

import (
	"testing"
)

func swap(a, b *int) { // TODO: swap the pointed-to values.
}

func TestKoan(t *testing.T) {
	for _, pair := range [][2]int{{2, 5}, {0, 7}, {-3, 4}} {
		a, b := pair[0], pair[1]
		swap(&a, &b)
		if a != pair[1] || b != pair[0] {
			t.Errorf("swap(%v)=[%d %d]", pair, a, b)
		}
	}
	value := 7
	swap(&value, &value)
	if value != 7 {
		t.Errorf("same-pointer swap=%d; want 7", value)
	}
}
