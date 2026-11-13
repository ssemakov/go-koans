package main

// Methods on named types
// Implement the Kilometers method on Meters, then call it on tc.meters in the test. Use
// the receiver value and a floating-point conversion; do not lose the fractional
// kilometer. Keep the checks.

import (
	"testing"
)

type Meters int

func (m Meters) Kilometers() float64 {
	return 0 // TODO: use the receiver and preserve the fraction.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		meters Meters
		want   float64
	}{{0, 0}, {1000, 1}, {1500, 1.5}, {250, 0.25}}
	for _, tc := range cases {
		got := float64(0) // TODO: call the Kilometers method on tc.meters.
		if got != tc.want {
			t.Errorf("%d meters=%v km; want %v", tc.meters, got, tc.want)
		}
	}
}
