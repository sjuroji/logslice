package timerange_test

import (
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/timerange"
)

var (
	t0 = time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	t1 = time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	t2 = time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC)
)

func TestNew_ValidRange(t *testing.T) {
	r, err := timerange.New(t0, t2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r == nil {
		t.Fatal("expected non-nil Range")
	}
}

func TestNew_InvalidRange(t *testing.T) {
	_, err := timerange.New(t2, t0)
	if err == nil {
		t.Fatal("expected error for start > end")
	}
}

func TestNew_EqualBounds(t *testing.T) {
	_, err := timerange.New(t1, t1)
	if err != nil {
		t.Fatalf("equal bounds should be valid, got: %v", err)
	}
}

func TestContains_WithinRange(t *testing.T) {
	r, _ := timerange.New(t0, t2)
	if !r.Contains(t1) {
		t.Error("expected t1 to be within range")
	}
}

func TestContains_OutsideRange(t *testing.T) {
	r, _ := timerange.New(t1, t2)
	if r.Contains(t0) {
		t.Error("expected t0 to be outside range")
	}
}

func TestContains_UnboundedStart(t *testing.T) {
	r, _ := timerange.New(time.Time{}, t2)
	if !r.Contains(t0) {
		t.Error("expected t0 to be within unbounded-start range")
	}
}

func TestContains_UnboundedEnd(t *testing.T) {
	r, _ := timerange.New(t0, time.Time{})
	if !r.Contains(t2) {
		t.Error("expected t2 to be within unbounded-end range")
	}
}

func TestIsUnbounded(t *testing.T) {
	r, _ := timerange.New(time.Time{}, time.Time{})
	if !r.IsUnbounded() {
		t.Error("expected range to be unbounded")
	}

	r2, _ := timerange.New(t0, t2)
	if r2.IsUnbounded() {
		t.Error("expected range to be bounded")
	}
}

func TestString(t *testing.T) {
	r, _ := timerange.New(time.Time{}, time.Time{})
	if got := r.String(); got != "[*, *]" {
		t.Errorf("unexpected String() = %q", got)
	}
}
