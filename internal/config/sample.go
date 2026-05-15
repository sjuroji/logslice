package config

import "flag"

// SampleN returns the --sample / -S flag value from fs.
// When n == 1 (the default) every line is emitted, which is equivalent
// to no sampling.
func SampleN(fs *flag.FlagSet) int {
	if fs == nil {
		return 1
	}
	f := fs.Lookup("sample")
	if f == nil {
		return 1
	}
	g, ok := f.Value.(flag.Getter)
	if !ok {
		return 1
	}
	v, ok := g.Get().(int)
	if !ok {
		return 1
	}
	if v < 1 {
		return 1
	}
	return v
}

// registerSampleFlag adds --sample and -S to fs.
func registerSampleFlag(fs *flag.FlagSet) {
	fs.Int("sample", 1, "emit every Nth matching line (1 = all lines)")
	fs.Int("S", 1, "shorthand for --sample")
}
