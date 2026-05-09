package config

// TailLines returns the number of trailing lines to emit, or 0 if the
// feature is disabled. A negative value is treated as 0.
func (c *Config) TailLines() int {
	if c == nil {
		return 0
	}
	if c.Tail < 0 {
		return 0
	}
	return c.Tail
}
