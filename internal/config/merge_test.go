package config

import (
	"flag"
	"testing"
)

func newMergeFS(t *testing.T, args ...string) *flag.FlagSet {
	t.Helper()
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	registerMergeFlags(fs)
	if err := fs.Parse(args); err != nil {
		t.Fatalf("parse: %v", err)
	}
	return fs
}

func TestMergeInputs_NilFlagSet(t *testing.T) {
	if MergeInputs(nil) {
		t.Fatal("expected false for nil FlagSet")
	}
}

func TestMergeInputs_DefaultConfig(t *testing.T) {
	fs := newMergeFS(t)
	if MergeInputs(fs) {
		t.Fatal("expected merge to be false by default")
	}
}

func TestMergeInputs_ExplicitFlag(t *testing.T) {
	fs := newMergeFS(t, "--merge")
	if !MergeInputs(fs) {
		t.Fatal("expected merge to be true when flag set")
	}
}

func TestMergeTimestampLayout_NilFlagSet(t *testing.T) {
	if got := MergeTimestampLayout(nil); got != "" {
		t.Fatalf("expected empty string, got %q", got)
	}
}

func TestMergeTimestampLayout_DefaultConfig(t *testing.T) {
	fs := newMergeFS(t)
	if got := MergeTimestampLayout(fs); got != "" {
		t.Fatalf("expected empty string by default, got %q", got)
	}
}

func TestMergeTimestampLayout_ExplicitValue(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	fs := newMergeFS(t, "--merge-layout="+layout)
	if got := MergeTimestampLayout(fs); got != layout {
		t.Fatalf("expected %q, got %q", layout, got)
	}
}

func TestMergeInputs_UnregisteredFlag(t *testing.T) {
	fs := flag.NewFlagSet("empty", flag.ContinueOnError)
	if MergeInputs(fs) {
		t.Fatal("expected false when flag not registered")
	}
}

func TestMergeTimestampLayout_UnregisteredFlag(t *testing.T) {
	fs := flag.NewFlagSet("empty", flag.ContinueOnError)
	if got := MergeTimestampLayout(fs); got != "" {
		t.Fatalf("expected empty string when flag not registered, got %q", got)
	}
}
