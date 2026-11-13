package main

// New allocates a zero value
// Implement newCounter using new(int). Each call must return a non-nil pointer to a fresh
// integer initialized to zero. Updating one counter must not change another. Keep the
// checks.

import (
	"testing"
)

func newCounter() *int {
	return nil // TODO: allocate a zero-initialized int with new.
}

func TestKoan(t *testing.T) {
	allocated := newCounter()
	if allocated == nil {
		t.Fatal("newCounter must allocate an integer")
	}
	if *allocated != 0 {
		t.Errorf("new counter=%d; want 0", *allocated)
	}
	*allocated = 4

	second := newCounter()
	if second == nil {
		t.Fatal("each call must allocate an integer")
	}
	if *second != 0 {
		t.Errorf("second counter=%d; want 0 after changing the first", *second)
	}
	*second = 9
	if *allocated != 4 {
		t.Errorf("first counter=%d; want 4 after changing the second", *allocated)
	}
}
