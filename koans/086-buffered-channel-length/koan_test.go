package main

// Buffered channel length
// Track queue length separately from capacity as values are sent and received. Fill the
// first received value and each queue length at the point where it is checked.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	ch := make(chan int, 2)

	wantInitial := -999 // TODO: replace this value.
	if got := len(ch); got != wantInitial {
		t.Errorf("Initial: got %#v; want %#v", got, wantInitial)
	}

	wantCapacity := -999 // TODO: replace this value.
	if got := cap(ch); got != wantCapacity {
		t.Errorf("Capacity: got %#v; want %#v", got, wantCapacity)
	}

	ch <- 7
	ch <- 0

	wantFilled := -999 // TODO: replace this value.
	if got := len(ch); got != wantFilled {
		t.Errorf("Filled: got %#v; want %#v", got, wantFilled)
	}

	wantFirst := -999 // TODO: replace this value.
	if got := <-ch; got != wantFirst {
		t.Errorf("First: got %#v; want %#v", got, wantFirst)
	}

	wantRemaining := -999 // TODO: replace this value.
	if got := len(ch); got != wantRemaining {
		t.Errorf("Remaining: got %#v; want %#v", got, wantRemaining)
	}
}
