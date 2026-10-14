package main

// Promoted methods
// Implement Tagged.Label to override the promoted method. Append a nonempty tag in
// brackets; keep Base.Label available without the tag.

import (
	"testing"
)

type Base struct{ Name string }

func (b Base) Label() string { return b.Name }

type Tagged struct {
	Base
	Tag string
}

func (t Tagged) Label() string {
	return t.Base.Label() // TODO: add the tag without changing Base.
}

func TestKoan(t *testing.T) {
	plain := struct{ Base }{Base{Name: "Ada"}}
	if got := plain.Label(); got != "Ada" {
		t.Errorf("promoted label=%q; want Ada", got)
	}
	for _, tc := range []struct{ name, tag, want string }{{"Ada", "admin", "Ada [admin]"}, {"Grace", "", "Grace"}, {"Go", "test", "Go [test]"}} {
		value := Tagged{Base: Base{Name: tc.name}, Tag: tc.tag}
		if got := value.Label(); got != tc.want {
			t.Errorf("overridden label=%q; want %q", got, tc.want)
		}
		if got := value.Base.Label(); got != tc.name {
			t.Errorf("base label=%q; want %q", got, tc.name)
		}
	}
}
