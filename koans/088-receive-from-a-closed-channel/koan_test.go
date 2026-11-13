package main

// Receive from a closed channel
// Fill the value and ok result immediately after each receive. Closing a buffered channel
// does not discard queued values; only a closed and drained channel returns zero and
// false.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 7
	ch <- 8
	close(ch)
	value, ok := <-ch

	wantFirst := -999 // TODO: replace this value.
	if got := value; got != wantFirst {
		t.Errorf("First: got %#v; want %#v", got, wantFirst)
	}

	wantFirstOK := false // TODO: replace this value.
	if got := ok; got != wantFirstOK {
		t.Errorf("FirstOK: got %#v; want %#v", got, wantFirstOK)
	}

	value, ok = <-ch

	wantSecond := -999 // TODO: replace this value.
	if got := value; got != wantSecond {
		t.Errorf("Second: got %#v; want %#v", got, wantSecond)
	}

	wantSecondOK := false // TODO: replace this value.
	if got := ok; got != wantSecondOK {
		t.Errorf("SecondOK: got %#v; want %#v", got, wantSecondOK)
	}

	value, ok = <-ch

	wantDrained := -999 // TODO: replace this value.
	if got := value; got != wantDrained {
		t.Errorf("Drained: got %#v; want %#v", got, wantDrained)
	}

	wantDrainedOK := true // TODO: replace this value.
	if got := ok; got != wantDrainedOK {
		t.Errorf("DrainedOK: got %#v; want %#v", got, wantDrainedOK)
	}
}
