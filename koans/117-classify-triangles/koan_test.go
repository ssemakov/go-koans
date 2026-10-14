package main

// Classify triangles
// Implement triangle. Reject nonpositive sides and degenerate or impossible triangles;
// return equilateral, isosceles, or scalene. Side order must not matter. Positive sides
// can be near the largest int, so avoid overflowing their sum.

import (
	"errors"
	"testing"
)

func triangle(a, b, c int) (string, error) {
	return "", errors.New("not implemented") // TODO: classify after validating all sides.
}

func TestKoan(t *testing.T) {
	largest := int(^uint(0) >> 1)
	valid := []struct {
		a, b, c int
		want    string
	}{
		{2, 2, 2, "equilateral"}, {10, 10, 10, "equilateral"}, {4, 4, 3, "isosceles"},
		{1, 2, 2, "isosceles"}, {3, 4, 5, "scalene"}, {2, 3, 4, "scalene"},
		{largest, largest, largest, "equilateral"}, {largest - 2, largest - 1, largest, "scalene"},
	}
	for _, tc := range valid {
		for _, sides := range [][3]int{{tc.a, tc.b, tc.c}, {tc.a, tc.c, tc.b}, {tc.b, tc.a, tc.c}, {tc.b, tc.c, tc.a}, {tc.c, tc.a, tc.b}, {tc.c, tc.b, tc.a}} {
			got, err := triangle(sides[0], sides[1], sides[2])
			if err != nil || got != tc.want {
				t.Errorf("triangle(%v)=(%q,%v); want (%q,nil)", sides, got, err, tc.want)
			}
		}
	}
	for _, sides := range [][3]int{{0, 0, 0}, {0, 2, 2}, {2, -1, 2}, {2, 2, -1}, {1, 1, 3}, {3, 1, 1}, {2, 4, 2}, {4, 2, 2}, {largest, 1, 1}} {
		if _, err := triangle(sides[0], sides[1], sides[2]); err == nil {
			t.Errorf("triangle(%v) should return an error", sides)
		}
	}
}
