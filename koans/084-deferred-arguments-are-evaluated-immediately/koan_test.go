package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	value := 1
	gotValue := 0
	func() { defer func(n int) { gotValue = n }(value); value = 2 }()
	want := -999 // TODO: replace this value.
	got := gotValue
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
