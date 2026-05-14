package config_test

import (
	"flag"
	"testing"

	"github.com/yourorg/logslice/internal/config"
)

func TestOutputFormat_NilFlagSet(t *testing.T) {
	get := config.OutputFormat(nil)
	fmt, color := get()
	if fmt != "plain" {
		t.Errorf("expected plain, got %q", fmt)
	}
	if color {
		t.Error("expected color=false for nil FlagSet")
	}
}

func TestOutputFormat_Defaults(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	get := config.OutputFormat(fs)
	_ = fs.Parse([]string{})

	fmt, color := get()
	if fmt != "plain" {
		t.Errorf("expected plain, got %q", fmt)
	}
	if color {
		t.Error("expected color=false by default")
	}
}

func TestOutputFormat_ExplicitFormat(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	get := config.OutputFormat(fs)
	_ = fs.Parse([]string{"--format", "json"})

	fmt, _ := get()
	if fmt != "json" {
		t.Errorf("expected json, got %q", fmt)
	}
}

func TestOutputFormat_ColorFlag(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	get := config.OutputFormat(fs)
	_ = fs.Parse([]string{"--color"})

	_, color := get()
	if !color {
		t.Error("expected color=true when --color passed")
	}
}

func TestFormat_NilConfig(t *testing.T) {
	if got := config.Format(nil); got != "plain" {
		t.Errorf("expected plain, got %q", got)
	}
}

func TestColorEnabled_NilConfig(t *testing.T) {
	if config.ColorEnabled(nil) {
		t.Error("expected false for nil config")
	}
}
