// Package reporter formats and prints processing statistics produced by
// logslice after one or more log files have been processed.
//
// Usage:
//
//	s := stats.New()
//	// ... process files, updating s ...
//	s.Finish()
//
//	r := reporter.New(os.Stderr, verbose)
//	if err := r.Print(s); err != nil {
//		log.Fatal(err)
//	}
//
// When verbose is false only the file count, lines read, lines matched and
// elapsed duration are printed.  When verbose is true the absolute start and
// finish timestamps are also included.
package reporter
