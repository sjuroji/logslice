package config

import (
	"flag"
	"testing"
)

func newContextFS(t *testing.T) *flag.FlagSet {
	t.Helper()
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	registerContextLinesFlags(fs)
	return fs
}

func TestBeforeLines_NilFlagSet(t *testing.T) {
	if got := BeforeLines(nil); got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}

func TestAfterLines_NilFlagSet(t *testing.T) {
	if got := AfterLines(nil); got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}

func TestBeforeLines_Default(t *testing.T) {
	fs := newContextFS(t)
	if err := fs.Parse(nil); err != nil {
		t.Fatal(err)
	}
	if got := BeforeLines(fs); got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}

func TestBeforeLines_ExplicitValue(t *testing.T) {
	fs := newContextFS(t)
	if err := fs.Parse([]string{"-before=3"}); err != nil {
		t.Fatal(err)
	}
	if got := BeforeLines(fs); got != 3 {
		t.Fatalf("expected 3, got %d", got)
	}
}

func TestAfterLines_ExplicitValue(t *testing.T) {
	fs := newContextFS(t)
	if err := fs.Parse([]string{"-after=5"}); err != nil {
		t.Fatal(err)
	}
	if got := AfterLines(fs); got != 5 {
		t.Fatalf("expected 5, got %d", got)
	}
}

func TestContextFlag_OverridesBefore(t *testing.T) {
	fs := newContextFS(t)
	if err := fs.Parse([]string{"-before=1", "-context=4"}); err != nil {
		t.Fatal(err)
	}
	if got := BeforeLines(fs); got != 4 {
		t.Fatalf("expected context(4) to override before(1), got %d", got)
	}
}

func TestContextFlag_OverridesAfter(t *testing.T) {
	fs := newContextFS(t)
	if err := fs.Parse([]string{"-after=2", "-context=6"}); err != nil {
		t.Fatal(err)
	}
	if got := AfterLines(fs); got != 6 {
		t.Fatalf("expected context(6) to override after(2), got %d", got)
	}
}

func TestContextFlag_ZeroDoesNotOverride(t *testing.T) {
	fs := newContextFS(t)
	if err := fs.Parse([]string{"-before=2", "-after=3"}); err != nil {
		t.Fatal(err)
	}
	if got := BeforeLines(fs); got != 2 {
		t.Fatalf("expected 2, got %d", got)
	}
	if got := AfterLines(fs); got != 3 {
		t.Fatalf("expected 3, got %d", got)
	}
}
