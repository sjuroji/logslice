package config

import "flag"

// DedupeLines returns true when the --dedupe flag is set, causing the
// pipeline to suppress duplicate output lines.
func DedupeLines(fs *flag.FlagSet) bool {
	if fs == nil {
		return false
	}
	f := fs.Lookup("dedupe")
	if f == nil {
		return false
	}
	return f.Value.String() == "true"
}

// registerDedupeFlag adds the --dedupe / -D flag to fs.
func registerDedupeFlag(fs *flag.FlagSet) {
	fs.Bool("dedupe", false, "suppress duplicate output lines")
	fs.Bool("D", false, "suppress duplicate output lines (shorthand)")
}
