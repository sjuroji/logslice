// Package grep provides compiled regular-expression pattern matching
// used by logslice to filter log lines.
package grep

import (
	"fmt"
	"regexp"
)

// Matcher holds a compiled regular expression and exposes a single
// Match method so callers never deal with *regexp.Regexp directly.
type Matcher struct {
	re *regexp.Regexp
}

// New compiles pattern and returns a Matcher. An empty pattern always
// matches every line. Returns an error when the pattern is invalid.
func New(pattern string) (*Matcher, error) {
	if pattern == "" {
		return &Matcher{}, nil
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("grep: compile pattern %q: %w", pattern, err)
	}
	return &Matcher{re: re}, nil
}

// Match reports whether line matches the compiled pattern.
// When no pattern was provided (empty Matcher) it always returns true.
func (m *Matcher) Match(line string) bool {
	if m.re == nil {
		return true
	}
	return m.re.MatchString(line)
}

// Pattern returns the original pattern string, or an empty string when
// the Matcher was created with no pattern.
func (m *Matcher) Pattern() string {
	if m.re == nil {
		return ""
	}
	return m.re.String()
}
