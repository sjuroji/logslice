package config

import (
	"flag"
	"testing"
	"time"
)

func TestTimeout_NilFlagSet(t *testing.T) {
	if got := Timeout(nil); got != 0 {
		t.Fatalf("expected 0, got %v", got)
	}
}

func TestTimeout_DefaultConfig(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	registerContextFlags(fs)
	_ = fs.Parse([]string{})
	if got := Timeout(fs); got != 0 {
		t.Fatalf("expected default 0, got %v", got)
	}
}

func TestTimeout_ExplicitValue(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	registerContextFlags(fs)
	_ = fs.Parse([]string{"-timeout", "45s"})
	want := 45 * time.Second
	if got := Timeout(fs); got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestTimeout_MinuteValue(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	registerContextFlags(fs)
	_ = fs.Parse([]string{"-timeout", "2m"})
	want := 2 * time.Minute
	if got := Timeout(fs); got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestTimeout_UnregisteredFlag(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	// intentionally do NOT register context flags
	_ = fs.Parse([]string{})
	if got := Timeout(fs); got != 0 {
		t.Fatalf("expected 0 for unregistered flag, got %v", got)
	}
}
