package walker

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Walker traverses a list of paths, expanding directories into individual
// files that match an optional suffix filter.
type Walker struct {
	suffixes []string
	recursive bool
}

// New returns a Walker. If suffixes is non-empty, only files whose names end
// with one of the provided suffixes are returned (e.g. ".log").
func New(recursive bool, suffixes ...string) *Walker {
	return &Walker{suffixes: suffixes, recursive: recursive}
}

// Expand takes a slice of paths (files or directories) and returns an ordered
// list of matching file paths. Directories are walked; plain files are
// included as-is (suffix filter still applies).
func (w *Walker) Expand(paths []string) ([]string, error) {
	var results []string

	for _, p := range paths {
		info, err := os.Stat(p)
		if err != nil {
			return nil, err
		}

		if !info.IsDir() {
			if w.matches(info.Name()) {
				results = append(results, p)
			}
			continue
		}

		if err := w.walkDir(p, &results); err != nil {
			return nil, err
		}
	}

	sort.Strings(results)
	return results, nil
}

func (w *Walker) walkDir(root string, results *[]string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if path != root && !w.recursive {
				return filepath.SkipDir
			}
			return nil
		}
		if w.matches(d.Name()) {
			*results = append(*results, path)
		}
		return nil
	})
}

func (w *Walker) matches(name string) bool {
	if len(w.suffixes) == 0 {
		return true
	}
	for _, s := range w.suffixes {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}
