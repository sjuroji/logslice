package reporter_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/reporter"
	"github.com/yourorg/logslice/internal/stats"
)

func buildStats(files, read, matched int) *stats.Stats {
	s := stats.New()
	for i := 0; i < files; i++ {
		s.AddFile()
	}
	s.AddRead(read)
	s.AddMatched(matched)
	s.Finish()
	return s
}

func TestPrint_ContainsCounts(t *testing.T) {
	var buf strings.Builder
	r := reporter.New(&buf, false)
	s := buildStats(2, 100, 42)

	if err := r.Print(s); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	for _, want := range []string{"files: 2", "read: 100", "matched: 42"} {
		if !strings.Contains(out, want) {
			t.Errorf("output %q missing %q", out, want)
		}
	}
}

func TestPrint_VerboseIncludesTimes(t *testing.T) {
	var buf strings.Builder
	r := reporter.New(&buf, true)
	s := buildStats(1, 10, 5)

	if err := r.Print(s); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	for _, want := range []string{"started:", "finished:"} {
		if !strings.Contains(out, want) {
			t.Errorf("verbose output %q missing %q", out, want)
		}
	}
}

func TestPrint_NonVerboseOmitsTimes(t *testing.T) {
	var buf strings.Builder
	r := reporter.New(&buf, false)
	s := buildStats(1, 10, 5)

	_ = r.Print(s)
	out := buf.String()

	if strings.Contains(out, "started:") {
		t.Errorf("non-verbose output should not contain timestamps, got: %q", out)
	}
}

func TestPrint_WriteError(t *testing.T) {
	r := reporter.New(&errWriter{}, false)
	s := buildStats(0, 0, 0)

	if err := r.Print(s); err == nil {
		t.Fatal("expected error from failing writer")
	}
}

type errWriter struct{}

func (e *errWriter) Write(_ []byte) (int, error) {
	return 0, fmt.Errorf("write error")
}
