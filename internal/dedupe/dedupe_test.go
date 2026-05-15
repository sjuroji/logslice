package dedupe_test

import (
	"testing"

	"github.com/logslice/logslice/internal/dedupe"
)

func TestNew_InitialisesEmpty(t *testing.T) {
	f := dedupe.New()
	if f.Unique() != 0 {
		t.Fatalf("expected 0 unique lines, got %d", f.Unique())
	}
}

func TestAccept_FirstOccurrenceAccepted(t *testing.T) {
	f := dedupe.New()
	if !f.Accept("hello world") {
		t.Fatal("expected first occurrence to be accepted")
	}
}

func TestAccept_DuplicateRejected(t *testing.T) {
	f := dedupe.New()
	f.Accept("hello world")
	if f.Accept("hello world") {
		t.Fatal("expected duplicate to be rejected")
	}
}

func TestAccept_DifferentLinesAccepted(t *testing.T) {
	f := dedupe.New()
	if !f.Accept("line one") {
		t.Fatal("expected line one to be accepted")
	}
	if !f.Accept("line two") {
		t.Fatal("expected line two to be accepted")
	}
}

func TestCount_TracksSuppressions(t *testing.T) {
	f := dedupe.New()
	f.Accept("dup")
	f.Accept("dup")
	f.Accept("dup")
	if got := f.Count("dup"); got != 2 {
		t.Fatalf("expected 2 suppressions, got %d", got)
	}
}

func TestCount_UnseenLineReturnsZero(t *testing.T) {
	f := dedupe.New()
	if got := f.Count("never seen"); got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}

func TestUnique_CountsDistinctLines(t *testing.T) {
	f := dedupe.New()
	f.Accept("a")
	f.Accept("b")
	f.Accept("a")
	f.Accept("c")
	if got := f.Unique(); got != 3 {
		t.Fatalf("expected 3 unique lines, got %d", got)
	}
}

func TestReset_ClearsState(t *testing.T) {
	f := dedupe.New()
	f.Accept("line")
	f.Reset()
	if f.Unique() != 0 {
		t.Fatalf("expected 0 after reset, got %d", f.Unique())
	}
	if !f.Accept("line") {
		t.Fatal("expected line to be accepted after reset")
	}
}
