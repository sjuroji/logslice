package walker_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/example/logslice/internal/walker"
)

// buildTree creates a temporary directory tree for testing.
//
//	root/
//	  a.log
//	  b.txt
//	  sub/
//	    c.log
func buildTree(t *testing.T) string {
	t.Helper()
	root := t.TempDir()

	writeFile(t, filepath.Join(root, "a.log"), "line1")
	writeFile(t, filepath.Join(root, "b.txt"), "line2")

	sub := filepath.Join(root, "sub")
	if err := os.Mkdir(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(sub, "c.log"), "line3")
	return root
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestExpand_SingleFile(t *testing.T) {
	root := buildTree(t)
	w := walker.New(false)
	got, err := w.Expand([]string{filepath.Join(root, "a.log")})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 {
		t.Fatalf("want 1 file, got %d", len(got))
	}
}

func TestExpand_DirectoryNonRecursive(t *testing.T) {
	root := buildTree(t)
	w := walker.New(false, ".log")
	got, err := w.Expand([]string{root})
	if err != nil {
		t.Fatal(err)
	}
	// Only a.log at root level; sub/ is skipped.
	if len(got) != 1 {
		t.Fatalf("want 1 file, got %d: %v", len(got), got)
	}
}

func TestExpand_DirectoryRecursive(t *testing.T) {
	root := buildTree(t)
	w := walker.New(true, ".log")
	got, err := w.Expand([]string{root})
	if err != nil {
		t.Fatal(err)
	}
	// a.log + sub/c.log
	if len(got) != 2 {
		t.Fatalf("want 2 files, got %d: %v", len(got), got)
	}
}

func TestExpand_SuffixFilter(t *testing.T) {
	root := buildTree(t)
	w := walker.New(false, ".txt")
	got, err := w.Expand([]string{root})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 {
		t.Fatalf("want 1 file, got %d: %v", len(got), got)
	}
}

func TestExpand_FileNotFound(t *testing.T) {
	w := walker.New(false)
	_, err := w.Expand([]string{"/does/not/exist.log"})
	if err == nil {
		t.Fatal("expected error for missing path")
	}
}

func TestExpand_ResultsAreSorted(t *testing.T) {
	root := buildTree(t)
	w := walker.New(true)
	got, err := w.Expand([]string{root})
	if err != nil {
		t.Fatal(err)
	}
	for i := 1; i < len(got); i++ {
		if got[i] < got[i-1] {
			t.Fatalf("results not sorted: %v", got)
		}
	}
}
