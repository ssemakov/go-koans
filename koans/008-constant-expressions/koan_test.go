package main

// Constant expressions
// Fill the conversion and declaration. An untyped constant can initialize several
// compatible types; a typed variable needs an explicit conversion.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	const unit = 3
	var small int8 = unit
	var wide int64 = unit
	typed := int8(7)
	converted := int64(0) // TODO: convert typed to int64.
	_ = typed
	if small != 3 || wide != 3 {
		t.Error("constant initialization changed the value")
	}
	if converted != 7 {
		t.Errorf("converted=%d; want 7", converted)
	}
	const third = 0 // TODO: use an exact constant expression: 1.0 / 3.
	const restored = third * 3
	if restored != 1 {
		t.Errorf("restored=%v; want 1", restored)
	}
}
