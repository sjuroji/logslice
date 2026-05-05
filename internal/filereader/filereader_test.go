package filereader_test

import (
	"os"
	"strings"
	"testing"

	"github.com/user/logslice/internal/filereader"
)

func TestNewFromReader_ReadLines(t *testing.T) {
	input := "line one\nline two\nline three\n"
	r := filereader.NewFromReader(strings.NewReader(input))

	var lines []string
	for r.Scan() {
		lines = append(lines, r.Text())
	}
	if err := r.Err(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
	if lines[1] != "line two" {
		t.Errorf("expected 'line two', got %q", lines[1])
	}
}

func TestNewFromReader_EmptyInput(t *testing.T) {
	r := filereader.NewFromReader(strings.NewReader(""))
	if r.Scan() {
		t.Fatal("expected no lines from empty input")
	}
	if err := r.Err(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNew_FileNotFound(t *testing.T) {
	_, err := filereader.New("/nonexistent/path/to/file.log")
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}

func TestNew_RealFile(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "logslice-*.log")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	_, _ = f.WriteString("alpha\nbeta\ngamma\n")
	f.Close()

	r, err := filereader.New(f.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer r.Close()

	if r.Path() != f.Name() {
		t.Errorf("expected path %q, got %q", f.Name(), r.Path())
	}

	var count int
	for r.Scan() {
		count++
	}
	if count != 3 {
		t.Errorf("expected 3 lines, got %d", count)
	}
}

func TestClose_NoopOnReaderBacked(t *testing.T) {
	r := filereader.NewFromReader(strings.NewReader("hello\n"))
	if err := r.Close(); err != nil {
		t.Errorf("expected nil error on Close, got %v", err)
	}
}
