package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func writeTempLog(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "test-*.log")
	if err != nil {
		t.Fatalf("create temp log: %v", err)
	}
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("write temp log: %v", err)
	}
	f.Close()
	return f.Name()
}

func buildBinary(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	bin := filepath.Join(dir, "logslice")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	cmd.Dir = "."
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build failed: %v\n%s", err, out)
	}
	return bin
}

func TestMain_NoArgs(t *testing.T) {
	bin := buildBinary(t)
	cmd := exec.Command(bin)
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("expected non-zero exit with no args")
	}
	if !strings.Contains(string(out), "usage:") {
		t.Errorf("expected usage message, got: %s", out)
	}
}

func TestMain_PatternFilter(t *testing.T) {
	bin := buildBinary(t)
	log := writeTempLog(t, "INFO start\nDEBUG skip\nINFO end\n")
	cmd := exec.Command(bin, "-pattern", "INFO", log)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) != 2 {
		t.Errorf("expected 2 matching lines, got %d: %v", len(lines), lines)
	}
}

func TestMain_MaxLines(t *testing.T) {
	bin := buildBinary(t)
	log := writeTempLog(t, "a\nb\nc\nd\ne\n")
	cmd := exec.Command(bin, "-n", "3", log)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) != 3 {
		t.Errorf("expected 3 lines, got %d", len(lines))
	}
}
