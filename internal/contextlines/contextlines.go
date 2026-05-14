// Package contextlines provides a sliding-window buffer that captures N lines
// of leading and trailing context around lines accepted by a match predicate,
// mirroring the behaviour of grep -B / -A / -C.
package contextlines

// Line pairs a raw text line with the 1-based line number at which it appeared
// in the source stream.
type Line struct {
	Number int
	Text   string
}

// Buffer accumulates context lines around matches.
type Buffer struct {
	before int
	after  int

	ring    []Line // circular buffer for pre-match lines
	head    int    // next write position in ring
	count   int    // lines currently stored in ring

	pending int    // remaining after-context lines to flush
	out     []Line // collected output lines
}

// New creates a Buffer that will emit up to before leading lines and after
// trailing lines around each accepted line. Both values are clamped to >= 0.
func New(before, after int) *Buffer {
	if before < 0 {
		before = 0
	}
	if after < 0 {
		after = 0
	}
	return &Buffer{
		before: before,
		after:  after,
		ring:   make([]Line, before),
	}
}

// Feed presents line to the buffer. matched indicates whether the line
// satisfies the caller's filter predicate. Feed returns any lines that are
// now ready to be emitted (pre-context + match + post-context).
func (b *Buffer) Feed(line Line, matched bool) []Line {
	b.out = b.out[:0]

	if matched {
		// Flush buffered pre-context lines in chronological order.
		if b.before > 0 && b.count > 0 {
			start := (b.head - b.count + len(b.ring)) % len(b.ring)
			for i := 0; i < b.count; i++ {
				b.out = append(b.out, b.ring[(start+i)%len(b.ring)])
			}
			b.count = 0
		}
		b.out = append(b.out, line)
		b.pending = b.after
	} else if b.pending > 0 {
		b.out = append(b.out, line)
		b.pending--
	} else if b.before > 0 {
		// Store in ring for potential future pre-context use.
		b.ring[b.head] = line
		b.head = (b.head + 1) % len(b.ring)
		if b.count < b.before {
			b.count++
		}
	}

	// Return a copy so callers own the slice.
	result := make([]Line, len(b.out))
	copy(result, b.out)
	return result
}
