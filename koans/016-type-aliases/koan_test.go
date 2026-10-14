package main

// Type aliases
// Complete the aliases. Count should be the same type as int, not a distinct defined type.
// Bytes should be the same type as []byte.

import (
	"reflect"
	"testing"
)

type Count int    // TODO: make Count an alias of int.
type Bytes []byte // TODO: make Bytes an alias of []byte.

func TestKoan(t *testing.T) {
	var plain int = 4
	var count Count = Count(plain)
	count += 3
	plain = int(count)
	if got := reflect.TypeOf(count).Name(); got != "int" {
		t.Errorf("Count type name=%q; want int", got)
	}
	if plain != 7 {
		t.Errorf("plain=%d; want 7", plain)
	}
	var bytes Bytes = []byte("Go")
	var ordinary []byte = bytes
	ordinary[0] = 'g'
	if reflect.TypeOf(bytes) != reflect.TypeOf(ordinary) {
		t.Error("Bytes must be the same type as []byte")
	}
	if string(bytes) != "go" {
		t.Errorf("bytes=%q; want go", bytes)
	}
}
