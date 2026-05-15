// Package fieldextract splits log lines by a configurable field separator
// and returns only the requested fields, joined back by the same separator.
//
// # Usage
//
//	e, err := fieldextract.New(":", []int{1, 3})
//	if err != nil {
//		log.Fatal(err)
//	}
//	output := e.Apply("2024-01-15:INFO:server started") // "2024-01-15:server started"
//
// Field indices are 1-based. An empty index slice causes Apply to return
// the original line unchanged, making the extractor a no-op suitable for
// conditional pipeline stages.
package fieldextract
