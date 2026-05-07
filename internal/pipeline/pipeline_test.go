package pipeline_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/example/logslice/internal/config"
	"github.com/example/logslice/internal/pipeline"
	"github.com/example/logslice/internal/stats"
)

func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	p := filepath.Join(dir, "test.log")
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatalf("writeTempFile: %v", err)
	}
	return p
}

func defaultCfg() *config.Config {
	return &config.Config{
		MaxLines: 0,
		LineNumbers: false,
	}
}

func TestRun_AllLines(t *testing.T) {
	path := writeTempFile(t, "alpha\nbeta\ngamma\n")
	var buf bytes.Buffer
	s := stats.New()
	pl := pipeline.New(defaultCfg(), &buf, s)

	if err := pl.Run(path); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(got) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(got))
	}
}

func TestRun_PatternFilter(t *testing.T) {
	path := writeTempFile(t, "alpha\nbeta\ngamma\n")
	var buf bytes.Buffer
	s := stats.New()
	cfg := defaultCfg()
	cfg.Pattern = "beta"
	pl := pipeline.New(cfg, &buf, s)

	if err := pl.Run(path); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(buf.String(), "beta") {
		t.Error("expected 'beta' in output")
	}
	if strings.Contains(buf.String(), "alpha") {
		t.Error("unexpected 'alpha' in output")
	}
}

func TestRun_FileNotFound(t *testing.T) {
	var buf bytes.Buffer
	s := stats.New()
	pl := pipeline.New(defaultCfg(), &buf, s)

	if err := pl.Run("/nonexistent/path/file.log"); err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestRun_StatsUpdated(t *testing.T) {
	path := writeTempFile(t, "line1\nline2\nline3\n")
	var buf bytes.Buffer
	s := stats.New()
	pl := pipeline.New(defaultCfg(), &buf, s)

	if err := pl.Run(path); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_ = time.Now() // ensure stats.Finish can be called later
	snap := s.Snapshot()
	if snap.FilesProcessed != 1 {
		t.Errorf("expected 1 file processed, got %d", snap.FilesProcessed)
	}
	if snap.LinesRead != 3 {
		t.Errorf("expected 3 lines read, got %d", snap.LinesRead)
	}
}
