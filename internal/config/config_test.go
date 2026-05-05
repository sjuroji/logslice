package config_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/config"
)

func parse(t *testing.T, args ...string) *config.Config {
	t.Helper()
	var out, errOut bytes.Buffer
	cfg, err := config.Parse(args, &out, &errOut)
	if err != nil {
		t.Fatalf("Parse(%v) unexpected error: %v", args, err)
	}
	return cfg
}

func TestParse_Defaults(t *testing.T) {
	cfg := parse(t, "file.log")
	if cfg.Pattern != "" {
		t.Errorf("expected empty pattern, got %q", cfg.Pattern)
	}
	if cfg.MaxLines != 0 {
		t.Errorf("expected MaxLines 0, got %d", cfg.MaxLines)
	}
	if cfg.ShowLineNumbers {
		t.Error("expected ShowLineNumbers false")
	}
	if len(cfg.Files) != 1 || cfg.Files[0] != "file.log" {
		t.Errorf("expected files [file.log], got %v", cfg.Files)
	}
}

func TestParse_Pattern(t *testing.T) {
	cfg := parse(t, "--pattern", "ERROR", "app.log")
	if cfg.Pattern != "ERROR" {
		t.Errorf("expected pattern ERROR, got %q", cfg.Pattern)
	}
}

func TestParse_MaxLines(t *testing.T) {
	cfg := parse(t, "--max-lines", "50")
	if cfg.MaxLines != 50 {
		t.Errorf("expected MaxLines 50, got %d", cfg.MaxLines)
	}
}

func TestParse_LineNumbers(t *testing.T) {
	cfg := parse(t, "--line-numbers")
	if !cfg.ShowLineNumbers {
		t.Error("expected ShowLineNumbers true")
	}
}

func TestParse_SinceUntil_RFC3339(t *testing.T) {
	cfg := parse(t,
		"--since", "2024-01-15T08:00:00Z",
		"--until", "2024-01-15T18:00:00Z",
	)
	wantSince := time.Date(2024, 1, 15, 8, 0, 0, 0, time.UTC)
	if !cfg.Since.Equal(wantSince) {
		t.Errorf("Since: got %v, want %v", cfg.Since, wantSince)
	}
	wantUntil := time.Date(2024, 1, 15, 18, 0, 0, 0, time.UTC)
	if !cfg.Until.Equal(wantUntil) {
		t.Errorf("Until: got %v, want %v", cfg.Until, wantUntil)
	}
}

func TestParse_Since_CustomFormat(t *testing.T) {
	cfg := parse(t,
		"--time-format", "2006-01-02",
		"--since", "2024-03-01",
	)
	want := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	if !cfg.Since.Equal(want) {
		t.Errorf("Since: got %v, want %v", cfg.Since, want)
	}
}

func TestParse_InvalidSince(t *testing.T) {
	var out, errOut bytes.Buffer
	_, err := config.Parse([]string{"--since", "not-a-date"}, &out, &errOut)
	if err == nil {
		t.Error("expected error for invalid --since, got nil")
	}
}

func TestParse_Prefix(t *testing.T) {
	cfg := parse(t, "--prefix", "[app]", "file.log")
	if cfg.Prefix != "[app]" {
		t.Errorf("expected prefix [app], got %q", cfg.Prefix)
	}
}
