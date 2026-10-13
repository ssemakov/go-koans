package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	left := "go"
	want := "__" // TODO: replace this value.
	got := left + "pher"
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
