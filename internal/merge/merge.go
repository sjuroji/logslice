// Package merge provides utilities for merging multiple sorted log streams
// into a single chronologically ordered output.
package merge

import (
	"bufio"
	"container/heap"
	"io"
	"time"

	"github.com/yourorg/logslice/internal/timeparser"
)

// entry holds a single log line together with its parsed timestamp and the
// index of the source reader it came from.
type entry struct {
	line      string
	ts        time.Time
	sourceIdx int
}

// minHeap implements heap.Interface over a slice of entries ordered by
// timestamp ascending.
type minHeap []entry

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].ts.Before(h[j].ts) }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(entry)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Merger merges multiple readers into a single ordered stream.
type Merger struct {
	scanners []*bufio.Scanner
	parser   *timeparser.Parser
}

// New returns a Merger that reads from each of the supplied readers and emits
// lines in ascending timestamp order.  A nil location defaults to UTC.
func New(readers []io.Reader, tp *timeparser.Parser) *Merger {
	scanners := make([]*bufio.Scanner, len(readers))
	for i, r := range readers {
		scanners[i] = bufio.NewScanner(r)
	}
	return &Merger{scanners: scanners, parser: tp}
}

// Merge reads all lines from every source reader, sorts them by the timestamp
// detected at the start of each line, and writes them in order to w.  Lines
// whose timestamps cannot be parsed are appended after all timestamped lines
// in their original source order.
func (m *Merger) Merge(w io.Writer) error {
	h := &minHeap{}
	heap.Init(h)

	var unparsed []string

	for idx, sc := range m.scanners {
		for sc.Scan() {
			line := sc.Text()
			ts, err := m.parser.Parse(line)
			if err != nil {
				unparsed = append(unparsed, line)
				continue
			}
			heap.Push(h, entry{line: line, ts: ts, sourceIdx: idx})
		}
		if err := sc.Err(); err != nil {
			return err
		}
	}

	bw := bufio.NewWriter(w)
	for h.Len() > 0 {
		e := heap.Pop(h).(entry)
		if _, err := bw.WriteString(e.line + "\n"); err != nil {
			return err
		}
	}
	for _, line := range unparsed {
		if _, err := bw.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return bw.Flush()
}
