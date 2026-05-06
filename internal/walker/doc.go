// Package walker provides directory and file-path expansion for logslice.
//
// Given a mix of plain file paths and directory paths, Walker.Expand returns
// a sorted, deduplicated list of file paths suitable for sequential
// processing by the log pipeline.
//
// Optional suffix filtering (e.g. ".log", ".txt") allows callers to restrict
// expansion to relevant file types without requiring glob patterns from the
// shell.
//
// Recursive mode controls whether sub-directories are descended into;
// when disabled, only the immediate contents of a directory are considered.
package walker
