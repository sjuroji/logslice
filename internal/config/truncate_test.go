package config_test

import (
	"flag"
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func newTruncateFS(args ...string) *flag.FlagSet {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	config.TruncateWidth(fs)
	config.TruncateSuffix(fs)
	_ = fs.Parse(args)
	return fs
}

func TestTruncateWidth_NilFlagSet(t *testing.T) {
	w := config.TruncateWidth(nil)
	if w != 0 {
		t.Fatalf("expected 0 for nil FlagSet, got %d", w)
	}
}

func TestTruncateWidth_DefaultConfig(t *testing.T) {
	fs := newTruncateFS()
	w := config.TruncateWidth(fs)
	if w != 0 {
		t.Fatalf("expected default 0, got %d", w)
	}
}

func TestTruncateWidth_ExplicitValue(t *testing.T) {
	fs := newTruncateFS("--truncate-width", "120")
	w := config.TruncateWidth(fs)
	if w != 120 {
		t.Fatalf("expected 120, got %d", w)
	}
}

func TestTruncateWidth_ShortFlag(t *testing.T) {
	fs := newTruncateFS("-W", "80")
	w := config.TruncateWidth(fs)
	if w != 80 {
		t.Fatalf("expected 80, got %d", w)
	}
}

func TestTruncateSuffix_NilFlagSet(t *testing.T) {
	s := config.TruncateSuffix(nil)
	if s != "" {
		t.Fatalf("expected empty string for nil FlagSet, got %q", s)
	}
}

func TestTruncateSuffix_DefaultConfig(t *testing.T) {
	fs := newTruncateFS()
	s := config.TruncateSuffix(fs)
	if s != "..." {
		t.Fatalf("expected default '...', got %q", s)
	}
}

func TestTruncateSuffix_ExplicitValue(t *testing.T) {
	fs := newTruncateFS("--truncate-suffix", "~")
	s := config.TruncateSuffix(fs)
	if s != "~" {
		t.Fatalf("expected '~', got %q", s)
	}
}
