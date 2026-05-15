// Package sample provides line sampling for log output,
// allowing every Nth line to be emitted while discarding the rest.
package sample

import "fmt"

// Sampler keeps a running counter and accepts every Nth line.
type Sampler struct {
	n       int
	counter int
}

// New returns a new Sampler that emits one line in every n.
// If n is less than 1 an error is returned.
func New(n int) (*Sampler, error) {
	if n < 1 {
		return nil, fmt.Errorf("sample: n must be >= 1, got %d", n)
	}
	return &Sampler{n: n}, nil
}

// Accept returns true when the current line should be emitted.
// It increments an internal counter on every call.
func (s *Sampler) Accept(_ string) bool {
	if s == nil {
		return true
	}
	s.counter++
	return s.counter%s.n == 0
}

// Reset sets the internal counter back to zero.
func (s *Sampler) Reset() {
	if s == nil {
		return
	}
	s.counter = 0
}

// N returns the configured sample interval.
func (s *Sampler) N() int {
	if s == nil {
		return 1
	}
	return s.n
}
