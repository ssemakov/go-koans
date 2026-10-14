package main

// Array length
// Complete the array declaration. It must hold exactly three readings, including an
// omitted zero-valued reading, and match the function's [3]int parameter.

import (
	"testing"
)

func totalReadings(values [3]int) int { return values[0] + values[1] + values[2] }

func TestKoan(t *testing.T) {
	readings := [3]int{} // TODO: initialize the first two readings to 4 and 7.
	if got := totalReadings(readings); got != 11 {
		t.Errorf("total=%d; want 11", got)
	}
	if readings != [3]int{4, 7, 0} {
		t.Errorf("readings=%v; want [4 7 0]", readings)
	}
	if len(readings) != 3 || cap(readings) != 3 {
		t.Error("an array has a fixed length and capacity")
	}
}
