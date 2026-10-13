package main

import (
	"errors"
	"testing"
)

func TestKoan(t *testing.T) {
	err := errors.New("missing")
	want := "__" // TODO: replace this value.
	got := err.Error()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
