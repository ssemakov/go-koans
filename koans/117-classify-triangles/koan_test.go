package main

import (
	"errors"
	"testing"
)

func triangle(a, b, c int) (string, error) {
	// TODO: implement.
	return "", errors.New("not implemented")
}
func TestKoan(t *testing.T) {
	cases := []struct {
		a, b, c int
		want    string
	}{{2, 2, 2, "equilateral"}, {4, 4, 3, "isosceles"}, {3, 4, 5, "scalene"}, {4, 3, 4, "isosceles"}}
	for _, tc := range cases {
		got, err := triangle(tc.a, tc.b, tc.c)
		if err != nil || got != tc.want {
			t.Fatalf("triangle(%d,%d,%d) = %q, %v; want %q", tc.a, tc.b, tc.c, got, err, tc.want)
		}
	}
	for _, tc := range [][3]int{{0, 0, 0}, {-1, 2, 2}, {1, 1, 3}, {2, 4, 2}} {
		if _, err := triangle(tc[0], tc[1], tc[2]); err == nil {
			t.Fatalf("expected error for %v", tc)
		}
	}
}
