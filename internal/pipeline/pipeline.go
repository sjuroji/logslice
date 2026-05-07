// Package pipeline wires together the core logslice processing components
// into a single reusable execution unit.
package pipeline

import (
	"io"

	"github.com/example/logslice/internal/config"
	"github.com/example/logslice/internal/filereader"
	"github.com/example/logslice/internal/logfilter"
	"github.com/example/logslice/internal/output"
	"github.com/example/logslice/internal/processor"
	"github.com/example/logslice/internal/stats"
)

// Pipeline holds the assembled components needed to process a single file.
type Pipeline struct {
	cfg      *config.Config
	out      *output.Writer
	stats    *stats.Stats
}

// New creates a Pipeline from the given configuration, writing results to w.
func New(cfg *config.Config, w io.Writer, s *stats.Stats) *Pipeline {
	return &Pipeline{
		cfg:   cfg,
		out:   output.New(w, cfg.LineNumbers, cfg.Prefix),
		stats: s,
	}
}

// Run opens path, applies filters defined in cfg, and writes matching lines.
// Statistics are updated on the supplied Stats instance.
func (p *Pipeline) Run(path string) error {
	reader, err := filereader.New(path)
	if err != nil {
		return err
	}
	defer reader.Close()

	p.stats.AddFile()

	filter, err := logfilter.New(
		cfg(p).Pattern,
		cfg(p).TimeField,
		cfg(p).TimeLayout,
		cfg(p).Start,
		cfg(p).End,
	)
	if err != nil {
		return err
	}

	proc := processor.New(reader, filter, p.out, p.cfg.MaxLines)
	read, matched, err := proc.Process()
	if err != nil {
		return err
	}

	p.stats.AddRead(read)
	p.stats.AddMatched(matched)
	return nil
}

// cfg is a small helper to avoid repeating the field access in Run.
func cfg(p *Pipeline) *config.Config { return p.cfg }
