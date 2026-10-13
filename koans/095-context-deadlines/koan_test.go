package main

import (
	"context"
	"testing"
	"time"
)

func TestKoan(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Unix(0, 0))
	defer cancel()
	<-ctx.Done()
	want := false // TODO: replace this value.
	got := ctx.Err() == context.DeadlineExceeded
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
