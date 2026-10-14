package main

// Recover a panic
// Implement catchPanic. Run work once; a panic becomes an error containing its value,
// while normal completion returns a real nil error. Use a deferred recovery in the same
// goroutine.

import (
	"fmt"
	"testing"
)

func catchPanic(work func()) (err error) {
	return fmt.Errorf("not implemented") // TODO: recover only when work panics.
}

func TestKoan(t *testing.T) {
	calls := 0
	if err := catchPanic(func() { calls++ }); err != nil {
		t.Errorf("normal work returned %v", err)
	}
	if calls != 1 {
		t.Errorf("calls=%d; want 1", calls)
	}
	for _, message := range []string{"stop", "bad input"} {
		err := catchPanic(func() { panic(message) })
		if err == nil || err.Error() != message {
			t.Errorf("panic(%q): error=%v", message, err)
		}
	}
	if err := catchPanic(func() { panic(7) }); err == nil || err.Error() != "7" {
		t.Errorf("numeric panic: error=%v", err)
	}
}
