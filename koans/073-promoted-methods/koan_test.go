package main

// Promoted methods
// Predict the method results. Plain promotes Base.Label without declaring a method.
// Tagged declares its own Label, which takes precedence; Tagged.Base.Label still
// calls the embedded method. Change only the expected strings.

import "testing"

type Base struct{ Name string }

func (b Base) Label() string { return b.Name }

type Plain struct{ Base }
type Tagged struct {
	Base
	Tag string
}

func (t Tagged) Label() string { return t.Base.Label() + " [" + t.Tag + "]" }

func TestKoan(t *testing.T) {
	plain := Plain{Base: Base{Name: "Ada"}}
	tagged := Tagged{Base: Base{Name: "Grace"}, Tag: "admin"}

	wantPromoted := "" // TODO: predict plain.Label().
	if got := plain.Label(); got != wantPromoted {
		t.Errorf("promoted label=%q; want %q", got, wantPromoted)
	}
	wantExplicit := "" // TODO: predict plain.Base.Label().
	if got := plain.Base.Label(); got != wantExplicit {
		t.Errorf("explicit base label=%q; want %q", got, wantExplicit)
	}
	wantOverride := "" // TODO: predict tagged.Label().
	if got := tagged.Label(); got != wantOverride {
		t.Errorf("overridden label=%q; want %q", got, wantOverride)
	}
	wantBase := "" // TODO: predict tagged.Base.Label().
	if got := tagged.Base.Label(); got != wantBase {
		t.Errorf("embedded label=%q; want %q", got, wantBase)
	}
}
