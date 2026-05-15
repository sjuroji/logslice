package config_test

import (
	"flag"
	"testing"

	"github.com/logslice/logslice/internal/config"
)

func newRateFS(args ...string) *flag.FlagSet {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	config.RateLimit(fs)
	_ = fs.Parse(args)
	return fs
}

func TestRateLimit_NilFlagSet(t *testing.T) {
	v := config.RateLimit(nil)
	if v != 0 {
		t.Fatalf("expected 0, got %d", v)
	}
}

func TestRateLimit_DefaultConfig(t *testing.T) {
	fs := newRateFS()
	v := config.RateLimit(fs)
	if v != 0 {
		t.Fatalf("expected default 0, got %d", v)
	}
}

func TestRateLimit_ExplicitValue(t *testing.T) {
	fs := newRateFS("--rate", "100")
	v := config.RateLimit(fs)
	if v != 100 {
		t.Fatalf("expected 100, got %d", v)
	}
}

func TestRateLimit_ShortFlag(t *testing.T) {
	fs := newRateFS("-r", "50")
	v := config.RateLimit(fs)
	if v != 50 {
		t.Fatalf("expected 50, got %d", v)
	}
}

func TestRateLimit_LargeValue(t *testing.T) {
	fs := newRateFS("--rate", "100000")
	v := config.RateLimit(fs)
	if v != 100000 {
		t.Fatalf("expected 100000, got %d", v)
	}
}
