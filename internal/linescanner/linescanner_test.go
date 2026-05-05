package linescanner_test

import (
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/linescanner"
)

func TestScan_NoFilter(t *testing.T) {
	input := "line one\nline two\nline three\n"
	s := linescanner.New(strings.NewReader(input), linescanner.Options{})

	lines, err := s.Scan()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
}

func TestScan_PatternFilter(t *testing.T) {
	input := "2024-01-01 ERROR something failed\n2024-01-01 INFO all good\n2024-01-01 ERROR another failure\n"
	s := linescanner.New(strings.NewReader(input), linescanner.Options{
		Pattern: "ERROR",
	})

	lines, err := s.Scan()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
	for _, l := range lines {
		if !strings.Contains(l, "ERROR") {
			t.Errorf("expected line to contain ERROR, got: %s", l)
		}
	}
}

func TestScan_MaxLines(t *testing.T) {
	input := "a\nb\nc\nd\ne\n"
	s := linescanner.New(strings.NewReader(input), linescanner.Options{
		MaxLines: 3,
	})

	lines, err := s.Scan()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
}

func TestScan_EmptyInput(t *testing.T) {
	s := linescanner.New(strings.NewReader(""), linescanner.Options{})

	lines, err := s.Scan()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 0 {
		t.Fatalf("expected 0 lines, got %d", len(lines))
	}
}

func TestScan_PatternAndMaxLines(t *testing.T) {
	input := "ERROR one\nINFO skip\nERROR two\nERROR three\n"
	s := linescanner.New(strings.NewReader(input), linescanner.Options{
		Pattern:  "ERROR",
		MaxLines: 2,
	})

	lines, err := s.Scan()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
}
