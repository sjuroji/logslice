// Package merge implements a k-way merge of multiple log streams into a single
// chronologically ordered output.
//
// # Overview
//
// When log data is spread across several files (e.g. from multiple replicas or
// rotated archives) it is often necessary to interleave the lines so that
// events can be read in the order they actually occurred.
//
// Merger reads all source readers into an in-memory min-heap keyed on the
// timestamp detected at the start of each line.  Lines that carry no
// recognisable timestamp are collected separately and appended verbatim after
// all timestamped output so that no data is silently discarded.
//
// # Usage
//
//	p, _ := timeparser.New(time.UTC)
//	m := merge.New([]io.Reader{fileA, fileB}, p)
//	if err := m.Merge(os.Stdout); err != nil {
//		log.Fatal(err)
//	}
package merge
