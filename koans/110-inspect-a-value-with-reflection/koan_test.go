package main

import (
	"reflect"
	"testing"
)

func TestKoan(t *testing.T) {
	value := reflect.TypeOf(7)
	want := false // TODO: replace this value.
	got := value.Kind() == reflect.Int
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
