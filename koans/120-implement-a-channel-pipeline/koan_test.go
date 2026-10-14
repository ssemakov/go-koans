package main

// Implement a channel pipeline
// Implement squares as a streaming pipeline. Return an output channel immediately; a
// goroutine squares and sends each input in order, then closes output when input closes.
// Consumers drain output, and input values fit safely when squared.

import (
	"slices"
	"testing"
	"time"
)

func squares(input <-chan int) <-chan int {
	output := make(chan int)
	close(output) // TODO: transform input in a goroutine, then close output.
	return output
}

func TestKoan(t *testing.T) {
	t.Run("closed empty", func(t *testing.T) {
		input := make(chan int)
		close(input)
		returned := make(chan (<-chan int), 1)
		go func() { returned <- squares(input) }()
		var output <-chan int
		select {
		case output = <-returned:
		case <-time.After(time.Second):
			t.Fatal("squares did not return for closed input")
		}
		select {
		case _, ok := <-output:
			if ok {
				t.Error("empty input must not produce a value")
			}
		case <-time.After(time.Second):
			t.Fatal("empty output did not close")
		}
	})
	t.Run("buffered values", func(t *testing.T) {
		input := make(chan int, 5)
		for _, value := range []int{2, 3, 4, -2, 0} {
			input <- value
		}
		close(input)
		finished := make(chan []int, 1)
		go func() {
			var values []int
			for value := range squares(input) {
				values = append(values, value)
			}
			finished <- values
		}()
		select {
		case got := <-finished:
			if !slices.Equal(got, []int{4, 9, 16, 4, 0}) {
				t.Errorf("squares=%v; want [4 9 16 4 0]", got)
			}
		case <-time.After(time.Second):
			t.Fatal("buffered pipeline did not finish")
		}
	})
	t.Run("stream before input closes", func(t *testing.T) {
		input := make(chan int)
		inputClosed := false
		defer func() {
			if !inputClosed {
				close(input)
			}
		}()
		returned := make(chan (<-chan int), 1)
		go func() { returned <- squares(input) }()
		var output <-chan int
		select {
		case output = <-returned:
		case <-time.After(time.Second):
			t.Fatal("squares waited for input before returning")
		}
		for _, value := range []int{3, -2, 0, 5} {
			select {
			case input <- value:
			case <-time.After(time.Second):
				t.Fatal("pipeline did not receive the next input")
			}
			select {
			case got, ok := <-output:
				if !ok || got != value*value {
					t.Errorf("input=%d: got (%d,%v); want (%d,true)", value, got, ok, value*value)
				}
			case <-time.After(time.Second):
				t.Fatal("pipeline waited for input closure before producing output")
			}
		}
		close(input)
		inputClosed = true
		select {
		case _, ok := <-output:
			if ok {
				t.Error("pipeline produced an extra value")
			}
		case <-time.After(time.Second):
			t.Fatal("output did not close after input")
		}
	})
}
