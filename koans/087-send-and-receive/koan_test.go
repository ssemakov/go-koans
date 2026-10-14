package main

// Send and receive
// Complete the send and receive loops. Use the supplied buffered channel to round-trip
// every string in order; return the values and the number still queued.

import (
	"slices"
	"testing"
)

func roundTrip(values []string) ([]string, int) {
	ch := make(chan string, len(values))
	for _, value := range values {
		_ = value
		// TODO: send each value to ch.
	}
	var result []string
	for range values {
		// TODO: receive one value and append it to result.
	}
	return result, len(ch)
}

func TestKoan(t *testing.T) {
	for _, values := range [][]string{nil, {}, {"Go"}, {"first", "", "last"}} {
		got, remaining := roundTrip(values)
		if !slices.Equal(got, values) {
			t.Errorf("roundTrip(%q)=%q", values, got)
		}
		if remaining != 0 {
			t.Errorf("remaining=%d; want 0", remaining)
		}
	}
}
