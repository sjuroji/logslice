package config

import (
	"flag"
	"time"
)

// Timeout returns the optional execution timeout parsed from the flag set.
// Returns zero if no timeout flag was registered or the value is zero.
func Timeout(fs *flag.FlagSet) time.Duration {
	if fs == nil {
		return 0
	}
	f := fs.Lookup("timeout")
	if f == nil {
		return 0
	}
	gv, ok := f.Value.(flag.Getter)
	if !ok {
		return 0
	}
	d, _ := gv.Get().(time.Duration)
	return d
}

// registerContextFlags adds context-related flags to the provided flag set.
func registerContextFlags(fs *flag.FlagSet) {
	fs.Duration("timeout", 0, "maximum execution time (e.g. 30s, 2m); 0 means no limit")
}
