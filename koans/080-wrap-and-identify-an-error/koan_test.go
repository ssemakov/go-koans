package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestKoan(t *testing.T) {
	base := errors.New("missing")
	err := fmt.Errorf("read: %w", base)
	want := false // TODO: replace this value.
	got := errors.Is(err, base)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
