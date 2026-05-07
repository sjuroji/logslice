// Package progress provides a simple progress reporter that writes
// per-file processing status to an io.Writer.
package progress

import (
	"fmt"
	"io"
	"sync"
)

// Reporter writes progress messages as files are processed.
type Reporter struct {
	w       io.Writer
	mu      sync.Mutex
	verbose bool
	total   int
	current int
}

// New creates a new Reporter that writes to w.
// When verbose is true, additional detail (matched/read counts) is printed.
func New(w io.Writer, total int, verbose bool) *Reporter {
	return &Reporter{
		w:       w,
		total:   total,
		verbose: verbose,
	}
}

// FileStart prints a message indicating that processing of path has begun.
func (r *Reporter) FileStart(path string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.current++
	if r.verbose {
		fmt.Fprintf(r.w, "[%d/%d] processing %s\n", r.current, r.total, path)
	}
}

// FileDone prints a summary line for path after processing is complete.
// read is the number of lines read; matched is the number that passed filters.
func (r *Reporter) FileDone(path string, read, matched int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.verbose {
		fmt.Fprintf(r.w, "[%d/%d] done    %s — read %d, matched %d\n",
			r.current, r.total, path, read, matched)
	}
}

// Error prints an error message for path.
func (r *Reporter) Error(path string, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	fmt.Fprintf(r.w, "error: %s: %v\n", path, err)
}
