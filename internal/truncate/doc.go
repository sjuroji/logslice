// Package truncate provides line-length limiting for log output.
//
// A Truncator shortens any line that exceeds a configured maximum width,
// appending a configurable suffix (e.g. "...") to indicate truncation.
//
// Usage:
//
//	t, err := truncate.New(80, "...")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	short := t.Apply("a very long line that exceeds the limit")
//
// A nil Truncator (returned when width <= 0) passes every line through
// unchanged, so callers need not special-case the disabled state.
package truncate
