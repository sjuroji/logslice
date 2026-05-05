package config

import (
	"testing"
	"time"
)

func TestValidate_NilConfig(t *testing.T) {
	if err := Validate(nil); err == nil {
		t.Fatal("expected error for nil config, got nil")
	}
}

func TestValidate_DefaultConfig(t *testing.T) {
	cfg := &Config{}
	if err := Validate(cfg); err != nil {
		t.Fatalf("unexpected error for default config: %v", err)
	}
}

func TestValidate_NegativeMaxLines(t *testing.T) {
	cfg := &Config{MaxLines: -1}
	err := Validate(cfg)
	if err == nil {
		t.Fatal("expected error for negative MaxLines, got nil")
	}
	var ve *ValidationError
	if ok := errorAs(err, &ve); !ok {
		t.Fatalf("expected *ValidationError, got %T", err)
	}
	if ve.Field != "max-lines" {
		t.Errorf("expected field 'max-lines', got %q", ve.Field)
	}
}

func TestValidate_EndBeforeStart(t *testing.T) {
	now := time.Now()
	cfg := &Config{
		Start: now,
		End:   now.Add(-time.Hour),
	}
	err := Validate(cfg)
	if err == nil {
		t.Fatal("expected error when end is before start, got nil")
	}
	var ve *ValidationError
	if ok := errorAs(err, &ve); !ok {
		t.Fatalf("expected *ValidationError, got %T", err)
	}
	if ve.Field != "end" {
		t.Errorf("expected field 'end', got %q", ve.Field)
	}
}

func TestValidate_ValidRange(t *testing.T) {
	now := time.Now()
	cfg := &Config{
		Start: now,
		End:   now.Add(time.Hour),
	}
	if err := Validate(cfg); err != nil {
		t.Fatalf("unexpected error for valid range: %v", err)
	}
}

func TestValidate_StartOnly(t *testing.T) {
	cfg := &Config{Start: time.Now()}
	if err := Validate(cfg); err != nil {
		t.Fatalf("unexpected error for start-only config: %v", err)
	}
}

func TestValidate_EndOnly(t *testing.T) {
	cfg := &Config{End: time.Now()}
	if err := Validate(cfg); err != nil {
		t.Fatalf("unexpected error for end-only config: %v", err)
	}
}

// errorAs is a thin wrapper to avoid importing errors in test helpers.
func errorAs(err error, target **ValidationError) bool {
	if ve, ok := err.(*ValidationError); ok {
		*target = ve
		return true
	}
	return false
}
