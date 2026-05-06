// Package config parses and validates CLI flags for logslice.
package config

import (
	"flag"
	"io"
	"time"
)

// Config holds the resolved configuration for a single logslice run.
type Config struct {
	// Positional arguments after flag parsing.
	Files []string

	// Filtering
	Pattern string
	Start   time.Time
	End     time.Time
	Layout  string

	// Output
	MaxLines    int
	LineNumbers bool
	Prefix      string

	// Behaviour
	Recursive bool
	Suffixes  []string
	Verbose   bool
	Stats     bool
}

// Parse reads flags from args, writing usage to w, and returns a Config.
// It does not call Validate; callers should do so separately.
func Parse(args []string, w io.Writer) (*Config, error) {
	fs := flag.NewFlagSet("logslice", flag.ContinueOnError)
	fs.SetOutput(w)

	var (
		pattern     = fs.String("pattern", "", "regex pattern to match against log lines")
		startStr    = fs.String("start", "", "start of time range (RFC3339 or date)")
		endStr      = fs.String("end", "", "end of time range (RFC3339 or date)")
		layout      = fs.String("layout", "", "explicit Go time layout for parsing timestamps")
		maxLines    = fs.Int("max-lines", 0, "maximum number of matching lines to output (0 = unlimited)")
		lineNumbers = fs.Bool("line-numbers", false, "prefix each output line with its line number")
		prefix      = fs.String("prefix", "", "static string prefix for every output line")
		recursive   = fs.Bool("recursive", false, "recurse into directories")
		suffix      = fs.String("suffix", ".log", "comma-separated file suffixes when expanding directories")
		verbose     = fs.Bool("verbose", false, "print verbose progress information")
		stats       = fs.Bool("stats", false, "print processing statistics after completion")
	)

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	cfg := &Config{
		Files:       fs.Args(),
		Pattern:     *pattern,
		Layout:      *layout,
		MaxLines:    *maxLines,
		LineNumbers: *lineNumbers,
		Prefix:      *prefix,
		Recursive:   *recursive,
		Verbose:     *verbose,
		Stats:       *stats,
		Suffixes:    splitSuffixes(*suffix),
	}

	if *startStr != "" {
		t, err := parseTime(*startStr, *layout)
		if err != nil {
			return nil, err
		}
		cfg.Start = t
	}

	if *endStr != "" {
		t, err := parseTime(*endStr, *layout)
		if err != nil {
			return nil, err
		}
		cfg.End = t
	}

	return cfg, nil
}

func splitSuffixes(s string) []string {
	if s == "" {
		return nil
	}
	var out []string
	for _, part := range splitComma(s) {
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}

func splitComma(s string) []string {
	var parts []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ',' {
			parts = append(parts, s[start:i])
			start = i + 1
		}
	}
	return append(parts, s[start:])
}
