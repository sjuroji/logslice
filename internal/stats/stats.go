// Package stats tracks processing statistics for a log slicing run.
package stats

import (
	"fmt"
	"io"
	"time"
)

// Stats holds counters accumulated during log processing.
type Stats struct {
	LinesRead    int
	LinesMatched int
	FilesRead    int
	StartedAt    time.Time
	FinishedAt   time.Time
}

// New returns a new Stats instance with the start time set to now.
func New() *Stats {
	return &Stats{StartedAt: time.Now()}
}

// Finish records the completion time.
func (s *Stats) Finish() {
	s.FinishedAt = time.Now()
}

// Elapsed returns the duration between start and finish.
// If Finish has not been called, it returns the duration since start.
func (s *Stats) Elapsed() time.Duration {
	if s.FinishedAt.IsZero() {
		return time.Since(s.StartedAt)
	}
	return s.FinishedAt.Sub(s.StartedAt)
}

// AddFile increments the file counter.
func (s *Stats) AddFile() {
	s.FilesRead++
}

// AddRead increments the lines-read counter by n.
func (s *Stats) AddRead(n int) {
	s.LinesRead += n
}

// AddMatched increments the lines-matched counter by n.
func (s *Stats) AddMatched(n int) {
	s.LinesMatched += n
}

// Write prints a human-readable summary to w.
func (s *Stats) Write(w io.Writer) {
	fmt.Fprintf(w, "files read   : %d\n", s.FilesRead)
	fmt.Fprintf(w, "lines read   : %d\n", s.LinesRead)
	fmt.Fprintf(w, "lines matched: %d\n", s.LinesMatched)
	fmt.Fprintf(w, "elapsed      : %s\n", s.Elapsed().Round(time.Millisecond))
}
