package main

// Read from a reader
// Implement readText with io.ReadAll. Return the text and read error, including any
// partial data received together with an error.

import (
	"errors"
	"io"
	"strings"
	"testing"
)

var ErrRead = errors.New("read failed")

type brokenReader struct{}

func (brokenReader) Read(p []byte) (int, error) { return copy(p, "partial"), ErrRead }
func readText(reader io.Reader) (string, error) {
	return "", nil // TODO: read all bytes without discarding errors or partial data.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		name   string
		reader io.Reader
		text   string
		err    error
	}{
		{"ordinary", strings.NewReader("Go"), "Go", nil}, {"empty", strings.NewReader(""), "", nil},
		{"unicode", strings.NewReader("é世"), "é世", nil}, {"partial error", brokenReader{}, "partial", ErrRead},
	}
	for _, tc := range cases {
		text, err := readText(tc.reader)
		if text != tc.text || !errors.Is(err, tc.err) {
			t.Errorf("%s: got (%q,%v); want (%q,%v)", tc.name, text, err, tc.text, tc.err)
		}
	}
}
