// Package timerange provides time range filtering for log lines.
// It determines whether a given timestamp falls within a specified
// start and end time window.
package timerange

import (
	"errors"
	"time"
)

// ErrInvalidRange is returned when the start time is after the end time.
var ErrInvalidRange = errors.New("timerange: start must be before or equal to end")

// Range represents an inclusive time window [Start, End].
type Range struct {
	Start time.Time
	End   time.Time
}

// New creates a new Range. If start is after end, ErrInvalidRange is returned.
// A zero-value start or end is treated as unbounded on that side.
func New(start, end time.Time) (*Range, error) {
	if !start.IsZero() && !end.IsZero() && start.After(end) {
		return nil, ErrInvalidRange
	}
	return &Range{Start: start, End: end}, nil
}

// Contains reports whether t falls within the range.
// A zero-value Start means no lower bound; a zero-value End means no upper bound.
func (r *Range) Contains(t time.Time) bool {
	if !r.Start.IsZero() && t.Before(r.Start) {
		return false
	}
	if !r.End.IsZero() && t.After(r.End) {
		return false
	}
	return true
}

// IsUnbounded reports whether the range has no constraints on either side.
func (r *Range) IsUnbounded() bool {
	return r.Start.IsZero() && r.End.IsZero()
}

// String returns a human-readable representation of the range.
func (r *Range) String() string {
	const layout = time.RFC3339
	start := "*"
	end := "*"
	if !r.Start.IsZero() {
		start = r.Start.Format(layout)
	}
	if !r.End.IsZero() {
		end = r.End.Format(layout)
	}
	return "[" + start + ", " + end + "]"
}
