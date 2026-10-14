package main

// Range over a channel
// Implement drain with range over the channel. Return all queued values, including zero,
// and finish when the sender closes the channel.

import (
	"slices"
	"testing"
	"time"
)

func drain(input <-chan int) []int {
	return nil // TODO: receive until the input closes.
}

func TestKoan(t *testing.T) {
	for _, values := range [][]int{nil, {0}, {2, 0, 3}} {
		input := make(chan int, len(values))
		for _, value := range values {
			input <- value
		}
		close(input)
		finished := make(chan []int, 1)
		go func() { finished <- drain(input) }()
		select {
		case got := <-finished:
			if !slices.Equal(got, values) {
				t.Errorf("drain(%v)=%v", values, got)
			}
		case <-time.After(time.Second):
			t.Fatal("drain did not stop at channel closure")
		}
	}
}
