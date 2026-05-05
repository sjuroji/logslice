package processor_test

import (
	"strings"
	"testing"
	"bytes"

	"github.com/yourorg/logslice/internal/logfilter"
	"github.com/yourorg/logslice/internal/output"
	"github.com/yourorg/logslice/internal/processor"
)

const sampleLog = `2024-01-10 10:00:01 INFO  service started
2024-01-10 10:00:02 DEBUG initialising cache
2024-01-10 10:00:03 ERROR connection refused
2024-01-10 10:00:04 INFO  request received
2024-01-10 10:00:05 ERROR timeout waiting for db
`

func newProcessor(t *testing.T, buf *bytes.Buffer, pattern string, maxLines int) *processor.Processor {
	t.Helper()
	f := logfilter.New(logfilter.Options{Pattern: pattern})
	w := output.New(buf, output.Options{})
	return processor.New(f, w, processor.Options{MaxLines: maxLines})
}

func TestProcess_NoFilter(t *testing.T) {
	var buf bytes.Buffer
	p := newProcessor(t, &buf, "", 0)

	n, err := p.Process(strings.NewReader(sampleLog))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 5 {
		t.Errorf("expected 5 lines written, got %d", n)
	}
}

func TestProcess_PatternFilter(t *testing.T) {
	var buf bytes.Buffer
	p := newProcessor(t, &buf, "ERROR", 0)

	n, err := p.Process(strings.NewReader(sampleLog))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 2 {
		t.Errorf("expected 2 ERROR lines, got %d", n)
	}
	if !strings.Contains(buf.String(), "connection refused") {
		t.Error("expected output to contain 'connection refused'")
	}
}

func TestProcess_MaxLines(t *testing.T) {
	var buf bytes.Buffer
	p := newProcessor(t, &buf, "", 2)

	n, err := p.Process(strings.NewReader(sampleLog))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 2 {
		t.Errorf("expected 2 lines written, got %d", n)
	}
}

func TestProcess_EmptyInput(t *testing.T) {
	var buf bytes.Buffer
	p := newProcessor(t, &buf, "", 0)

	n, err := p.Process(strings.NewReader(""))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 0 {
		t.Errorf("expected 0 lines, got %d", n)
	}
}

func TestProcess_PatternAndMaxLines(t *testing.T) {
	var buf bytes.Buffer
	p := newProcessor(t, &buf, "ERROR", 1)

	n, err := p.Process(strings.NewReader(sampleLog))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 1 {
		t.Errorf("expected 1 line, got %d", n)
	}
}
