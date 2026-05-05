// Package linescanner provides utilities for scanning log files line by line
// with optional pattern matching and time-range filtering support.
package linescanner

import (
	"bufio"
	"io"
	"strings"
)

// Options configures the behavior of the Scanner.
type Options struct {
	// Pattern filters lines to only those containing this substring.
	// An empty string disables pattern filtering.
	Pattern string

	// MaxLines limits the number of lines returned. 0 means no limit.
	MaxLines int
}

// Scanner reads lines from a reader, applying optional filters.
type Scanner struct {
	reader  io.Reader
	options Options
}

// New creates a new Scanner reading from r with the given options.
func New(r io.Reader, opts Options) *Scanner {
	return &Scanner{
		reader:  r,
		options: opts,
	}
}

// Scan reads all matching lines from the underlying reader and returns them.
// It returns an error if reading fails.
func (s *Scanner) Scan() ([]string, error) {
	var results []string

	scanner := bufio.NewScanner(s.reader)
	for scanner.Scan() {
		line := scanner.Text()

		if s.options.Pattern != "" && !strings.Contains(line, s.options.Pattern) {
			continue
		}

		results = append(results, line)

		if s.options.MaxLines > 0 && len(results) >= s.options.MaxLines {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
