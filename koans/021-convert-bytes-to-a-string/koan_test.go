package main

// Convert bytes to a string
// Complete the two conversions. Keep the original string independent of later byte edits,
// and make an editable byte copy of it.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	data := []byte("Go")
	original := "" // TODO: convert data to a string.
	data[0] = 'N'
	editable := []byte{} // TODO: convert original to bytes.
	if original != "Go" {
		t.Fatalf("original=%q; want Go after mutating data", original)
	}
	if len(editable) != 2 {
		t.Fatalf("editable=%v; want two bytes", editable)
	}
	editable[1] = '!'
	if string(data) != "No" {
		t.Errorf("data=%q; want No", data)
	}
	if string(editable) != "G!" {
		t.Errorf("editable=%q; want G!", editable)
	}
	if original != "Go" {
		t.Errorf("editing a byte copy changed original=%q", original)
	}
}
