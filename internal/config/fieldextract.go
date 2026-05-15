package config

import (
	"flag"
	"strconv"
	"strings"
)

// FieldExtractSeparator returns the field separator configured via --field-sep.
// Returns an empty string (no extraction) when fs is nil.
func FieldExtractSeparator(fs *flag.FlagSet) string {
	if fs == nil {
		return ""
	}
	f := fs.Lookup("field-sep")
	if f == nil {
		return ""
	}
	return f.Value.String()
}

// FieldExtractIndices returns the parsed 1-based field indices from --fields.
// Returns nil when fs is nil or the flag is not set / empty.
func FieldExtractIndices(fs *flag.FlagSet) []int {
	if fs == nil {
		return nil
	}
	f := fs.Lookup("fields")
	if f == nil {
		return nil
	}
	raw := f.Value.String()
	if raw == "" {
		return nil
	}
	var out []int
	for _, part := range strings.Split(raw, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		n, err := strconv.Atoi(part)
		if err == nil && n >= 1 {
			out = append(out, n)
		}
	}
	return out
}

// registerFieldExtractFlags registers --field-sep and --fields on fs.
func registerFieldExtractFlags(fs *flag.FlagSet) {
	fs.String("field-sep", " ", "field `separator` for --fields extraction")
	fs.String("fields", "", "comma-separated list of 1-based field `indices` to extract")
}
