package main

import (
	"testing"
)

func TestKoan(t *testing.T) {

	want := false // TODO: replace this value.
	got := true && !false
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
