package main

// Boolean zero value
// Remove the explicit initializer so enabled receives Go's default boolean value. Do not
// replace true with false. Keep the assertion.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	var enabled bool = true

	if enabled {
		t.Error("an uninitialized boolean should be false")
	}
}
