package main

// Parse an integer
// Implement parseNumbers. Parse every string with strconv.Atoi. If any conversion fails,
// return no values and wrap the error with item N: using the zero-based index.

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func parseNumbers(texts []string) ([]int, error) {
	return nil, nil // TODO: parse all values and preserve the first error.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		texts    []string
		want     []int
		badIndex int
		err      error
	}{
		{nil, nil, -1, nil}, {[]string{"42", "-7", "0"}, []int{42, -7, 0}, -1, nil},
		{[]string{"1", "bad", "3"}, nil, 1, strconv.ErrSyntax}, {[]string{""}, nil, 0, strconv.ErrSyntax},
		{[]string{"99999999999999999999999999999"}, nil, 0, strconv.ErrRange},
	}
	for _, tc := range cases {
		got, err := parseNumbers(tc.texts)
		if tc.err == nil {
			if err != nil || !slices.Equal(got, tc.want) {
				t.Errorf("parseNumbers(%q)=(%v,%v); want (%v,nil)", tc.texts, got, err, tc.want)
			}
		} else {
			if got != nil || !errors.Is(err, tc.err) {
				t.Errorf("invalid input: got (%v,%v)", got, err)
			}
			prefix := fmt.Sprintf("item %d:", tc.badIndex)
			if err == nil || !strings.HasPrefix(err.Error(), prefix) {
				t.Errorf("error=%v; want prefix %q", err, prefix)
			}
		}
	}
}
