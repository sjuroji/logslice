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
		t.Fatal("expected error for negative MaxLines")
	}
	var ve *ValidationError
	if ok := isValidationError(err, &ve); !ok {
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
		t.Fatal("expected error when end is before start")
	}
	var ve *ValidationError
	if ok := isValidationError(err, &ve); !ok {
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

func TestValidate_EqualStartEnd(t *testing.T) {
	now := time.Now()
	cfg := &Config{Start: now, End: now}
	if err := Validate(cfg); err != nil {
		t.Fatalf("unexpected error for equal start/end: %v", err)
	}
}

// isValidationError is a helper to type-assert without generics for Go 1.21 compat.
func isValidationError(err error, out **ValidationError) bool {
	if ve, ok := err.(*ValidationError); ok {
		*out = ve
		return true
	}
	return false
}
