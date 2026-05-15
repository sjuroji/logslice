package config

import (
	"flag"
	"strings"
)

// FieldSeparator returns the field separator string used when splitting log
// lines into columns. Defaults to a single space. Registered under --field-sep
// and -F.
func FieldSeparator(fs *flag.FlagSet) string {
	if fs == nil {
		return ""
	}
	f := fs.Lookup("field-sep")
	if f == nil {
		return ""
	}
	return f.Value.String()
}

// FieldIndices returns the list of 1-based column indices to extract when
// --fields is supplied. An empty slice means all fields are emitted.
func FieldIndices(fs *flag.FlagSet) []string {
	if fs == nil {
		return nil
	}
	f := fs.Lookup("fields")
	if f == nil {
		return nil
	}
	v := f.Value.String()
	if v == "" {
		return nil
	}
	return splitComma(v)
}

func registerFieldFlags(fs *flag.FlagSet) {
	fs.String("field-sep", " ", "field separator for column extraction (default space)")
	fs.String("F", " ", "shorthand for --field-sep")
	fs.String("fields", "", "comma-separated list of 1-based field indices to extract")
}

// ApplyFieldFilter extracts the requested fields from line using sep as the
// delimiter. If indices is empty the original line is returned unchanged.
func ApplyFieldFilter(line, sep string, indices []string) string {
	if len(indices) == 0 || sep == "" {
		return line
	}
	parts := strings.Split(line, sep)
	out := make([]string, 0, len(indices))
	for _, idx := range indices {
		n := 0
		for _, ch := range idx {
			if ch >= '0' && ch <= '9' {
				n = n*10 + int(ch-'0')
			}
		}
		if n >= 1 && n <= len(parts) {
			out = append(out, parts[n-1])
		}
	}
	return strings.Join(out, sep)
}
