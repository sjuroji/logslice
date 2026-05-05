// Package timeparser provides flexible timestamp parsing for log lines.
//
// It supports a variety of common log timestamp formats found in popular
// logging frameworks and web servers, including ISO 8601, Apache/Nginx
// combined log format, syslog, and others.
//
// Usage:
//
//	p := timeparser.New(time.UTC)
//
//	// Auto-detect format
//	t, format, err := p.Parse("2024-03-15T08:30:00Z")
//
//	// Use a specific format
//	t, err = p.ParseWithFormat("15/Mar/2024:08:30:00 +0000", "02/Jan/2006:15:04:05 -0700")
//
// Custom formats can be added by appending to the Formats slice on the Parser.
package timeparser
