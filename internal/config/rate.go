package config

import "flag"

// RateLimit returns the maximum number of lines per second to emit, or 0 if
// no rate limiting has been requested.
func RateLimit(fs *flag.FlagSet) int {
	if fs == nil {
		return 0
	}
	f := fs.Lookup("rate")
	if f == nil {
		return 0
	}
	g, ok := f.Value.(flag.Getter)
	if !ok {
		return 0
	}
	v, _ := g.Get().(int)
	return v
}

// registerRateFlag adds the --rate / -r flag to fs.
func registerRateFlag(fs *flag.FlagSet) {
	fs.Int("rate", 0, "maximum lines per second to emit (0 = unlimited)")
	fs.Int("r", 0, "maximum lines per second to emit (shorthand)")
}
