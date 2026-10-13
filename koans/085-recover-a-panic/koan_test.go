package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	message := ""
	func() {
		defer func() {
			if r := recover(); r != nil {
				message = r.(string)
			}
		}()
		panic("stop")
	}()
	want := "__" // TODO: replace this value.
	got := message
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
