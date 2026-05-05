package timeparser

import (
	"fmt"
	"time"
)

// Common log timestamp formats to try when parsing
var supportedFormats = []string{
	"2006-01-02T15:04:05Z07:00",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04:05.000",
	"2006-01-02 15:04:05.000000",
	"02/Jan/2006:15:04:05 -0700",
	"Jan 02 15:04:05",
	"2006/01/02 15:04:05",
}

// Parser holds configuration for time parsing.
type Parser struct {
	Location *time.Location
	Formats  []string
}

// New creates a Parser with default formats and the given timezone.
func New(loc *time.Location) *Parser {
	if loc == nil {
		loc = time.UTC
	}
	return &Parser{
		Location: loc,
		Formats:  supportedFormats,
	}
}

// Parse attempts to parse a timestamp string using all supported formats.
// Returns the parsed time and the format that matched, or an error.
func (p *Parser) Parse(s string) (time.Time, string, error) {
	for _, format := range p.Formats {
		t, err := time.ParseInLocation(format, s, p.Location)
		if err == nil {
			return t, format, nil
		}
	}
	return time.Time{}, "", fmt.Errorf("timeparser: unable to parse %q with any known format", s)
}

// ParseWithFormat parses a timestamp using a specific format string.
func (p *Parser) ParseWithFormat(s, format string) (time.Time, error) {
	t, err := time.ParseInLocation(format, s, p.Location)
	if err != nil {
		return time.Time{}, fmt.Errorf("timeparser: %w", err)
	}
	return t, nil
}
