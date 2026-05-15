// Package inverse provides a line filter that inverts the result of
// another filter, accepting lines that the wrapped filter would reject
// and rejecting lines that it would accept.
package inverse

// Accepter is the interface satisfied by any filter that can decide
// whether a line should be kept.
type Accepter interface {
	Accept(line string) bool
}

// Filter wraps an Accepter and inverts its decision.
type Filter struct {
	inner Accepter
}

// New returns a *Filter that inverts inner. If inner is nil, New returns
// nil; callers should treat a nil *Filter as a no-op (all lines accepted).
func New(inner Accepter) *Filter {
	if inner == nil {
		return nil
	}
	return &Filter{inner: inner}
}

// Accept returns true when the wrapped Accepter returns false, and vice
// versa. A nil receiver always returns true so that a nil *Filter can be
// used safely without a guard.
func (f *Filter) Accept(line string) bool {
	if f == nil {
		return true
	}
	return !f.inner.Accept(line)
}
