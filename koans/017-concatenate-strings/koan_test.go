package main

// Concatenate strings
// Implement joinName using concatenation. Add one space between nonempty parts; an empty
// part must not add a leading or trailing space.

import (
	"testing"
)

func joinName(first, last string) string {
	return first // TODO: implement.
}

func TestKoan(t *testing.T) {
	cases := []struct{ first, last, want string }{
		{"Ada", "Lovelace", "Ada Lovelace"}, {"Go", "", "Go"}, {"", "Gopher", "Gopher"}, {"", "", ""},
	}
	for _, tc := range cases {
		first, last := tc.first, tc.last
		if got := joinName(first, last); got != tc.want {
			t.Errorf("joinName(%q,%q)=%q; want %q", first, last, got, tc.want)
		}
		if first != tc.first || last != tc.last {
			t.Error("the inputs changed")
		}
	}
}
