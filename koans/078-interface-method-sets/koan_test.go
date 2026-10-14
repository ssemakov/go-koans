package main

// Interface method sets
// Repair the receiver so only *Counter implements Incrementer. The pointer must mutate its
// counter; a Counter value must fail the runtime interface assertion.

import (
	"testing"
)

type Incrementer interface{ Increment() }
type Counter struct{ N int }

func (c Counter) Increment() { // TODO: choose the pointer receiver.
	c.N++
}

func TestKoan(t *testing.T) {
	_, valueOK := any(Counter{}).(Incrementer)
	if valueOK {
		t.Error("Counter value must not implement Incrementer")
	}
	counter := &Counter{}
	incrementer, pointerOK := any(counter).(Incrementer)
	if !pointerOK {
		t.Fatal("*Counter must implement Incrementer")
	}
	incrementer.Increment()
	incrementer.Increment()
	if counter.N != 2 {
		t.Errorf("counter.N=%d; want 2", counter.N)
	}
}
