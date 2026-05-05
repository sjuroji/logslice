package output_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/output"
)

func TestWriteLine_Plain(t *testing.T) {
	var buf bytes.Buffer
	w := output.New(&buf, output.Options{})

	if err := w.WriteLine(1, "hello world"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	w.Flush()

	if got := buf.String(); got != "hello world\n" {
		t.Errorf("expected %q, got %q", "hello world\n", got)
	}
}

func TestWriteLine_WithLineNumbers(t *testing.T) {
	var buf bytes.Buffer
	w := output.New(&buf, output.Options{ShowLineNumbers: true})

	w.WriteLine(42, "some log entry")
	w.Flush()

	if got := buf.String(); got != "42: some log entry\n" {
		t.Errorf("expected %q, got %q", "42: some log entry\n", got)
	}
}

func TestWriteLine_WithPrefix(t *testing.T) {
	var buf bytes.Buffer
	w := output.New(&buf, output.Options{Prefix: "[INFO] "})

	w.WriteLine(1, "startup complete")
	w.Flush()

	if got := buf.String(); got != "[INFO] startup complete\n" {
		t.Errorf("expected %q, got %q", "[INFO] startup complete\n", got)
	}
}

func TestWriteLine_WithLineNumbersAndPrefix(t *testing.T) {
	var buf bytes.Buffer
	w := output.New(&buf, output.Options{ShowLineNumbers: true, Prefix: ">> "})

	w.WriteLine(7, "disk full")
	w.Flush()

	if got := buf.String(); got != ">> 7: disk full\n" {
		t.Errorf("expected %q, got %q", ">> 7: disk full\n", got)
	}
}

func TestWriteLine_MultipleLines(t *testing.T) {
	var buf bytes.Buffer
	w := output.New(&buf, output.Options{ShowLineNumbers: true})

	lines := []string{"alpha", "beta", "gamma"}
	for i, l := range lines {
		w.WriteLine(i+1, l)
	}
	w.Flush()

	result := buf.String()
	for i, l := range lines {
		expected := fmt.Sprintf("%d: %s\n", i+1, l)
		if !strings.Contains(result, expected) {
			t.Errorf("output missing line %q", expected)
		}
	}
}

func init() {
	// ensure fmt is used
	_ = fmt.Sprintf
}

import "fmt"
