// Package count provides lightweight line-count primitives used throughout
// the logslice pipeline.
//
// Counter is a thread-safe atomic counter suitable for tracking matched or
// read lines inside a concurrent pipeline stage.
//
// Accumulator collects per-file line counts and maintains a running total,
// making it easy to produce summary statistics after processing a set of
// log files.
//
// Example usage:
//
//	c := count.New()
//	for _, line := range lines {
//		if filter.Match(line) {
//			c.Inc()
//		}
//	}
//	fmt.Println("matched:", c.Value())
package count
