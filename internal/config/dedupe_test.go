package config_test

import (
	"flag"
	"testing"

	"github.com/logslice/logslice/internal/config"
)

func newDedupeFS(args ...string) *flag.FlagSet {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.Bool("dedupe", false, "")
	fs.Bool("D", false, "")
	_ = fs.Parse(args)
	return fs
}

func TestDedupeLines_NilFlagSet(t *testing.T) {
	if config.DedupeLines(nil) {
		t.Fatal("expected false for nil FlagSet")
	}
}

func TestDedupeLines_DefaultConfig(t *testing.T) {
	fs := newDedupeFS()
	if config.DedupeLines(fs) {
		t.Fatal("expected false when flag not set")
	}
}

func TestDedupeLines_ExplicitFlag(t *testing.T) {
	fs := newDedupeFS("--dedupe")
	if !config.DedupeLines(fs) {
		t.Fatal("expected true when --dedupe is set")
	}
}

func TestDedupeLines_ShortFlag(t *testing.T) {
	fs := newDedupeFS("-D")
	// Short flag is a separate bool; check it is registered.
	f := fs.Lookup("D")
	if f == nil {
		t.Fatal("expected -D flag to be registered")
	}
	if f.Value.String() != "true" {
		t.Fatalf("expected -D to be true, got %s", f.Value.String())
	}
}

func TestDedupeLines_UnregisteredFlag(t *testing.T) {
	fs := flag.NewFlagSet("empty", flag.ContinueOnError)
	_ = fs.Parse(nil)
	if config.DedupeLines(fs) {
		t.Fatal("expected false when flag is not registered")
	}
}
