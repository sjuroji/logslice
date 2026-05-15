// Package sample implements a deterministic line sampler for log processing.
//
// A Sampler accepts every Nth line from a stream, discarding the rest.
// This is useful when log files are extremely large and only a statistical
// sample of the output is required rather than every matching line.
//
// Usage:
//
//	s, err := sample.New(10) // emit every 10th line
//	if err != nil {
//		log.Fatal(err)
//	}
//	if s.Accept(line) {
//		// process line
//	}
//
// A nil *Sampler is safe to use and always returns true from Accept,
// making it easy to disable sampling without conditional logic.
package sample
