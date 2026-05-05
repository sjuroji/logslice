package config

import (
	"fmt"
	"time"
)

// knownFormats lists common timestamp layouts tried in order when no
// explicit format is provided.
var knownFormats = []string{
	time.RFC3339,
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02",
	"01/02/2006 15:04:05",
	"01/02/2006",
}

// parseTime attempts to parse value using layout if provided, otherwise
// tries each entry in knownFormats. It returns the first successful parse
// in UTC, or an error if no format matches.
func parseTime(value, layout string) (time.Time, error) {
	if layout != "" {
		t, err := time.Parse(layout, value)
		if err != nil {
			return time.Time{}, fmt.Errorf("cannot parse %q with format %q: %w", value, layout, err)
		}
		return t.UTC(), nil
	}

	for _, fmt := range knownFormats {
		if t, err := time.Parse(fmt, value); err == nil {
			return t.UTC(), nil
		}
	}

	return time.Time{}, fmt.Errorf("cannot parse timestamp %q: no matching format found", value)
}
