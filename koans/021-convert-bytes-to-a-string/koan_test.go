package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	value := []byte{71, 111}
	want := "__" // TODO: replace this value.
	got := string(value)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
