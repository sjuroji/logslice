package highlight_test

import (
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/highlight"
)

func TestNew_EmptyPatternReturnsNil(t *testing.T) {
	h, err := highlight.New("", highlight.ColorYellow)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h != nil {
		t.Fatal("expected nil highlighter for empty pattern")
	}
}

func TestNew_InvalidPatternReturnsError(t *testing.T) {
	_, err := highlight.New("[", highlight.ColorYellow)
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestApply_NilHighlighterReturnsOriginal(t *testing.T) {
	var h *highlight.Highlighter
	line := "hello world"
	if got := h.Apply(line); got != line {
		t.Fatalf("expected %q, got %q", line, got)
	}
}

func TestApply_SingleMatch(t *testing.T) {
	h, _ := highlight.New("error", highlight.ColorRed)
	line := "an error occurred"
	got := h.Apply(line)
	if !strings.Contains(got, "\033[31m") {
		t.Error("expected red ANSI code in output")
	}
	if !strings.Contains(got, "error") {
		t.Error("expected original text preserved")
	}
	if !strings.Contains(got, "\033[0m") {
		t.Error("expected reset code in output")
	}
}

func TestApply_MultipleMatches(t *testing.T) {
	h, _ := highlight.New("o", highlight.ColorCyan)
	line := "foo bar boo"
	got := h.Apply(line)
	count := strings.Count(got, "\033[36m")
	if count != 3 {
		t.Fatalf("expected 3 highlights, got %d", count)
	}
}

func TestApply_NoMatch(t *testing.T) {
	h, _ := highlight.New("xyz", highlight.ColorYellow)
	line := "nothing to see here"
	got := h.Apply(line)
	if got != line {
		t.Fatalf("expected unchanged line, got %q", got)
	}
}

func TestApply_DefaultColorYellow(t *testing.T) {
	h, _ := highlight.New("warn", "unknown-color")
	line := "warn: disk low"
	got := h.Apply(line)
	if !strings.Contains(got, "\033[33m") {
		t.Error("expected yellow ANSI code as default")
	}
}

func TestApply_RegexGroup(t *testing.T) {
	h, _ := highlight.New(`\d+`, highlight.ColorCyan)
	line := "pid=1234 code=500"
	got := h.Apply(line)
	if strings.Count(got, "\033[36m") != 2 {
		t.Error("expected two numeric highlights")
	}
}
