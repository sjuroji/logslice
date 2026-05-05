package logfilter_test

import (
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/logfilter"
)

var (
	base  = time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	start = base
	end   = base.Add(2 * time.Hour)
)

func newFilter(t *testing.T, cfg logfilter.Config) *logfilter.Filter {
	t.Helper()
	f, err := logfilter.New(cfg)
	if err != nil {
		t.Fatalf("logfilter.New: %v", err)
	}
	return f
}

func TestMatch_NoFilters(t *testing.T) {
	f := newFilter(t, logfilter.Config{})
	if !f.Match("any log line") {
		t.Error("expected match with no filters")
	}
}

func TestMatch_PatternMatch(t *testing.T) {
	f := newFilter(t, logfilter.Config{Pattern: "ERROR"})
	if !f.Match("2024-01-15 10:30:00 ERROR something broke") {
		t.Error("expected match for line containing ERROR")
	}
	if f.Match("2024-01-15 10:30:00 INFO all good") {
		t.Error("expected no match for line without ERROR")
	}
}

func TestMatch_TimeRange_Inside(t *testing.T) {
	f := newFilter(t, logfilter.Config{Start: start, End: end})
	// Line timestamp within range
	if !f.Match("2024-01-15 11:00:00 INFO within range") {
		t.Error("expected match for timestamp within range")
	}
}

func TestMatch_TimeRange_Outside(t *testing.T) {
	f := newFilter(t, logfilter.Config{Start: start, End: end})
	// Line timestamp outside range
	if f.Match("2024-01-15 13:00:00 INFO outside range") {
		t.Error("expected no match for timestamp outside range")
	}
}

func TestMatch_TimeRange_UnparsableLine(t *testing.T) {
	f := newFilter(t, logfilter.Config{Start: start, End: end})
	if f.Match("no timestamp here") {
		t.Error("expected no match when timestamp cannot be parsed and range is set")
	}
}

func TestMatch_PatternAndTimeRange(t *testing.T) {
	f := newFilter(t, logfilter.Config{
		Start:   start,
		End:     end,
		Pattern: "WARN",
	})
	// Both conditions satisfied
	if !f.Match("2024-01-15 10:45:00 WARN disk usage high") {
		t.Error("expected match when both pattern and time match")
	}
	// Pattern matches but time is outside
	if f.Match("2024-01-15 13:00:00 WARN too late") {
		t.Error("expected no match when time is outside range")
	}
	// Time matches but pattern does not
	if f.Match("2024-01-15 10:45:00 INFO no pattern") {
		t.Error("expected no match when pattern is absent")
	}
}

func TestNew_InvalidRange(t *testing.T) {
	_, err := logfilter.New(logfilter.Config{
		Start: end,
		End:   start, // reversed
	})
	if err == nil {
		t.Error("expected error for reversed time range")
	}
}
