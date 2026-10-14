package main

// Create an error
// Implement requirePositive. Return the provided ErrNonpositive for zero or negative
// input, and an actual nil error for positive input.

import (
	"errors"
	"testing"
)

var ErrNonpositive = errors.New("value must be positive")

func requirePositive(value int) error {
	return nil // TODO: validate the value.
}

func TestKoan(t *testing.T) {
	for _, tc := range []struct {
		value   int
		invalid bool
	}{{-7, true}, {0, true}, {1, false}, {9, false}} {
		err := requirePositive(tc.value)
		if tc.invalid && err != ErrNonpositive {
			t.Errorf("requirePositive(%d)=%v; want ErrNonpositive", tc.value, err)
		}
		if !tc.invalid && err != nil {
			t.Errorf("requirePositive(%d)=%v; want nil", tc.value, err)
		}
	}
}
