package main

// Deferred calls run in reverse order
// Fill the complete execution trace. Include ordinary statements and all deferred calls in
// the order they actually run.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	trace := ""
	func() {
		trace += "start "
		defer func() { trace += "first " }()
		defer func() { trace += "second " }()
		trace += "return "
	}()

	wantTrace := "TODO" // TODO: replace this value.
	if got := trace; got != wantTrace {
		t.Errorf("Trace: got %#v; want %#v", got, wantTrace)
	}
}
