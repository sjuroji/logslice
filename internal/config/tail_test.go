package config

import (
	"testing"
)

func TestTailLines_NilConfig(t *testing.T) {
	var c *Config
	if got := c.TailLines(); got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}

func TestTailLines_DefaultConfig(t *testing.T) {
	c := &Config{}
	if got := c.TailLines(); got != 0 {
		t.Fatalf("expected 0 for default config, got %d", got)
	}
}

func TestTailLines_ExplicitValue(t *testing.T) {
	c := &Config{Tail: 20}
	if got := c.TailLines(); got != 20 {
		t.Fatalf("expected 20, got %d", got)
	}
}

func TestTailLines_NegativeValue(t *testing.T) {
	c := &Config{Tail: -5}
	if got := c.TailLines(); got != 0 {
		t.Fatalf("expected 0 for negative tail, got %d", got)
	}
}

func TestTailLines_ShortFlag(t *testing.T) {
	args := []string{"--tail", "50", "file.log"}
	c, err := Parse(args)
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}
	if got := c.TailLines(); got != 50 {
		t.Fatalf("expected 50, got %d", got)
	}
}
