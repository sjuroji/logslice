package config

import "flag"

// MergeInputs returns true when the --merge flag is set, indicating that
// multiple input files should be merged and sorted by timestamp before output.
func MergeInputs(fs *flag.FlagSet) bool {
	if fs == nil {
		return false
	}
	f := fs.Lookup("merge")
	if f == nil {
		return false
	}
	return f.Value.String() == "true"
}

// MergeTimestampLayout returns the timestamp layout used when parsing lines
// during a merge operation. Falls back to an empty string if not set.
func MergeTimestampLayout(fs *flag.FlagSet) string {
	if fs == nil {
		return ""
	}
	f := fs.Lookup("merge-layout")
	if f == nil {
		return ""
	}
	return f.Value.String()
}

// registerMergeFlags registers --merge and --merge-layout onto fs.
func registerMergeFlags(fs *flag.FlagSet) {
	fs.Bool("merge", false, "merge multiple input files sorted by timestamp")
	fs.String("merge-layout", "", "timestamp `layout` used when merging (default: auto-detect)")
}
