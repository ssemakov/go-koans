package main

// Wrap and identify an error
// Repair readNumber's error wrapping. Add parse context while preserving strconv's
// underlying error so callers can recognize syntax and range errors with errors.Is.

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func readNumber(text string) (int, error) {
	value, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("parse: %v", err) // TODO: preserve the underlying error.
	}
	return value, nil
}

func TestKoan(t *testing.T) {
	cases := []struct {
		text string
		want int
		err  error
	}{
		{"42", 42, nil}, {"-7", -7, nil}, {"bad", 0, strconv.ErrSyntax}, {"", 0, strconv.ErrSyntax},
		{"999999999999999999999999999999", 0, strconv.ErrRange},
	}
	for _, tc := range cases {
		value, err := readNumber(tc.text)
		if tc.err == nil {
			if err != nil || value != tc.want {
				t.Errorf("readNumber(%q)=(%d,%v); want (%d,nil)", tc.text, value, err, tc.want)
			}
		} else {
			if !errors.Is(err, tc.err) {
				t.Errorf("readNumber(%q): error %v does not wrap %v", tc.text, err, tc.err)
			}
			if err == nil || !strings.HasPrefix(err.Error(), "parse: ") {
				t.Errorf("missing parse context: %v", err)
			}
		}
	}
	sameMessage := errors.New(strconv.ErrSyntax.Error())
	if errors.Is(sameMessage, strconv.ErrSyntax) {
		t.Error("matching error text does not imply matching error identity")
	}
}
