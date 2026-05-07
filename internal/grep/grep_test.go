package grep_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/grep"
)

func TestNew_EmptyPattern(t *testing.T) {
	m, err := grep.New("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if m == nil {
		t.Fatal("expected non-nil Matcher")
	}
}

func TestNew_InvalidPattern(t *testing.T) {
	_, err := grep.New("[invalid")
	if err == nil {
		t.Fatal("expected error for invalid pattern")
	}
}

func TestMatch_EmptyPatternMatchesAll(t *testing.T) {
	m, _ := grep.New("")
	lines := []string{"hello world", "", "2024-01-01 ERROR foo"}
	for _, l := range lines {
		if !m.Match(l) {
			t.Errorf("empty pattern should match %q", l)
		}
	}
}

func TestMatch_PatternMatches(t *testing.T) {
	m, err := grep.New(`ERROR`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !m.Match("2024-01-01 ERROR something failed") {
		t.Error("expected match")
	}
	if m.Match("2024-01-01 INFO all good") {
		t.Error("expected no match")
	}
}

func TestMatch_RegexPattern(t *testing.T) {
	m, err := grep.New(`^\d{4}-\d{2}-\d{2}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !m.Match("2024-06-15 INFO started") {
		t.Error("expected match for date-prefixed line")
	}
	if m.Match("INFO no date here") {
		t.Error("expected no match for non-date-prefixed line")
	}
}

func TestPattern_ReturnsOriginal(t *testing.T) {
	pattern := `ERROR|WARN`
	m, _ := grep.New(pattern)
	if m.Pattern() != pattern {
		t.Errorf("Pattern() = %q; want %q", m.Pattern(), pattern)
	}
}

func TestPattern_EmptyWhenNoPattern(t *testing.T) {
	m, _ := grep.New("")
	if m.Pattern() != "" {
		t.Errorf("Pattern() = %q; want empty string", m.Pattern())
	}
}
