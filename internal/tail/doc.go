// Package tail implements a fixed-size ring-buffer reader that retains
// only the last N lines consumed from an io.Reader.
//
// It is intended for use with the --tail flag in logslice, allowing users
// to inspect the final N lines of a log file — after all other filters
// (pattern, time-range) have been applied by the pipeline.
//
// Usage:
//
//	tr := tail.New(100)
//	if err := tr.ReadAll(reader); err != nil {
//	    return err
//	}
//	for _, line := range tr.Lines() {
//	    fmt.Println(line)
//	}
//
// A nil Tailer is returned when n <= 0, signalling that tailing is
// disabled and the caller should pass lines through unchanged.
package tail
