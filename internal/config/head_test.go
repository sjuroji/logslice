package config_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestHeadLines_NilConfig(t *testing.T) {
	var c *config.Config
	if got := c.HeadLines(); got != 0 {
		t.Fatalf("expected 0 for nil Config, got %d", got)
	}
}

func TestHeadLines_DefaultConfig(t *testing.T) {
	cfg, err := config.Parse([]string{})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if got := cfg.HeadLines(); got != 0 {
		t.Fatalf("expected 0 by default, got %d", got)
	}
}

func TestHeadLines_ExplicitValue(t *testing.T) {
	cfg, err := config.Parse([]string{"--head", "20"})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if got := cfg.HeadLines(); got != 20 {
		t.Fatalf("expected 20, got %d", got)
	}
}

func TestHeadLines_ShortFlag(t *testing.T) {
	cfg, err := config.Parse([]string{"-H", "5"})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if got := cfg.HeadLines(); got != 5 {
		t.Fatalf("expected 5, got %d", got)
	}
}
