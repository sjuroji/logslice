package config

import "flag"

// Patterns returns the list of grep patterns from the parsed flag set.
// Returns nil if cfg is nil or no patterns were provided.
func Patterns(cfg *flag.FlagSet) []string {
	if cfg == nil {
		return nil
	}
	f := cfg.Lookup("pattern")
	if f == nil {
		return nil
	}
	raw := f.Value.String()
	if raw == "" {
		return nil
	}
	return splitComma(raw)
}

// ExcludePatterns returns the list of exclusion patterns from the parsed flag set.
// Returns nil if cfg is nil or no patterns were provided.
func ExcludePatterns(cfg *flag.FlagSet) []string {
	if cfg == nil {
		return nil
	}
	f := cfg.Lookup("exclude")
	if f == nil {
		return nil
	}
	raw := f.Value.String()
	if raw == "" {
		return nil
	}
	return splitComma(raw)
}

// registerGrepFlags registers --pattern and --exclude flags on the given flag set.
func registerGrepFlags(fs *flag.FlagSet) {
	fs.String("pattern", "", "comma-separated list of patterns to include (regex supported)")
	fs.String("p", "", "alias for --pattern")
	fs.String("exclude", "", "comma-separated list of patterns to exclude (regex supported)")
	fs.String("e", "", "alias for --exclude")
}
