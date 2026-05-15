package config

import "flag"

// TruncateWidth returns the maximum line width configured via --truncate / -T.
// A value of 0 means no truncation. Negative values are treated as 0.
func TruncateWidth(fs *flag.FlagSet) int {
	if fs == nil {
		return 0
	}
	f := fs.Lookup("truncate")
	if f == nil {
		return 0
	}
	v, ok := f.Value.(interface{ Get() interface{} })
	if !ok {
		return 0
	}
	n, _ := v.Get().(int)
	if n < 0 {
		return 0
	}
	return n
}

// TruncateSuffix returns the suffix appended to truncated lines, configured
// via --truncate-suffix. Defaults to "...".
func TruncateSuffix(fs *flag.FlagSet) string {
	if fs == nil {
		return "..."
	}
	f := fs.Lookup("truncate-suffix")
	if f == nil {
		return "..."
	}
	return f.Value.String()
}

// registerTruncateFlags registers --truncate / -T and --truncate-suffix on fs.
func registerTruncateFlags(fs *flag.FlagSet) {
	fs.Int("truncate", 0, "truncate lines to `N` bytes (0 = disabled)")
	fs.Int("T", 0, "alias for --truncate")
	fs.String("truncate-suffix", "...", "suffix appended to truncated lines")
}
