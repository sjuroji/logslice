package grep_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/grep"
)

func TestNewMulti_EmptySlice(t *testing.T) {
	mu, err := grep.NewMulti(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if mu.Len() != 0 {
		t.Errorf("Len() = %d; want 0", mu.Len())
	}
}

func TestNewMulti_SkipsEmptyPatterns(t *testing.T) {
	mu, err := grep.NewMulti([]string{"", ""})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if mu.Len() != 0 {
		t.Errorf("Len() = %d; want 0", mu.Len())
	}
}

func TestNewMulti_InvalidPattern(t *testing.T) {
	_, err := grep.NewMulti([]string{"valid", "[bad"})
	if err == nil {
		t.Fatal("expected error for invalid pattern")
	}
}

func TestMulti_MatchAll(t *testing.T) {
	mu, err := grep.NewMulti([]string{"ERROR", "database"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !mu.Match("ERROR database connection refused") {
		t.Error("expected match when all patterns present")
	}
	if mu.Match("ERROR network timeout") {
		t.Error("expected no match when only one pattern present")
	}
	if mu.Match("INFO database ready") {
		t.Error("expected no match when only second pattern present")
	}
}

func TestMulti_EmptyMatchesAll(t *testing.T) {
	mu, err := grep.NewMulti([]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := []string{"anything", "", "ERROR foo"}
	for _, l := range lines {
		if !mu.Match(l) {
			t.Errorf("empty Multi should match %q", l)
		}
	}
}

func TestMulti_Len(t *testing.T) {
	mu, _ := grep.NewMulti([]string{"a", "b", "c"})
	if mu.Len() != 3 {
		t.Errorf("Len() = %d; want 3", mu.Len())
	}
}
