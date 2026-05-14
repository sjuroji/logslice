package config

import "flag"

// CountOnly returns true when the user requested only a line count (--count / -c)
// rather than printing matching lines.
func CountOnly(cfg *Config) bool {
	if cfg == nil {
		return false
	}
	return cfg.CountOnly
}

// registerCountFlag wires the --count / -c flag into the provided FlagSet and
// stores the result in cfg.
func registerCountFlag(fs *flag.FlagSet, cfg *Config) {
	fs.BoolVar(&cfg.CountOnly, "count", false, "print only a count of matching lines per file")
	fs.BoolVar(&cfg.CountOnly, "c", false, "print only a count of matching lines per file (shorthand)")
}
