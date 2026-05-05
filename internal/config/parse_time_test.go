package config

import (
	"testing"
	"time"
)

func TestParseTime_RFC3339(t *testing.T) {
	got, err := parseTime("2024-06-01T12:30:00Z", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := time.Date(2024, 6, 1, 12, 30, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParseTime_DateOnly(t *testing.T) {
	got, err := parseTime("2024-06-01", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParseTime_SpaceSeparated(t *testing.T) {
	got, err := parseTime("2024-06-01 09:15:00", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := time.Date(2024, 6, 1, 9, 15, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParseTime_ExplicitLayout(t *testing.T) {
	got, err := parseTime("01/15/2024", "01/02/2006")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParseTime_UnknownFormat(t *testing.T) {
	_, err := parseTime("garbage-value", "")
	if err == nil {
		t.Error("expected error for unknown format, got nil")
	}
}

func TestParseTime_BadExplicitLayout(t *testing.T) {
	_, err := parseTime("2024-06-01", "Mon Jan 2")
	if err == nil {
		t.Error("expected error for mismatched explicit layout, got nil")
	}
}
