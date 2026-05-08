// Package head provides a line-count limiter for the leading lines of a log
// stream — the counterpart to the tail package which reads trailing lines.
//
// # Overview
//
// A [Reader] wraps a positive integer n and exposes an [Reader.Accept] method
// that returns true for the first n calls and false thereafter. This makes it
// easy to short-circuit processing loops without threading a counter through
// every layer of the pipeline.
//
// A nil *Reader is intentionally valid and means "no limit": both Accept and
// Done return their zero-value results (true and false respectively), so
// callers never need to guard against a nil check in hot paths.
//
// # Usage
//
//	r := head.New(cfg.HeadLines) // nil when HeadLines <= 0
//	for scanner.Scan() {
//	    if !r.Accept() {
//	        break
//	    }
//	    process(scanner.Text())
//	}
package head
