package inverse_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/inverse"
)

// stubAccepter accepts only lines equal to "keep".
type stubAccepter struct{}

func (s *stubAccepter) Accept(line string) bool { return line == "keep" }

func TestNew_NilInnerReturnsNil(t *testing.T) {
	f := inverse.New(nil)
	if f != nil {
		t.Fatal("expected nil for nil inner")
	}
}

func TestNew_NonNilInnerReturnsFilter(t *testing.T) {
	f := inverse.New(&stubAccepter{})
	if f == nil {
		t.Fatal("expected non-nil Filter")
	}
}

func TestAccept_NilFilterAlwaysTrue(t *testing.T) {
	var f *inverse.Filter
	for _, line := range []string{"", "keep", "anything"} {
		if !f.Accept(line) {
			t.Errorf("nil Filter.Accept(%q) = false, want true", line)
		}
	}
}

func TestAccept_InvertsInnerDecision(t *testing.T) {
	f := inverse.New(&stubAccepter{})

	// inner accepts "keep", so inverse must reject it
	if f.Accept("keep") {
		t.Error("Accept(\"keep\") = true, want false")
	}

	// inner rejects everything else, so inverse must accept it
	for _, line := range []string{"", "drop", "other"} {
		if !f.Accept(line) {
			t.Errorf("Accept(%q) = false, want true", line)
		}
	}
}

func TestAccept_AllLinesRejectedByInner(t *testing.T) {
	// rejectAll never accepts anything
	type rejectAll struct{}
	accept := func(line string) bool { return false }
	_ = accept

	// Use an anonymous accepter via a small adapter.
	f := inverse.New(&alwaysReject{})
	for _, line := range []string{"", "foo", "bar"} {
		if !f.Accept(line) {
			t.Errorf("Accept(%q) = false, want true", line)
		}
	}
}

func TestAccept_AllLinesAcceptedByInner(t *testing.T) {
	f := inverse.New(&alwaysAccept{})
	for _, line := range []string{"", "foo", "bar"} {
		if f.Accept(line) {
			t.Errorf("Accept(%q) = true, want false", line)
		}
	}
}

type alwaysReject struct{}

func (a *alwaysReject) Accept(_ string) bool { return false }

type alwaysAccept struct{}

func (a *alwaysAccept) Accept(_ string) bool { return true }
