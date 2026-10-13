package main

import (
	"strconv"
	"testing"
)

func TestKoan(t *testing.T) {
	value, err := strconv.Atoi("42")
	if err != nil {
		t.Fatal(err)
	}
	want := -999 // TODO: replace this value.
	got := value
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
