package main

// Context deadlines
// Repair cancellable so its child retains the parent's deadline and cancellation. An
// expired deadline must remain DeadlineExceeded rather than being replaced with Canceled.

import (
	"context"
	"errors"
	"testing"
	"time"
)

func cancellable(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background()) // TODO: derive from parent.
}

func TestKoan(t *testing.T) {
	expired, stop := context.WithDeadline(context.Background(), time.Unix(0, 0))
	defer stop()
	child, cancelChild := cancellable(expired)
	defer cancelChild()
	deadline, ok := child.Deadline()
	if !ok || !deadline.Equal(time.Unix(0, 0)) {
		t.Errorf("child deadline=(%v,%v); want Unix(0,0)", deadline, ok)
	}
	select {
	case <-child.Done():
		if !errors.Is(child.Err(), context.DeadlineExceeded) {
			t.Errorf("child error=%v; want DeadlineExceeded", child.Err())
		}
	case <-time.After(time.Second):
		t.Error("expired parent did not stop the child")
	}
	parent, cancelParent := context.WithCancel(context.Background())
	other, cancelOther := cancellable(parent)
	defer cancelOther()
	cancelParent()
	select {
	case <-other.Done():
		if !errors.Is(other.Err(), context.Canceled) {
			t.Errorf("child error=%v; want Canceled", other.Err())
		}
	case <-time.After(time.Second):
		t.Error("parent cancellation did not stop the child")
	}
}
