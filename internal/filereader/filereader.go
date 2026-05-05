package filereader

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Reader wraps file access and provides line-by-line reading.
type Reader struct {
	path string
	file *os.File
	scanner *bufio.Scanner
}

// New opens the file at path and returns a Reader ready for scanning.
// The caller must call Close when done.
func New(path string) (*Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("filereader: open %q: %w", path, err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	return &Reader{
		path:    path,
		file:    f,
		scanner: scanner,
	}, nil
}

// NewFromReader creates a Reader backed by an existing io.Reader.
// Close on this Reader is a no-op.
func NewFromReader(r io.Reader) *Reader {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	return &Reader{
		scanner: scanner,
	}
}

// Scan advances to the next line. Returns true if a line is available.
func (r *Reader) Scan() bool {
	return r.scanner.Scan()
}

// Text returns the current line text.
func (r *Reader) Text() string {
	return r.scanner.Text()
}

// Err returns any non-EOF error encountered during scanning.
func (r *Reader) Err() error {
	return r.scanner.Err()
}

// Path returns the file path associated with this Reader.
func (r *Reader) Path() string {
	return r.path
}

// Close releases the underlying file resource.
func (r *Reader) Close() error {
	if r.file != nil {
		return r.file.Close()
	}
	return nil
}
