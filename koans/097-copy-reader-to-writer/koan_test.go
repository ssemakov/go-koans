package main

// Copy reader to writer
// Implement copyAll with io.Copy. Return the actual byte count and preserve reader or
// writer errors instead of treating partial copying as success.

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

var ErrRead = errors.New("read failed")
var ErrWrite = errors.New("write failed")

type brokenReader struct{}

func (brokenReader) Read(p []byte) (int, error) { return copy(p, "partial"), ErrRead }

type brokenWriter struct{}

func (brokenWriter) Write([]byte) (int, error) { return 0, ErrWrite }
func copyAll(writer io.Writer, reader io.Reader) (int64, error) {
	return 0, nil // TODO: copy and return both results.
}

func TestKoan(t *testing.T) {
	for _, text := range []string{"", "hello", "é世"} {
		var out bytes.Buffer
		count, err := copyAll(&out, strings.NewReader(text))
		if err != nil || count != int64(len(text)) || out.String() != text {
			t.Errorf("copyAll(%q)=(%d,%v,%q)", text, count, err, out.String())
		}
	}
	count, err := copyAll(brokenWriter{}, strings.NewReader("hello"))
	if count != 0 || !errors.Is(err, ErrWrite) {
		t.Errorf("writer failure=(%d,%v); want (0,ErrWrite)", count, err)
	}
	var partial bytes.Buffer
	count, err = copyAll(&partial, brokenReader{})
	if count != 7 || !errors.Is(err, ErrRead) || partial.String() != "partial" {
		t.Errorf("partial read=(%d,%v,%q)", count, err, partial.String())
	}
}
