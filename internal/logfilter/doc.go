// Package logfilter provides a unified Filter type that combines
// time-range filtering and substring pattern matching for log lines.
//
// A Filter is constructed from a Config that may specify:
//   - Start / End: an inclusive time window; lines whose parsed timestamp
//     falls outside this window are rejected.
//   - Pattern: a plain-text substring that must appear in the line.
//
// Either constraint is optional. When neither is set every line is accepted.
//
// Example:
//
//	f, err := logfilter.New(logfilter.Config{
//		Start:   start,
//		End:     end,
//		Pattern: "ERROR",
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, line := range lines {
//		if f.Match(line) {
//			fmt.Println(line)
//		}
//	}
package logfilter
