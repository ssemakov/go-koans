package main

// Receive from a closed channel
// Fill the value and ok result for each receive. Closing a buffered channel does not
// discard queued values; only a closed and drained channel returns zero and false.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 7
	ch <- 8
	close(ch)
	first, firstOK := <-ch
	second, secondOK := <-ch
	third, thirdOK := <-ch

	wantFirst := -999 // TODO: replace this value.
	if got := first; got != wantFirst {
		t.Errorf("First: got %#v; want %#v", got, wantFirst)
	}

	wantFirstOK := false // TODO: replace this value.
	if got := firstOK; got != wantFirstOK {
		t.Errorf("FirstOK: got %#v; want %#v", got, wantFirstOK)
	}

	wantSecond := -999 // TODO: replace this value.
	if got := second; got != wantSecond {
		t.Errorf("Second: got %#v; want %#v", got, wantSecond)
	}

	wantSecondOK := false // TODO: replace this value.
	if got := secondOK; got != wantSecondOK {
		t.Errorf("SecondOK: got %#v; want %#v", got, wantSecondOK)
	}

	wantDrained := -999 // TODO: replace this value.
	if got := third; got != wantDrained {
		t.Errorf("Drained: got %#v; want %#v", got, wantDrained)
	}

	wantDrainedOK := true // TODO: replace this value.
	if got := thirdOK; got != wantDrainedOK {
		t.Errorf("DrainedOK: got %#v; want %#v", got, wantDrainedOK)
	}
}
