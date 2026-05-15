// Package fieldextract provides utilities for extracting and filtering
// fields from a log line using a configurable separator.
package fieldextract

import (
	"fmt"
	"strings"
)

// Extractor splits lines by a separator and returns selected fields.
type Extractor struct {
	sep     string
	indices []int // 1-based field indices
}

// New creates an Extractor that splits on sep and returns the fields at
// the given 1-based indices. An empty indices slice means return all fields
// (joined by sep). Returns an error if any index is less than 1.
func New(sep string, indices []int) (*Extractor, error) {
	if sep == "" {
		sep = " "
	}
	for _, idx := range indices {
		if idx < 1 {
			return nil, fmt.Errorf("fieldextract: index must be >= 1, got %d", idx)
		}
	}
	copied := make([]int, len(indices))
	copy(copied, indices)
	return &Extractor{sep: sep, indices: copied}, nil
}

// Apply extracts the requested fields from line and returns them joined by
// the separator. If indices is empty the original line is returned unchanged.
// Fields that are out of range for a given line are silently skipped.
func (e *Extractor) Apply(line string) string {
	if e == nil || len(e.indices) == 0 {
		return line
	}
	parts := strings.Split(line, e.sep)
	var out []string
	for _, idx := range e.indices {
		if idx <= len(parts) {
			out = append(out, parts[idx-1])
		}
	}
	if len(out) == 0 {
		return ""
	}
	return strings.Join(out, e.sep)
}

// Fields returns the number of fields in line when split by the extractor's
// separator.
func (e *Extractor) Fields(line string) int {
	if e == nil {
		return 0
	}
	return len(strings.Split(line, e.sep))
}
