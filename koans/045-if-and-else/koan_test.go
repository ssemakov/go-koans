package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	result := ""
	if 3 > 2 {
		result = "yes"
	} else {
		result = "no"
	}
	want := "__" // TODO: replace this value.
	got := result
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
