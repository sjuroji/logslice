// Package pipeline wires together all internal components and drives
// the per-file processing loop.
package pipeline

import (
	"fmt"
	"io"

	"github.com/yourorg/logslice/internal/config"
	"github.com/yourorg/logslice/internal/filereader"
	"github.com/yourorg/logslice/internal/head"
	"github.com/yourorg/logslice/internal/logfilter"
	"github.com/yourorg/logslice/internal/output"
	"github.com/yourorg/logslice/internal/processor"
	"github.com/yourorg/logslice/internal/progress"
	"github.com/yourorg/logslice/internal/stats"
	"github.com/yourorg/logslice/internal/tail"
)

// Pipeline holds the shared components used across every file processed in
// a single logslice invocation.
type Pipeline struct {
	cfg      *config.Config
	out      *output.Writer
	prog     *progress.Reporter
	stats    *stats.Stats
}

// New constructs a Pipeline from the supplied configuration, writing
// matching lines to w and progress/diagnostic messages to errW.
func New(cfg *config.Config, w io.Writer, errW io.Writer, s *stats.Stats) *Pipeline {
	out := output.New(w, cfg.LineNumbers, cfg.Prefix)
	prog := progress.New(errW, cfg.Verbose)
	return &Pipeline{
		cfg:   cfg,
		out:   out,
		prog:  prog,
		stats: s,
	}
}

// Run processes a single file identified by path, applying all filters
// defined in the pipeline configuration. It updates the shared Stats and
// reports progress via the progress reporter.
func (p *Pipeline) Run(path string) error {
	fr, err := filereader.New(path)
	if err != nil {
		p.prog.Error(path, err)
		return fmt.Errorf("open %s: %w", path, err)
	}
	defer fr.Close()

	p.stats.AddFile()
	p.prog.FileStart(path)

	// Build the log-line filter from config.
	f, err := cfg(p.cfg)
	if err != nil {
		return fmt.Errorf("build filter: %w", err)
	}

	proc := processor.New(fr, f, p.out, p.stats)

	// Wrap with head / tail limiters when requested.
	if n := p.cfg.HeadLines(); n > 0 {
		h := head.New(n)
		read, matched, err := proc.ProcessHead(h)
		p.prog.FileDone(path, read, matched)
		return err
	}

	if n := p.cfg.TailLines(); n > 0 {
		t := tail.New(n)
		read, matched, err := proc.ProcessTail(t, p.out)
		p.prog.FileDone(path, read, matched)
		return err
	}

	read, matched, err := proc.Process()
	p.prog.FileDone(path, read, matched)
	return err
}

// cfg builds a logfilter.Filter from the pipeline configuration.
func cfg(c *config.Config) (*logfilter.Filter, error) {
	opts := []logfilter.Option{}

	if c.TimeRange != nil {
		opts = append(opts, logfilter.WithTimeRange(c.TimeRange))
	}

	if len(c.Patterns) > 0 {
		opts = append(opts, logfilter.WithPatterns(c.Patterns))
	}

	if len(c.ExcludePatterns) > 0 {
		opts = append(opts, logfilter.WithExcludePatterns(c.ExcludePatterns))
	}

	return logfilter.New(opts...)
}
