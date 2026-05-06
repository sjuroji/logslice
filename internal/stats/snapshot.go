package stats

import "time"

// Snapshot is an immutable copy of Stats values at a point in time.
type Snapshot struct {
	Files        int
	LinesRead    int
	LinesMatched int
	StartedAt    time.Time
	FinishedAt   time.Time
}

// Snapshot returns an immutable copy of the current Stats values.
// It is safe to call concurrently with other Stats methods.
func (s *Stats) Snapshot() Snapshot {
	s.mu.Lock()
	defer s.mu.Unlock()

	return Snapshot{
		Files:        s.files,
		LinesRead:    s.linesRead,
		LinesMatched: s.linesMatched,
		StartedAt:    s.startedAt,
		FinishedAt:   s.finishedAt,
	}
}
