// Package highlight provides ANSI colour highlighting for matched
// substrings within log lines.
package highlight

import (
	"regexp"
	"strings"
)

const (
	colorReset  = "\033[0m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
	colorCyan   = "\033[36m"
)

// Color names accepted by New.
const (
	ColorYellow = "yellow"
	ColorRed    = "red"
	ColorCyan   = "cyan"
)

// Highlighter wraps matched substrings with ANSI escape codes.
type Highlighter struct {
	re    *regexp.Regexp
	open  string
	close string
}

// New returns a Highlighter that marks every match of pattern with the
// requested colour. An empty pattern returns nil without error so callers
// can treat a nil *Highlighter as a no-op. color defaults to yellow when
// an unrecognised name is supplied.
func New(pattern, color string) (*Highlighter, error) {
	if pattern == "" {
		return nil, nil
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	open := ansiOpen(color)
	return &Highlighter{re: re, open: open, close: colorReset}, nil
}

// Apply returns line with every match of the pattern wrapped in ANSI colour
// escapes. If h is nil the original line is returned unchanged.
func (h *Highlighter) Apply(line string) string {
	if h == nil {
		return line
	}
	var sb strings.Builder
	last := 0
	for _, loc := range h.re.FindAllStringIndex(line, -1) {
		sb.WriteString(line[last:loc[0]])
		sb.WriteString(h.open)
		sb.WriteString(line[loc[0]:loc[1]])
		sb.WriteString(h.close)
		last = loc[1]
	}
	sb.WriteString(line[last:])
	return sb.String()
}

func ansiOpen(color string) string {
	switch strings.ToLower(color) {
	case ColorRed:
		return colorRed
	case ColorCyan:
		return colorCyan
	default:
		return colorYellow
	}
}
