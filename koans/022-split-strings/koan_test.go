package main

import (
	"strings"
	"testing"
)

func TestKoan(t *testing.T) {
	parts := strings.Split("a,b,c", ",")
	want := -999 // TODO: replace this value.
	got := len(parts)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
