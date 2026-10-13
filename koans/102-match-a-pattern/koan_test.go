package main

import (
	"regexp"
	"testing"
)

func TestKoan(t *testing.T) {
	pattern := regexp.MustCompile(`^go[0-9]+$`)
	want := false // TODO: replace this value.
	got := pattern.MatchString("go123")
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
