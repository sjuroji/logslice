// Package config handles command-line argument parsing for logslice.
//
// It exposes a single Parse function that accepts a slice of argument
// strings (compatible with os.Args[1:]) and returns a Config struct
// that drives the rest of the pipeline.
//
// Supported flags:
//
//	--pattern      regex filter applied to each log line
//	--since        lower bound timestamp (inclusive)
//	--until        upper bound timestamp (inclusive)
//	--max-lines    cap on output lines (0 = unlimited)
//	--line-numbers prepend line numbers to output
//	--prefix       prepend a custom string to every output line
//	--time-format  explicit Go time layout for --since / --until
//
// When --time-format is omitted, Parse tries a set of common layouts
// automatically (RFC3339, ISO 8601 variants, US date formats).
package config
