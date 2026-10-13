package main

import (
	"errors"
	"fmt"
	"testing"
)

type PathError struct{ Path string }

func (e *PathError) Error() string { return e.Path }
func TestKoan(t *testing.T) {
	err := fmt.Errorf("read: %w", &PathError{Path: "x"})
	var target *PathError
	want := false // TODO: replace this value.
	got := errors.As(err, &target)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
