// Package processor ties together log filtering and output writing
// to process log lines from an io.Reader source.
package processor

import (
	"bufio"
	"io"

	"github.com/yourorg/logslice/internal/logfilter"
	"github.com/yourorg/logslice/internal/output"
)

// Options holds configuration for a processing run.
type Options struct {
	// MaxLines limits the number of matching lines written. Zero means unlimited.
	MaxLines int
}

// Processor reads lines from a source, applies a filter, and writes
// matching lines to a writer.
type Processor struct {
	filter  *logfilter.Filter
	writer  *output.Writer
	options Options
}

// New creates a Processor with the given filter, writer, and options.
func New(f *logfilter.Filter, w *output.Writer, opts Options) *Processor {
	return &Processor{
		filter:  f,
		writer:  w,
		options: opts,
	}
}

// Process reads all lines from r, filters them, and writes matches.
// It returns the number of lines written and any write error encountered.
func (p *Processor) Process(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	written := 0

	for scanner.Scan() {
		line := scanner.Text()

		if !p.filter.Match(line) {
			continue
		}

		if err := p.writer.WriteLine(line); err != nil {
			return written, err
		}

		written++

		if p.options.MaxLines > 0 && written >= p.options.MaxLines {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return written, err
	}

	return written, nil
}
