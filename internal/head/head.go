package head

import "io"

// Reader returns up to N leading lines from a source.
// A nil Reader is returned when n <= 0, meaning no limit applies.
type Reader struct {
	n    int
	seen int
}

// New returns a *Reader that will emit at most n lines.
// If n <= 0, New returns nil to signal "no head limit".
func New(n int) *Reader {
	if n <= 0 {
		return nil
	}
	return &Reader{n: n}
}

// Done reports whether the Reader has already consumed its quota of lines.
// Calling Done on a nil *Reader always returns false.
func (r *Reader) Done() bool {
	if r == nil {
		return false
	}
	return r.seen >= r.n
}

// Accept records one line and returns true when the line falls within the
// first n lines. Once the quota is exhausted it returns false for every
// subsequent call. Calling Accept on a nil *Reader always returns true.
func (r *Reader) Accept() bool {
	if r == nil {
		return true
	}
	if r.seen >= r.n {
		return false
	}
	r.seen++
	return true
}

// ReadAll reads all lines from lines and returns at most n of them.
// It is a convenience wrapper for callers that already have a slice.
func ReadAll(lines []string, n int) []string {
	r := New(n)
	if r == nil {
		out := make([]string, len(lines))
		copy(out, lines)
		return out
	}
	out := make([]string, 0, n)
	for _, l := range lines {
		if !r.Accept() {
			break
		}
		out = append(out, l)
	}
	return out
}

// ErrEOF is a sentinel returned by some callers to signal early termination.
var ErrEOF = io.EOF
