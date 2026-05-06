package config

import (
	"testing"
	"time"
)

func TestValidate_NilConfig(t *testing.T) {
	if err := Validate(nil); err == nil {
		t.Fatal("expected error for nil config")
	}
}

func TestValidate_DefaultConfig(t *testing.T) {
	cfg := &Config{Workers: 1}
	if err := Validate(cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidate_NegativeMaxLines(t *testing.T) {
	cfg := &Config{MaxLines: -1, Workers: 1}
	err := Validate(cfg)
	if err == nil {
		t.Fatal("expected error for negative max-lines")
	}
	var ve *ValidationError
	if ok := errorAs(err, &ve); !ok || ve.Field != "max-lines" {
		t.Fatalf("expected ValidationError on max-lines, got %v", err)
	}
}

func TestValidate_EndBeforeStart(t *testing.T) {
	now := time.Now()
	cfg := &Config{
		Start:   now,
		End:     now.Add(-time.Hour),
		Workers: 1,
	}
	err := Validate(cfg)
	if err == nil {
		t.Fatal("expected error when end is before start")
	}
	var ve *ValidationError
	if ok := errorAs(err, &ve); !ok || ve.Field != "end" {
		t.Fatalf("expected ValidationError on end, got %v", err)
	}
}

func TestValidate_ValidRange(t *testing.T) {
	now := time.Now()
	cfg := &Config{
		Start:   now,
		End:     now.Add(time.Hour),
		Workers: 1,
	}
	if err := Validate(cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidate_ZeroWorkers(t *testing.T) {
	cfg := &Config{Workers: 0}
	err := Validate(cfg)
	if err == nil {
		t.Fatal("expected error for zero workers")
	}
	var ve *ValidationError
	if ok := errorAs(err, &ve); !ok || ve.Field != "workers" {
		t.Fatalf("expected ValidationError on workers, got %v", err)
	}
}

func TestValidate_EqualStartEnd(t *testing.T) {
	now := time.Now()
	cfg := &Config{Start: now, End: now, Workers: 1}
	if err := Validate(cfg); err != nil {
		t.Fatalf("equal start/end should be valid, got: %v", err)
	}
}

// errorAs is a thin wrapper so the test file has no direct errors import.
func errorAs(err error, target **ValidationError) bool {
	if ve, ok := err.(*ValidationError); ok {
		*target = ve
		return true
	}
	return false
}
