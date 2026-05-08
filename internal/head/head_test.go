package head_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/head"
)

func TestNew_ZeroReturnsNil(t *testing.T) {
	if head.New(0) != nil {
		t.Fatal("expected nil for n=0")
	}
}

func TestNew_NegativeReturnsNil(t *testing.T) {
	if head.New(-5) != nil {
		t.Fatal("expected nil for n<0")
	}
}

func TestNew_PositiveReturnsReader(t *testing.T) {
	if head.New(3) == nil {
		t.Fatal("expected non-nil for n=3")
	}
}

func TestAccept_NilAlwaysTrue(t *testing.T) {
	var r *head.Reader
	for i := 0; i < 100; i++ {
		if !r.Accept() {
			t.Fatalf("nil Reader.Accept() returned false on call %d", i)
		}
	}
}

func TestAccept_LimitsToN(t *testing.T) {
	r := head.New(3)
	for i := 0; i < 3; i++ {
		if !r.Accept() {
			t.Fatalf("expected true on call %d", i)
		}
	}
	if r.Accept() {
		t.Fatal("expected false after quota exhausted")
	}
}

func TestDone_NilAlwaysFalse(t *testing.T) {
	var r *head.Reader
	if r.Done() {
		t.Fatal("nil Reader.Done() should return false")
	}
}

func TestDone_TrueAfterQuota(t *testing.T) {
	r := head.New(2)
	r.Accept()
	r.Accept()
	if !r.Done() {
		t.Fatal("expected Done() == true after quota")
	}
}

func TestReadAll_FewerLinesThanN(t *testing.T) {
	lines := []string{"a", "b"}
	got := head.ReadAll(lines, 10)
	if len(got) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(got))
	}
}

func TestReadAll_MoreLinesThanN(t *testing.T) {
	lines := []string{"a", "b", "c", "d", "e"}
	got := head.ReadAll(lines, 3)
	if len(got) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(got))
	}
	for i, want := range []string{"a", "b", "c"} {
		if got[i] != want {
			t.Errorf("line %d: got %q want %q", i, got[i], want)
		}
	}
}

func TestReadAll_ExactlyN(t *testing.T) {
	lines := []string{"x", "y", "z"}
	got := head.ReadAll(lines, 3)
	if len(got) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(got))
	}
}

func TestReadAll_ZeroN_ReturnsAll(t *testing.T) {
	lines := []string{"a", "b", "c"}
	got := head.ReadAll(lines, 0)
	if len(got) != 3 {
		t.Fatalf("expected all 3 lines for n=0, got %d", len(got))
	}
}
