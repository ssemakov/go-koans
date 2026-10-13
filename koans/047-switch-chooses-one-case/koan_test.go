package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	result := ""
	switch 2 {
	case 1:
		result = "one"
	case 2:
		result = "two"
	default:
		result = "other"
	}
	want := "__" // TODO: replace this value.
	got := result
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
