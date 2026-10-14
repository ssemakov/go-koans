package main

// Pass a pointer
// Repair reset. Changing a local pointer variable does not update the object it originally
// pointed to; write through the pointer instead. Inputs are non-nil.

import (
	"testing"
)

func reset(value *int) {
	value = new(int) // TODO: reset the pointed-to integer.
}

func TestKoan(t *testing.T) {
	for _, initial := range []int{7, -3, 0} {
		value := initial
		pointer := &value
		originalAddress := pointer
		reset(pointer)
		if value != 0 {
			t.Errorf("reset(%d): value=%d; want 0", initial, value)
		}
		if pointer != originalAddress {
			t.Error("reset changed the caller's pointer")
		}
	}
}
