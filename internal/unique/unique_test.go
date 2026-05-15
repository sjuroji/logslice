package unique

import (
	"testing"
)

func TestNew_ZeroWindowReturnsFilter(t *testing.T) {
	f, err := New(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f == nil {
		t.Fatal("expected non-nil filter")
	}
}

func TestNew_NegativeWindowReturnsError(t *testing.T) {
	_, err := New(-1)
	if err == nil {
		t.Fatal("expected error for negative window")
	}
}

func TestAccept_NilAlwaysTrue(t *testing.T) {
	var f *Filter
	if !f.Accept("anything") {
		t.Error("nil filter should always accept")
	}
}

func TestAccept_FirstOccurrenceAccepted(t *testing.T) {
	f, _ := New(0)
	if !f.Accept("hello") {
		t.Error("first occurrence should be accepted")
	}
}

func TestAccept_DuplicateRejected(t *testing.T) {
	f, _ := New(0)
	f.Accept("hello")
	if f.Accept("hello") {
		t.Error("duplicate should be rejected")
	}
}

func TestAccept_DifferentLinesAccepted(t *testing.T) {
	f, _ := New(0)
	f.Accept("line1")
	if !f.Accept("line2") {
		t.Error("different line should be accepted")
	}
}

func TestSuppressed_CountsDuplicates(t *testing.T) {
	f, _ := New(0)
	f.Accept("a")
	f.Accept("a")
	f.Accept("a")
	if f.Suppressed() != 2 {
		t.Errorf("expected 2 suppressed, got %d", f.Suppressed())
	}
}

func TestAccept_WindowEvictsOldEntries(t *testing.T) {
	f, _ := New(2)
	f.Accept("a") // seen: [a]
	f.Accept("b") // seen: [a, b]
	f.Accept("c") // seen: [b, c] — "a" evicted
	if !f.Accept("a") {
		t.Error("'a' should be accepted after eviction from window")
	}
}

func TestAccept_WindowDuplicateWithinWindow(t *testing.T) {
	f, _ := New(3)
	f.Accept("x")
	if f.Accept("x") {
		t.Error("'x' should be rejected while still in window")
	}
}

func TestReset_ClearsState(t *testing.T) {
	f, _ := New(0)
	f.Accept("hello")
	f.Accept("hello")
	f.Reset()
	if !f.Accept("hello") {
		t.Error("after reset, 'hello' should be accepted again")
	}
	if f.Suppressed() != 0 {
		t.Errorf("expected 0 suppressed after reset, got %d", f.Suppressed())
	}
}
