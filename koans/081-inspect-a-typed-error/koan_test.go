package main

// Inspect a typed error
// Implement errorPath with errors.As. Find a *PathError through wrapping and return its
// Path; unrelated errors, nil errors, or a typed nil *PathError return empty and false.

import (
	"errors"
	"fmt"
	"testing"
)

type PathError struct{ Path string }

func (e *PathError) Error() string { return e.Path }
func errorPath(err error) (string, bool) {
	return "", false // TODO: extract the typed error and its path.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		err  error
		path string
		ok   bool
	}{
		{&PathError{Path: "a.txt"}, "a.txt", true},
		{fmt.Errorf("read: %w", &PathError{Path: "b.txt"}), "b.txt", true},
		{errors.New("a.txt"), "", false}, {nil, "", false}, {(*PathError)(nil), "", false},
	}
	for _, tc := range cases {
		path, ok := errorPath(tc.err)
		if path != tc.path || ok != tc.ok {
			t.Errorf("errorPath(%v)=(%q,%v); want (%q,%v)", tc.err, path, ok, tc.path, tc.ok)
		}
	}
}
