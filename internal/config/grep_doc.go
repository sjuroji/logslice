// Package config provides flag parsing and configuration helpers for logslice.
//
// # Grep flags
//
// The grep sub-module registers two flags:
//
//   --pattern / -p
//     A comma-separated list of patterns (plain string or regular expression)
//     that a log line must match to be included in the output.  When multiple
//     patterns are supplied the line must satisfy at least one of them
//     (logical OR).  When omitted every line is considered a match.
//
//   --exclude / -e
//     A comma-separated list of patterns whose matching lines are dropped from
//     the output even when they satisfy --pattern.  Exclusions are evaluated
//     after inclusions, so they act as a post-filter veto.
//
// Both flags delegate to the internal/grep package for compilation and
// matching, which means full RE2-compatible regular expressions are supported.
package config
