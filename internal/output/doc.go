// Package output provides a buffered writer for formatted log output.
//
// It wraps an io.Writer and supports optional features such as
// line number annotations and custom line prefixes, making it
// suitable for presenting filtered log results to stdout or a file.
//
// Example usage:
//
//	w := output.New(os.Stdout, output.Options{
//		ShowLineNumbers: true,
//		Prefix:          "[match] ",
//	})
//	w.WriteLine(12, "2024-01-01 ERROR connection refused")
//	w.Flush()
package output
