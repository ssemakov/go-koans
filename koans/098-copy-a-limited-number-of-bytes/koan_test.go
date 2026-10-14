package main

// Copy a limited number of bytes
// Repair copyPrefix with io.CopyN. Copy at most n bytes; if the reader ends early, return
// the bytes copied and io.EOF. n is nonnegative.

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func copyPrefix(writer io.Writer, reader io.Reader, n int64) (int64, error) {
	return io.Copy(writer, reader) // TODO: limit the copy to n bytes.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		text  string
		n     int64
		want  string
		count int64
		err   error
	}{
		{"hello world", 5, "hello", 5, nil}, {"hello", 0, "", 0, nil},
		{"Go", 2, "Go", 2, nil}, {"Go", 5, "Go", 2, io.EOF}, {"", 1, "", 0, io.EOF},
	}
	for _, tc := range cases {
		var out bytes.Buffer
		count, err := copyPrefix(&out, strings.NewReader(tc.text), tc.n)
		if count != tc.count || !errors.Is(err, tc.err) || out.String() != tc.want {
			t.Errorf("text=%q n=%d: got (%d,%v,%q); want (%d,%v,%q)", tc.text, tc.n, count, err, out.String(), tc.count, tc.err, tc.want)
		}
	}
}
