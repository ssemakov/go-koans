package main

import (
	"fmt"
	"testing"
)

func TestKoan(t *testing.T) {

	want := "__" // TODO: replace this value.
	got := fmt.Sprintf("%s:%d", "go", 7)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
