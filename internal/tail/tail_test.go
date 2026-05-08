package tail_test

import (
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/tail"
)

func TestNew_ZeroReturnsNil(t *testing.T) {
	if tail.New(0) != nil {
		t.Fatal("expected nil for n=0")
	}
}

func TestNew_NegativeReturnsNil(t *testing.T) {
	if tail.New(-5) != nil {
		t.Fatal("expected nil for n=-5")
	}
}

func TestReadAll_FewerLinesThanN(t *testing.T) {
	tr := tail.New(10)
	input := "line1\nline2\nline3"
	if err := tr.ReadAll(strings.NewReader(input)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := tr.Lines()
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
	if lines[0] != "line1" || lines[2] != "line3" {
		t.Errorf("unexpected lines: %v", lines)
	}
}

func TestReadAll_MoreLinesThanN(t *testing.T) {
	tr := tail.New(3)
	input := "a\nb\nc\nd\ne"
	if err := tr.ReadAll(strings.NewReader(input)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := tr.Lines()
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
	if lines[0] != "c" || lines[1] != "d" || lines[2] != "e" {
		t.Errorf("expected last 3 lines [c d e], got %v", lines)
	}
}

func TestReadAll_ExactlyN(t *testing.T) {
	tr := tail.New(3)
	input := "x\ny\nz"
	_ = tr.ReadAll(strings.NewReader(input))
	lines := tr.Lines()
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
}

func TestReadAll_EmptyInput(t *testing.T) {
	tr := tail.New(5)
	_ = tr.ReadAll(strings.NewReader(""))
	if tr.Len() != 0 {
		t.Fatalf("expected 0 lines, got %d", tr.Len())
	}
}

func TestLen_ReflectsRetainedLines(t *testing.T) {
	tr := tail.New(4)
	_ = tr.ReadAll(strings.NewReader("a\nb"))
	if tr.Len() != 2 {
		t.Fatalf("expected Len()=2, got %d", tr.Len())
	}
}
