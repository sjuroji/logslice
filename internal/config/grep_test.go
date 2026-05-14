package config

import (
	"flag"
	"testing"
)

func TestPatterns_NilConfig(t *testing.T) {
	if got := Patterns(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestPatterns_DefaultConfig(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	registerGrepFlags(fs)
	_ = fs.Parse([]string{})
	if got := Patterns(fs); got != nil {
		t.Fatalf("expected nil for empty pattern, got %v", got)
	}
}

func TestPatterns_SingleValue(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	registerGrepFlags(fs)
	_ = fs.Parse([]string{"--pattern", "ERROR"})
	got := Patterns(fs)
	if len(got) != 1 || got[0] != "ERROR" {
		t.Fatalf("expected [ERROR], got %v", got)
	}
}

func TestPatterns_MultipleValues(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	registerGrepFlags(fs)
	_ = fs.Parse([]string{"--pattern", "ERROR,WARN,INFO"})
	got := Patterns(fs)
	if len(got) != 3 {
		t.Fatalf("expected 3 patterns, got %d: %v", len(got), got)
	}
}

func TestExcludePatterns_NilConfig(t *testing.T) {
	if got := ExcludePatterns(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestExcludePatterns_DefaultConfig(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	registerGrepFlags(fs)
	_ = fs.Parse([]string{})
	if got := ExcludePatterns(fs); got != nil {
		t.Fatalf("expected nil for empty exclude, got %v", got)
	}
}

func TestExcludePatterns_SingleValue(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	registerGrepFlags(fs)
	_ = fs.Parse([]string{"--exclude", "DEBUG"})
	got := ExcludePatterns(fs)
	if len(got) != 1 || got[0] != "DEBUG" {
		t.Fatalf("expected [DEBUG], got %v", got)
	}
}

func TestExcludePatterns_MultipleValues(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	registerGrepFlags(fs)
	_ = fs.Parse([]string{"--exclude", "DEBUG,TRACE"})
	got := ExcludePatterns(fs)
	if len(got) != 2 {
		t.Fatalf("expected 2 patterns, got %d: %v", len(got), got)
	}
}
