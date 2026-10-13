package main

import (
	"testing"
	"unicode/utf8"
)

func TestKoan(t *testing.T) {
	value := "é世"
	want := -999 // TODO: replace this value.
	got := utf8.RuneCountInString(value)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
