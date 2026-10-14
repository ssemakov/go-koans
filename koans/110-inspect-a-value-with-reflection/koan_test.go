package main

// Inspect a value with reflection
// Repair inspect. Return the type's Name and its Kind string, not its full type string. A
// nil interface has no type; a typed nil pointer still does.

import (
	"reflect"
	"testing"
)

type Count int

var CountValue Count = 7

func inspect(value any) (string, string) {
	typ := reflect.TypeOf(value)
	if typ == nil {
		return "", "nil"
	}
	return typ.String(), typ.String() // TODO: distinguish type name from kind.
}

func TestKoan(t *testing.T) {
	var nilPointer *Count
	cases := []struct {
		value      any
		name, kind string
	}{
		{7, "int", "int"}, {Count(7), "Count", "int"}, {&CountValue, "", "ptr"},
		{[]int{1}, "", "slice"}, {nil, "", "nil"}, {nilPointer, "", "ptr"},
	}
	for _, tc := range cases {
		name, kind := inspect(tc.value)
		if name != tc.name || kind != tc.kind {
			t.Errorf("inspect(%T)=(%q,%q); want (%q,%q)", tc.value, name, kind, tc.name, tc.kind)
		}
	}
}
