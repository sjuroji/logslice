// Package stats provides a lightweight counter set for tracking log-processing
// metrics across one or more files.
//
// Usage:
//
//	s := stats.New()
//
//	// ... process files ...
//	s.AddFile()
//	s.AddRead(linesScanned)
//	s.AddMatched(linesEmitted)
//
//	s.Finish()
//	s.Write(os.Stderr)
//
// The Stats type is not safe for concurrent use; callers that process files
// in parallel should maintain per-goroutine instances and merge results after
// all goroutines complete.
package stats
