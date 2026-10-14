package main

// Select with a default
// Implement tryReceive with select and default. Return value, ready, and open: a closed
// channel is ready but not open; an unready or nil channel returns zero, false, false.

import (
	"testing"
	"time"
)

func tryReceive(ch <-chan int) (int, bool, bool) {
	return 0, false, false // TODO: perform a nonblocking receive.
}

func TestKoan(t *testing.T) {
	buffered := make(chan int, 2)
	type result struct {
		value       int
		ready, open bool
	}
	buffered <- 7
	buffered <- 0
	closed := make(chan int)
	close(closed)
	cases := []struct {
		name        string
		ch          <-chan int
		value       int
		ready, open bool
	}{
		{"nil", nil, 0, false, false}, {"unready", make(chan int), 0, false, false},
		{"buffered", buffered, 7, true, true}, {"buffered zero", buffered, 0, true, true},
		{"closed", closed, 0, true, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			finished := make(chan result, 1)
			go func() {
				value, ready, open := tryReceive(tc.ch)
				finished <- result{value, ready, open}
			}()
			select {
			case got := <-finished:
				if got.value != tc.value || got.ready != tc.ready || got.open != tc.open {
					t.Errorf("got %+v; want (value=%d,ready=%v,open=%v)", got, tc.value, tc.ready, tc.open)
				}
			case <-time.After(time.Second):
				t.Fatal("tryReceive blocked")
			}
		})
	}
}
