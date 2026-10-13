package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	value := "gopher"
	want := "__" // TODO: replace this value.
	got := value[2:5]
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
