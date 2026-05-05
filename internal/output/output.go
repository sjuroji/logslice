// Package output handles writing filtered log lines to a destination.
package output

import (
	"bufio"
	"fmt"
	"io"
)

// Options configures the output writer.
type Options struct {
	// ShowLineNumbers prefixes each line with its original line number.
	ShowLineNumbers bool
	// Prefix is an optional string prepended to every output line.
	Prefix string
}

// Writer wraps an io.Writer and writes log lines with optional formatting.
type Writer struct {
	w    *bufio.Writer
	opts Options
}

// New creates a new Writer that writes to dst with the given options.
func New(dst io.Writer, opts Options) *Writer {
	return &Writer{
		w:    bufio.NewWriter(dst),
		opts: opts,
	}
}

// WriteLine writes a single log line to the destination.
// lineNum is the 1-based line number from the source file.
func (w *Writer) WriteLine(lineNum int, line string) error {
	var formatted string
	switch {
	case w.opts.ShowLineNumbers && w.opts.Prefix != "":
		formatted = fmt.Sprintf("%s%d: %s\n", w.opts.Prefix, lineNum, line)
	case w.opts.ShowLineNumbers:
		formatted = fmt.Sprintf("%d: %s\n", lineNum, line)
	case w.opts.Prefix != "":
		formatted = fmt.Sprintf("%s%s\n", w.opts.Prefix, line)
	default:
		formatted = line + "\n"
	}
	_, err := fmt.Fprint(w.w, formatted)
	return err
}

// Flush flushes any buffered data to the underlying writer.
func (w *Writer) Flush() error {
	return w.w.Flush()
}
