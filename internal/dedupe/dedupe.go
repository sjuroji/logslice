// Package dedupe provides line deduplication for log output.
// It tracks previously seen lines and filters out exact duplicates,
// optionally reporting the number of times a line was suppressed.
package dedupe

import "sync"

// Filter tracks seen lines and suppresses duplicates.
type Filter struct {
	mu   sync.Mutex
	seen map[string]int
}

// New returns a new duplicate-line Filter.
func New() *Filter {
	return &Filter{seen: make(map[string]int)}
}

// Accept returns true if line has not been seen before.
// Subsequent calls with the same line return false and increment
// the suppression counter for that line.
func (f *Filter) Accept(line string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, exists := f.seen[line]; exists {
		f.seen[line]++
		return false
	}
	f.seen[line] = 0
	return true
}

// Count returns the number of times line was suppressed (i.e. seen
// after the first occurrence). Returns 0 for unseen lines.
func (f *Filter) Count(line string) int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.seen[line]
}

// Reset clears all tracked lines.
func (f *Filter) Reset() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.seen = make(map[string]int)
}

// Unique returns the number of distinct lines seen so far.
func (f *Filter) Unique() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.seen)
}
