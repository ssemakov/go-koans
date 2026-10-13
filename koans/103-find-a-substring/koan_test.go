package main

import (
	"regexp"
	"testing"
)

func TestKoan(t *testing.T) {
	pattern := regexp.MustCompile(`[0-9]+`)
	want := "__" // TODO: replace this value.
	got := pattern.FindString("go123x")
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
