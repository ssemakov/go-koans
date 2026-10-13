package main

import (
	"testing"
	"time"
)

func squares(input <-chan int) <-chan int {
	output := make(chan int)
	// TODO: implement.
	close(output)
	return output
}
func TestKoan(t *testing.T) {
	input := make(chan int, 3)
	input <- 2
	input <- 3
	input <- 4
	close(input)
	output := squares(input)
	for _, want := range []int{4, 9, 16} {
		select {
		case got, ok := <-output:
			if !ok || got != want {
				t.Fatalf("got %d, open=%v; want %d", got, ok, want)
			}
		case <-time.After(time.Second):
			t.Fatal("pipeline did not produce a value")
		}
	}
	select {
	case _, ok := <-output:
		if ok {
			t.Fatal("output must close")
		}
	case <-time.After(time.Second):
		t.Fatal("output did not close")
	}
}
