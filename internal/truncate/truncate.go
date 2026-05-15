// Package truncate provides a line truncator that clips long lines to a
// maximum byte length, optionally appending a suffix such as "..." to
// indicate that the line was shortened.
package truncate

import "unicode/utf8"

// Truncator clips lines that exceed a maximum length.
type Truncator struct {
	maxLen int
	suffix string
}

// New returns a Truncator that clips lines longer than maxLen bytes.
// If suffix is non-empty it is appended to truncated lines; the suffix
// itself counts toward maxLen so the output never exceeds maxLen bytes.
// New returns nil when maxLen is zero or negative.
func New(maxLen int, suffix string) *Truncator {
	if maxLen <= 0 {
		return nil
	}
	if len(suffix) >= maxLen {
		suffix = ""
	}
	return &Truncator{maxLen: maxLen, suffix: suffix}
}

// Apply returns the line unchanged when the Truncator is nil or the line
// fits within the configured limit. Otherwise it returns a clipped copy
// with the suffix appended. Clipping always occurs on a valid UTF-8
// rune boundary so the result is never a malformed string.
func (t *Truncator) Apply(line string) string {
	if t == nil || len(line) <= t.maxLen {
		return line
	}
	cutAt := t.maxLen - len(t.suffix)
	// Walk back to a valid rune boundary.
	for cutAt > 0 && !utf8.RuneStart(line[cutAt]) {
		cutAt--
	}
	return line[:cutAt] + t.suffix
}

// MaxLen returns the configured maximum line length.
func (t *Truncator) MaxLen() int {
	if t == nil {
		return 0
	}
	return t.maxLen
}
