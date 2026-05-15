// Package unique provides a line-uniqueness filter that passes only the first
// occurrence of each distinct line, optionally limited to a sliding window.
package unique

import "errors"

// Filter accepts only lines not seen before within an optional window size.
// A window of 0 means unlimited (keep all seen lines in memory).
type Filter struct {
	seen   map[string]struct{}
	window int
	queue  []string
	count  int
}

// New creates a new Filter. window controls how many recent lines are tracked;
// pass 0 for an unlimited history. Returns an error if window is negative.
func New(window int) (*Filter, error) {
	if window < 0 {
		return nil, errors.New("unique: window must be >= 0")
	}
	return &Filter{
		seen:   make(map[string]struct{}),
		window: window,
	}, nil
}

// Accept returns true if line has not been seen within the current window.
// A nil Filter always returns true.
func (f *Filter) Accept(line string) bool {
	if f == nil {
		return true
	}
	if _, dup := f.seen[line]; dup {
		f.count++
		return false
	}
	f.seen[line] = struct{}{}
	if f.window > 0 {
		f.queue = append(f.queue, line)
		if len(f.queue) > f.window {
			old := f.queue[0]
			f.queue = f.queue[1:]
			delete(f.seen, old)
		}
	}
	return true
}

// Suppressed returns the number of duplicate lines rejected so far.
func (f *Filter) Suppressed() int {
	if f == nil {
		return 0
	}
	return f.count
}

// Reset clears all tracked state.
func (f *Filter) Reset() {
	if f == nil {
		return
	}
	f.seen = make(map[string]struct{})
	f.queue = f.queue[:0]
	f.count = 0
}
