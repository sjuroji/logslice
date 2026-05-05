// Package config provides CLI flag parsing and configuration
// for the logslice tool.
package config

import (
	"flag"
	"fmt"
	"io"
	"time"
)

// Config holds all runtime options parsed from command-line flags.
type Config struct {
	// Files is the list of log files to process.
	Files []string

	// Pattern is an optional regex pattern to filter lines.
	Pattern string

	// Since filters lines with timestamps after this time (inclusive).
	Since time.Time

	// Until filters lines with timestamps before this time (inclusive).
	Until time.Time

	// MaxLines limits the number of output lines (0 = unlimited).
	MaxLines int

	// ShowLineNumbers prepends line numbers to output.
	ShowLineNumbers bool

	// Prefix prepends a custom string to each output line.
	Prefix string

	// TimeFormat is the timestamp layout used when parsing Since/Until.
	TimeFormat string
}

// Parse reads flags from args and returns a populated Config.
// Output and error messages are written to out and errOut respectively.
func Parse(args []string, out io.Writer, errOut io.Writer) (*Config, error) {
	fs := flag.NewFlagSet("logslice", flag.ContinueOnError)
	fs.SetOutput(errOut)

	pattern := fs.String("pattern", "", "regex pattern to filter log lines")
	since := fs.String("since", "", "include lines at or after this timestamp")
	until := fs.String("until", "", "include lines at or before this timestamp")
	maxLines := fs.Int("max-lines", 0, "maximum number of lines to output (0 = unlimited)")
	lineNumbers := fs.Bool("line-numbers", false, "prepend line numbers to output")
	prefix := fs.String("prefix", "", "prepend a custom string to each output line")
	timeFormat := fs.String("time-format", "", "timestamp layout for --since and --until parsing")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	cfg := &Config{
		Files:          fs.Args(),
		Pattern:        *pattern,
		MaxLines:       *maxLines,
		ShowLineNumbers: *lineNumbers,
		Prefix:         *prefix,
		TimeFormat:     *timeFormat,
	}

	var err error
	if *since != "" {
		cfg.Since, err = parseTime(*since, *timeFormat)
		if err != nil {
			return nil, fmt.Errorf("--since: %w", err)
		}
	}
	if *until != "" {
		cfg.Until, err = parseTime(*until, *timeFormat)
		if err != nil {
			return nil, fmt.Errorf("--until: %w", err)
		}
	}

	return cfg, nil
}
