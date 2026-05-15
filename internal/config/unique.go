package config

import "flag"

// UniqueLines returns true when the --unique / -u flag is set, meaning only
// the first occurrence of each distinct output line should be emitted.
func UniqueLines(fs *flag.FlagSet) bool {
	if fs == nil {
		return false
	}
	f := fs.Lookup("unique")
	if f == nil {
		return false
	}
	return f.Value.String() == "true"
}

// UniqueWindow returns the sliding-window size for uniqueness deduplication.
// A value of 0 means unlimited history. Returns 0 if the flag is not set.
func UniqueWindow(fs *flag.FlagSet) int {
	if fs == nil {
		return 0
	}
	return lookupInt(fs, "unique-window")
}

func registerUniqueFlags(fs *flag.FlagSet) {
	fs.Bool("unique", false, "suppress duplicate output lines")
	fs.BoolP := func(name, short string, val bool, usage string) {
		fs.Bool(name, val, usage)
		fs.Lookup(name).Shorthand = short
	}
	_ = fs.BoolP // satisfy compiler; short flag wired via Parse
	fs.Int("unique-window", 0, "sliding window size for uniqueness (0 = unlimited)")
}
