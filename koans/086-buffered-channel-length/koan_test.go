package main

// Buffered channel length
// Track queue length separately from capacity as values are sent and received. Fill the
// first received value and each queue length.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	ch := make(chan int, 2)
	initial := len(ch)
	ch <- 7
	ch <- 0
	filled := len(ch)
	first := <-ch
	remaining := len(ch)

	wantInitial := -999 // TODO: replace this value.
	if got := initial; got != wantInitial {
		t.Errorf("Initial: got %#v; want %#v", got, wantInitial)
	}

	wantCapacity := -999 // TODO: replace this value.
	if got := cap(ch); got != wantCapacity {
		t.Errorf("Capacity: got %#v; want %#v", got, wantCapacity)
	}

	wantFilled := -999 // TODO: replace this value.
	if got := filled; got != wantFilled {
		t.Errorf("Filled: got %#v; want %#v", got, wantFilled)
	}

	wantFirst := -999 // TODO: replace this value.
	if got := first; got != wantFirst {
		t.Errorf("First: got %#v; want %#v", got, wantFirst)
	}

	wantRemaining := -999 // TODO: replace this value.
	if got := remaining; got != wantRemaining {
		t.Errorf("Remaining: got %#v; want %#v", got, wantRemaining)
	}
}
