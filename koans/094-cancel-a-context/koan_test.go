package main

import (
	"context"
	"testing"
)

func TestKoan(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	<-ctx.Done()
	want := false // TODO: replace this value.
	got := ctx.Err() == context.Canceled
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
