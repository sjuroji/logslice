// Package logfilter combines time-range and pattern filtering for log lines.
package logfilter

import (
	"strings"
	"time"

	"github.com/yourorg/logslice/internal/timeparser"
	"github.com/yourorg/logslice/internal/timerange"
)

// Filter holds the configuration for filtering log lines.
type Filter struct {
	tr      *timerange.TimeRange
	pattern string
	parser  *timeparser.Parser
}

// Config holds the options for creating a new Filter.
type Config struct {
	Start    time.Time
	End      time.Time
	Pattern  string
	Location *time.Location
}

// New creates a new Filter from the given Config.
// If both Start and End are zero, time filtering is disabled.
func New(cfg Config) (*Filter, error) {
	var tr *timerange.TimeRange

	if !cfg.Start.IsZero() || !cfg.End.IsZero() {
		var err error
		tr, err = timerange.New(cfg.Start, cfg.End)
		if err != nil {
			return nil, err
		}
	}

	loc := cfg.Location
	if loc == nil {
		loc = time.UTC
	}

	return &Filter{
		tr:      tr,
		pattern: cfg.Pattern,
		parser:  timeparser.New(loc),
	}, nil
}

// Match reports whether a raw log line passes all active filters.
// Time filtering is attempted by parsing the line; if parsing fails the
// line is included when no time range is set, excluded otherwise.
func (f *Filter) Match(line string) bool {
	if f.pattern != "" && !strings.Contains(line, f.pattern) {
		return false
	}

	if f.tr == nil {
		return true
	}

	t, err := f.parser.Parse(line)
	if err != nil {
		// Cannot determine timestamp — exclude from time-filtered output.
		return false
	}

	return f.tr.Contains(t)
}
