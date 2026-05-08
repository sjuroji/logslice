// Package tail provides support for reading the last N lines of a file
// or stream, useful for tailing log files before applying filters.
package tail

import (
	"bufio"
	"container/ring"
	"io"
)

// Tailer holds the last N lines read from a source.
type Tailer struct {
	n int
	buf *ring.Ring
}

// New creates a Tailer that retains the last n lines.
// If n <= 0, New returns nil and callers should skip tailing.
func New(n int) *Tailer {
	if n <= 0 {
		return nil
	}
	return &Tailer{
		n:   n,
		buf: ring.New(n),
	}
}

// ReadAll consumes r, keeping only the last n lines in memory.
func (t *Tailer) ReadAll(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		t.buf.Value = line
		t.buf = t.buf.Next()
	}
	return scanner.Err()
}

// Lines returns the retained lines in order from oldest to newest.
// Lines that were never written (ring slots still nil) are omitted.
func (t *Tailer) Lines() []string {
	var out []string
	t.buf.Do(func(v any) {
		if v != nil {
			out = append(out, v.(string))
		}
	})
	return out
}

// Len returns the number of lines currently retained.
func (t *Tailer) Len() int {
	return len(t.Lines())
}
