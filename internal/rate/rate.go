// Package rate provides a line-rate limiter that accepts at most N lines
// per second, dropping lines that exceed the budget for the current window.
package rate

import (
	"errors"
	"time"
)

// Limiter accepts lines up to a fixed rate (lines per second).
// A nil *Limiter is valid and accepts every line.
type Limiter struct {
	perSecond int
	windowStart time.Time
	count       int
	nowFn       func() time.Time
}

// New creates a Limiter that allows at most perSecond lines per second.
// Returns an error if perSecond is less than 1.
func New(perSecond int) (*Limiter, error) {
	if perSecond < 1 {
		return nil, errors.New("rate: perSecond must be >= 1")
	}
	return &Limiter{
		perSecond:   perSecond,
		windowStart: time.Now(),
		nowFn:       time.Now,
	}, nil
}

// Accept returns true if the line should be passed through, false if it
// should be dropped to stay within the configured rate.
// A nil Limiter always returns true.
func (l *Limiter) Accept(_ string) bool {
	if l == nil {
		return true
	}
	now := l.nowFn()
	if now.Sub(l.windowStart) >= time.Second {
		l.windowStart = now
		l.count = 0
	}
	if l.count < l.perSecond {
		l.count++
		return true
	}
	return false
}

// Rate returns the configured lines-per-second limit.
func (l *Limiter) Rate() int {
	if l == nil {
		return 0
	}
	return l.perSecond
}
