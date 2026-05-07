package progress_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/progress"
)

func TestFileStart_VerbosePrintsLine(t *testing.T) {
	var buf bytes.Buffer
	r := progress.New(&buf, 3, true)
	r.FileStart("foo.log")
	if !strings.Contains(buf.String(), "foo.log") {
		t.Errorf("expected path in output, got: %q", buf.String())
	}
	if !strings.Contains(buf.String(), "1/3") {
		t.Errorf("expected counter 1/3 in output, got: %q", buf.String())
	}
}

func TestFileStart_NonVerboseSilent(t *testing.T) {
	var buf bytes.Buffer
	r := progress.New(&buf, 1, false)
	r.FileStart("foo.log")
	if buf.Len() != 0 {
		t.Errorf("expected no output in non-verbose mode, got: %q", buf.String())
	}
}

func TestFileDone_VerbosePrintsCounts(t *testing.T) {
	var buf bytes.Buffer
	r := progress.New(&buf, 2, true)
	r.FileStart("bar.log")
	buf.Reset()
	r.FileDone("bar.log", 100, 42)
	out := buf.String()
	if !strings.Contains(out, "100") {
		t.Errorf("expected read count in output, got: %q", out)
	}
	if !strings.Contains(out, "42") {
		t.Errorf("expected matched count in output, got: %q", out)
	}
}

func TestFileDone_NonVerboseSilent(t *testing.T) {
	var buf bytes.Buffer
	r := progress.New(&buf, 1, false)
	r.FileStart("bar.log")
	buf.Reset()
	r.FileDone("bar.log", 50, 10)
	if buf.Len() != 0 {
		t.Errorf("expected no output in non-verbose mode, got: %q", buf.String())
	}
}

func TestError_AlwaysPrints(t *testing.T) {
	var buf bytes.Buffer
	r := progress.New(&buf, 1, false)
	r.Error("missing.log", errors.New("file not found"))
	out := buf.String()
	if !strings.Contains(out, "missing.log") {
		t.Errorf("expected path in error output, got: %q", out)
	}
	if !strings.Contains(out, "file not found") {
		t.Errorf("expected error message in output, got: %q", out)
	}
}

func TestCounter_IncrementsAcrossFiles(t *testing.T) {
	var buf bytes.Buffer
	r := progress.New(&buf, 3, true)
	for _, f := range []string{"a.log", "b.log", "c.log"} {
		r.FileStart(f)
	}
	if !strings.Contains(buf.String(), "3/3") {
		t.Errorf("expected final counter 3/3, got: %q", buf.String())
	}
}
