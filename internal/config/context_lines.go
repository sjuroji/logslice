package config

import "flag"

// ContextLines registers and reads the -before / -after / -context flags
// that control how many surrounding lines are printed around each match,
// mirroring the behaviour of grep -B / -A / -C.

func registerContextLinesFlags(fs *flag.FlagSet) {
	fs.Int("before", 0, "print `N` lines of context before each matching line (like grep -B)")
	fs.Int("after", 0, "print `N` lines of context after each matching line (like grep -A)")
	fs.Int("context", 0, "print `N` lines of context before and after each matching line (like grep -C)")
}

// BeforeLines returns the number of leading context lines requested.
// A non-zero -context value overrides -before.
func BeforeLines(fs *flag.FlagSet) int {
	if fs == nil {
		return 0
	}
	if c := lookupInt(fs, "context"); c > 0 {
		return c
	}
	return lookupInt(fs, "before")
}

// AfterLines returns the number of trailing context lines requested.
// A non-zero -context value overrides -after.
func AfterLines(fs *flag.FlagSet) int {
	if fs == nil {
		return 0
	}
	if c := lookupInt(fs, "context"); c > 0 {
		return c
	}
	return lookupInt(fs, "after")
}

// lookupInt retrieves an integer flag value from fs, returning 0 if the flag
// does not exist or has not been set.
func lookupInt(fs *flag.FlagSet, name string) int {
	f := fs.Lookup(name)
	if f == nil {
		return 0
	}
	v, ok := f.Value.(interface{ Get() interface{} })
	if !ok {
		return 0
	}
	if n, ok := v.Get().(int); ok {
		return n
	}
	return 0
}
