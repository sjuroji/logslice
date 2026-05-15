package merge_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/merge"
	"github.com/yourorg/logslice/internal/timeparser"
)

func newParser(t *testing.T) *timeparser.Parser {
	t.Helper()
	p, err := timeparser.New(time.UTC)
	if err != nil {
		t.Fatalf("timeparser.New: %v", err)
	}
	return p
}

func TestMerge_SingleReader(t *testing.T) {
	input := "2024-01-01T00:00:01Z line-a\n2024-01-01T00:00:03Z line-c\n"
	m := merge.New([]interface{ Read([]byte) (int, error) }{strings.NewReader(input)}, newParser(t))
	var buf bytes.Buffer
	if err := m.Merge(&buf); err != nil {
		t.Fatalf("Merge: %v", err)
	}
	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	if len(lines) != 2 {
		t.Fatalf("want 2 lines, got %d", len(lines))
	}
	if !strings.Contains(lines[0], "line-a") {
		t.Errorf("line 0: want line-a, got %q", lines[0])
	}
}

func TestMerge_TwoReadersInterleaved(t *testing.T) {
	a := strings.NewReader("2024-01-01T00:00:01Z alpha\n2024-01-01T00:00:05Z gamma\n")
	b := strings.NewReader("2024-01-01T00:00:02Z beta\n2024-01-01T00:00:07Z delta\n")
	m := merge.New([]interface{ Read([]byte) (int, error) }{a, b}, newParser(t))
	var buf bytes.Buffer
	if err := m.Merge(&buf); err != nil {
		t.Fatalf("Merge: %v", err)
	}
	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	want := []string{"alpha", "beta", "gamma", "delta"}
	if len(lines) != len(want) {
		t.Fatalf("want %d lines, got %d", len(want), len(lines))
	}
	for i, w := range want {
		if !strings.Contains(lines[i], w) {
			t.Errorf("line %d: want %q, got %q", i, w, lines[i])
		}
	}
}

func TestMerge_UnparsedLinesAppendedLast(t *testing.T) {
	a := strings.NewReader("2024-01-01T00:00:01Z first\nno-timestamp-here\n")
	m := merge.New([]interface{ Read([]byte) (int, error) }{a}, newParser(t))
	var buf bytes.Buffer
	if err := m.Merge(&buf); err != nil {
		t.Fatalf("Merge: %v", err)
	}
	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	if len(lines) != 2 {
		t.Fatalf("want 2 lines, got %d", len(lines))
	}
	if !strings.Contains(lines[0], "first") {
		t.Errorf("expected timestamped line first, got %q", lines[0])
	}
	if !strings.Contains(lines[1], "no-timestamp-here") {
		t.Errorf("expected unparsed line last, got %q", lines[1])
	}
}

func TestMerge_EmptyReaders(t *testing.T) {
	a := strings.NewReader("")
	b := strings.NewReader("")
	m := merge.New([]interface{ Read([]byte) (int, error) }{a, b}, newParser(t))
	var buf bytes.Buffer
	if err := m.Merge(&buf); err != nil {
		t.Fatalf("Merge: %v", err)
	}
	if buf.Len() != 0 {
		t.Errorf("expected empty output, got %q", buf.String())
	}
}

func TestMerge_NoReaders(t *testing.T) {
	m := merge.New(nil, newParser(t))
	var buf bytes.Buffer
	if err := m.Merge(&buf); err != nil {
		t.Fatalf("Merge: %v", err)
	}
	if buf.Len() != 0 {
		t.Errorf("expected empty output, got %q", buf.String())
	}
}
