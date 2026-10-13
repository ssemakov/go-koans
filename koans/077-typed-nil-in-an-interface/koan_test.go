package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	var pointer *int
	var value any = pointer
	want := true // TODO: replace this value.
	got := value == nil
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
