// Package count provides a line-counting reader that tracks how many lines
// pass through a processing pipeline without buffering the full output.
package count

import "sync/atomic"

// Counter tracks the number of lines observed.
type Counter struct {
	n atomic.Int64
}

// New returns a new Counter initialised to zero.
func New() *Counter {
	return &Counter{}
}

// Add increments the counter by delta.
func (c *Counter) Add(delta int64) {
	c.n.Add(delta)
}

// Inc increments the counter by one.
func (c *Counter) Inc() {
	c.n.Add(1)
}

// Value returns the current count.
func (c *Counter) Value() int64 {
	return c.n.Load()
}

// Reset sets the counter back to zero.
func (c *Counter) Reset() {
	c.n.Store(0)
}

// Accumulator collects per-file line counts and exposes an aggregate total.
type Accumulator struct {
	files map[string]int64
	total int64
}

// NewAccumulator returns an initialised Accumulator.
func NewAccumulator() *Accumulator {
	return &Accumulator{files: make(map[string]int64)}
}

// Record stores the line count for the named file and adds it to the total.
func (a *Accumulator) Record(name string, lines int64) {
	a.files[name] = lines
	a.total += lines
}

// Total returns the sum of all recorded line counts.
func (a *Accumulator) Total() int64 { return a.total }

// Files returns a snapshot copy of the per-file counts.
func (a *Accumulator) Files() map[string]int64 {
	out := make(map[string]int64, len(a.files))
	for k, v := range a.files {
		out[k] = v
	}
	return out
}
