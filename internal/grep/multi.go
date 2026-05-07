package grep

// Multi holds one or more Matchers and reports whether a line satisfies
// ALL of them (AND semantics). An empty Multi always matches.
type Multi struct {
	matchers []*Matcher
}

// NewMulti compiles each non-empty pattern in patterns and returns a
// Multi. Returns the first compilation error encountered.
func NewMulti(patterns []string) (*Multi, error) {
	ms := make([]*Matcher, 0, len(patterns))
	for _, p := range patterns {
		if p == "" {
			continue
		}
		m, err := New(p)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return &Multi{matchers: ms}, nil
}

// Match returns true when line satisfies every compiled pattern.
// If no patterns were provided it always returns true.
func (mu *Multi) Match(line string) bool {
	for _, m := range mu.matchers {
		if !m.Match(line) {
			return false
		}
	}
	return true
}

// Len returns the number of active patterns in the Multi.
func (mu *Multi) Len() int {
	return len(mu.matchers)
}
