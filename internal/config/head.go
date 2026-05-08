package config

// HeadLines returns the number of leading lines to emit, or 0 when no head
// limit is configured. It is the --head / -H flag counterpart to --tail / -T.
//
// The value is stored on Config alongside MaxLines so that the pipeline can
// treat them uniformly: MaxLines caps total output, HeadLines caps from the
// start, and tail.Reader caps from the end.
func (c *Config) HeadLines() int {
	if c == nil {
		return 0
	}
	return c.Head
}
