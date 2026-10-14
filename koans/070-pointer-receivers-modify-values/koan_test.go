package main

// Pointer receivers modify values
// Repair the method receiver so Add updates the caller's Counter. Repeated positive and
// negative additions must accumulate.

import (
	"testing"
)

type Counter struct{ N int }

func (c Counter) Add(delta int) { // TODO: use a receiver that updates the caller.
	c.N += delta
}

func TestKoan(t *testing.T) {
	counter := Counter{N: 2}
	for _, step := range []struct{ delta, want int }{{3, 5}, {-2, 3}, {0, 3}} {
		counter.Add(step.delta)
		if counter.N != step.want {
			t.Errorf("Add(%d): N=%d; want %d", step.delta, counter.N, step.want)
		}
	}
}
