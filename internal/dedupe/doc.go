// Package dedupe provides a concurrency-safe duplicate-line filter
// for use in log processing pipelines.
//
// A Filter tracks every distinct line it has seen. On the first
// occurrence of a line, Accept returns true so the caller may emit
// it. On every subsequent occurrence, Accept returns false and
// increments an internal suppression counter.
//
// Usage:
//
//	f := dedupe.New()
//	for _, line := range lines {
//		if f.Accept(line) {
//			fmt.Println(line)
//		}
//	}
//
The filter is safe for concurrent use from multiple goroutines.
package dedupe
