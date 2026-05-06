package reporter

import (
	"fmt"
	"io"
	"time"

	"github.com/yourorg/logslice/internal/stats"
)

// Reporter formats and writes processing statistics to an output writer.
type Reporter struct {
	w       io.Writer
	verbose bool
}

// New creates a new Reporter that writes to w.
// If verbose is true, additional detail is included in the report.
func New(w io.Writer, verbose bool) *Reporter {
	return &Reporter{w: w, verbose: verbose}
}

// Print writes a summary of s to the reporter's writer.
func (r *Reporter) Print(s *stats.Stats) error {
	snap := s.Snapshot()

	duration := snap.FinishedAt.Sub(snap.StartedAt).Round(time.Millisecond)
	if snap.FinishedAt.IsZero() {
		duration = time.Since(snap.StartedAt).Round(time.Millisecond)
	}

	_, err := fmt.Fprintf(r.w,
		"files: %d  read: %d  matched: %d  duration: %s\n",
		snap.Files,
		snap.LinesRead,
		snap.LinesMatched,
		duration,
	)
	if err != nil {
		return err
	}

	if r.verbose {
		_, err = fmt.Fprintf(r.w, "started: %s  finished: %s\n",
			snap.StartedAt.Format(time.RFC3339),
			snap.FinishedAt.Format(time.RFC3339),
		)
	}
	return err
}
