package main

// Synchronize a goroutine
// Implement squareAsync. Launch a goroutine that sends the square on the supplied
// unbuffered output, then closes it. Return the channel before the receiver starts.

import (
	"testing"
	"time"
)

func squareAsync(value int) <-chan int {
	output := make(chan int)
	close(output) // TODO: send and close from a goroutine.
	return output
}

func TestKoan(t *testing.T) {
	for _, tc := range []struct {
		name        string
		value, want int
	}{{"positive", 3, 9}, {"negative", -4, 16}, {"zero", 0, 0}} {
		t.Run(tc.name, func(t *testing.T) {
			returned := make(chan (<-chan int), 1)
			go func() { returned <- squareAsync(tc.value) }()
			var output <-chan int
			select {
			case output = <-returned:
			case <-time.After(time.Second):
				t.Fatal("squareAsync did not return before a receiver started")
			}
			select {
			case got, ok := <-output:
				if !ok || got != tc.want {
					t.Errorf("got (%d,%v); want (%d,true)", got, ok, tc.want)
				}
			case <-time.After(time.Second):
				t.Fatal("no squared value arrived")
			}
			select {
			case _, ok := <-output:
				if ok {
					t.Error("output should close after one value")
				}
			case <-time.After(time.Second):
				t.Fatal("output did not close")
			}
		})
	}
}
