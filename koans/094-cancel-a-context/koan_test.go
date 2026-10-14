package main

// Cancel a context
// Implement waitForValue with select. Return a received value, or return ctx.Err when
// cancellation ends the wait. A nil input disables the receive branch.

import (
	"context"
	"errors"
	"testing"
	"time"
)

func waitForValue(ctx context.Context, input <-chan int) (int, error) {
	return 0, nil // TODO: receive a value or respond to cancellation.
}

func TestKoan(t *testing.T) {
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	ready := make(chan int, 1)
	ready <- 7
	zero := make(chan int, 1)
	zero <- 0
	expired, stop := context.WithDeadline(context.Background(), time.Unix(0, 0))
	defer stop()
	cases := []struct {
		name  string
		ctx   context.Context
		input <-chan int
		value int
		err   error
	}{
		{"ready", context.Background(), ready, 7, nil}, {"zero", context.Background(), zero, 0, nil},
		{"cancelled", cancelled, nil, 0, context.Canceled}, {"expired", expired, nil, 0, context.DeadlineExceeded},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			type result struct {
				value int
				err   error
			}
			finished := make(chan result, 1)
			go func() { value, err := waitForValue(tc.ctx, tc.input); finished <- result{value, err} }()
			select {
			case got := <-finished:
				if got.value != tc.value || !errors.Is(got.err, tc.err) {
					t.Errorf("got (%d,%v); want (%d,%v)", got.value, got.err, tc.value, tc.err)
				}
			case <-time.After(time.Second):
				t.Fatal("waitForValue did not respond")
			}
		})
	}
}
