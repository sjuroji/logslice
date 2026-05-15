package truncate_test

import (
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/truncate"
)

func TestNew_ZeroReturnsNil(t *testing.T) {
	if truncate.New(0, "...") != nil {
		t.Fatal("expected nil for maxLen=0")
	}
}

func TestNew_NegativeReturnsNil(t *testing.T) {
	if truncate.New(-1, "...") != nil {
		t.Fatal("expected nil for maxLen=-1")
	}
}

func TestNew_PositiveReturnsNonNil(t *testing.T) {
	if truncate.New(80, "...") == nil {
		t.Fatal("expected non-nil truncator")
	}
}

func TestApply_NilPassesThrough(t *testing.T) {
	var tr *truncate.Truncator
	line := "hello world"
	if got := tr.Apply(line); got != line {
		t.Fatalf("expected %q, got %q", line, got)
	}
}

func TestApply_ShortLineUnchanged(t *testing.T) {
	tr := truncate.New(20, "...")
	line := "short"
	if got := tr.Apply(line); got != line {
		t.Fatalf("expected %q, got %q", line, got)
	}
}

func TestApply_ExactLengthUnchanged(t *testing.T) {
	tr := truncate.New(5, "...")
	line := "hello"
	if got := tr.Apply(line); got != line {
		t.Fatalf("expected %q, got %q", line, got)
	}
}

func TestApply_LongLineClipped(t *testing.T) {
	tr := truncate.New(10, "...")
	line := "hello world this is long"
	got := tr.Apply(line)
	if len(got) > 10 {
		t.Fatalf("result length %d exceeds maxLen 10: %q", len(got), got)
	}
	if !strings.HasSuffix(got, "...") {
		t.Fatalf("expected suffix '...', got %q", got)
	}
}

func TestApply_NoSuffix(t *testing.T) {
	tr := truncate.New(5, "")
	line := "hello world"
	got := tr.Apply(line)
	if got != "hello" {
		t.Fatalf("expected %q, got %q", "hello", got)
	}
}

func TestApply_UTF8RuneBoundary(t *testing.T) {
	// Each Japanese character is 3 bytes; maxLen=5 must not split a rune.
	tr := truncate.New(5, "")
	line := "日本語"
	got := tr.Apply(line)
	if !strings.HasPrefix(line, got) {
		t.Fatalf("result %q is not a prefix of original", got)
	}
}

func TestMaxLen_NilReturnsZero(t *testing.T) {
	var tr *truncate.Truncator
	if tr.MaxLen() != 0 {
		t.Fatal("expected 0 for nil truncator")
	}
}

func TestMaxLen_ReturnsConfigured(t *testing.T) {
	tr := truncate.New(120, "...")
	if tr.MaxLen() != 120 {
		t.Fatalf("expected 120, got %d", tr.MaxLen())
	}
}
