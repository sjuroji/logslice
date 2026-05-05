// Package timeparser provides utilities for parsing time strings
// in multiple common formats used in log files.
package timeparser

import (
	"fmt"
	"time"
)

// SupportedFormats lists the time layouts tried in order when no explicit
// format is provided.
var SupportedFormats = []string{
	time.RFC3339,
	time.RFC3339Nano,
	"2006-01-02 15:04:05",
	"2006-01-02 15:04:05.999999999",
	"2006-01-02T15:04:05",
	"2006-01-02",
	"02/Jan/2006:15:04:05 -0700",
}

// Parser parses time strings using a fixed location.
type Parser struct {
	loc *time.Location
}

// New creates a Parser that interprets times without timezone info in loc.
// If loc is nil, time.UTC is used.
func New(loc *time.Location) *Parser {
	if loc == nil {
		loc = time.UTC
	}
	return &Parser{loc: loc}
}

// Parse attempts to parse s using each of SupportedFormats in order.
// It returns the first successful result or an error if none match.
func (p *Parser) Parse(s string) (time.Time, error) {
	for _, layout := range SupportedFormats {
		t, err := time.ParseInLocation(layout, s, p.loc)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("timeparser: unrecognised time string %q", s)
}

// ParseWithFormat parses s using the explicit layout, honouring the parser
// location for timezone-naive strings.
func (p *Parser) ParseWithFormat(layout, s string) (time.Time, error) {
	t, err := time.ParseInLocation(layout, s, p.loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("timeparser: cannot parse %q with layout %q: %w", s, layout, err)
	}
	return t, nil
}
