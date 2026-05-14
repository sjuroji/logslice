package config

import "flag"

// OutputFormat registers and reads output-format flags on the given FlagSet.
//
// Supported flags:
//
//	--format, -f   output format: plain (default), json, tsv
//	--color        enable ANSI colour highlighting of matched patterns
func OutputFormat(fs *flag.FlagSet) func() (format string, color bool) {
	if fs == nil {
		return func() (string, bool) { return "plain", false }
	}

	format := fs.String("format", "plain", "output format: plain, json, tsv")
	fs.String("f", "plain", "output format shorthand")

	color := fs.Bool("color", false, "enable ANSI colour highlighting of matched patterns")
	fs.Bool("colour", false, "alias for --color")

	return func() (string, bool) {
		f := *format
		// prefer short flag if explicitly set
		if sh := fs.Lookup("f"); sh != nil && sh.Value.String() != "plain" {
			f = sh.Value.String()
		}
		c := *color
		if alt := fs.Lookup("colour"); alt != nil && alt.Value.String() == "true" {
			c = true
		}
		return f, c
	}
}

// Format returns the output format string from a parsed Config.
// Returns "plain" when cfg is nil.
func Format(cfg *Config) string {
	if cfg == nil {
		return "plain"
	}
	return cfg.Format
}

// ColorEnabled returns whether ANSI colour output is enabled.
// Returns false when cfg is nil.
func ColorEnabled(cfg *Config) bool {
	if cfg == nil {
		return false
	}
	return cfg.Color
}
