package rotate_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yourorg/logslice/internal/rotate"
)

func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	p := filepath.Join(dir, "app.log")
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatalf("write temp file: %v", err)
	}
	return p
}

func TestNew_FileNotFound(t *testing.T) {
	_, err := rotate.New("/nonexistent/path/app.log")
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}

func TestCheck_UnchangedFile(t *testing.T) {
	p := writeTempFile(t, "line1\nline2\n")
	d, err := rotate.New(p)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if err := d.Check(); err != nil {
		t.Fatalf("expected nil for unchanged file, got %v", err)
	}
}

func TestCheck_ShrunkFile(t *testing.T) {
	p := writeTempFile(t, "line1\nline2\nline3\n")
	d, err := rotate.New(p)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	// Truncate the file to simulate rotation.
	if err := os.WriteFile(p, []byte("x\n"), 0o644); err != nil {
		t.Fatalf("truncate: %v", err)
	}
	if err := d.Check(); !rotate.IsRotated(err) {
		t.Fatalf("expected ErrRotated for shrunk file, got %v", err)
	}
}

func TestCheck_DeletedFile(t *testing.T) {
	p := writeTempFile(t, "data\n")
	d, err := rotate.New(p)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if err := os.Remove(p); err != nil {
		t.Fatalf("remove: %v", err)
	}
	if err := d.Check(); err == nil {
		t.Fatal("expected error for deleted file, got nil")
	}
}

func TestReset_UpdatesBaseline(t *testing.T) {
	p := writeTempFile(t, "initial\n")
	d, err := rotate.New(p)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	// Shrink the file.
	if err := os.WriteFile(p, []byte("x\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := d.Reset(); err != nil {
		t.Fatalf("Reset: %v", err)
	}
	// After reset the smaller file should not look rotated.
	if err := d.Check(); err != nil {
		t.Fatalf("expected nil after reset, got %v", err)
	}
}

func TestIsRotated_TrueForErrRotated(t *testing.T) {
	if !rotate.IsRotated(rotate.ErrRotated) {
		t.Fatal("IsRotated should return true for ErrRotated")
	}
}

func TestIsRotated_FalseForOtherError(t *testing.T) {
	if rotate.IsRotated(os.ErrNotExist) {
		t.Fatal("IsRotated should return false for os.ErrNotExist")
	}
}
