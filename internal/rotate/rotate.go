// Package rotate detects and handles rotated log files by tracking
// inode changes or size reductions between successive reads.
package rotate

import (
	"errors"
	"os"
)

// ErrRotated is returned when a log rotation is detected.
var ErrRotated = errors.New("log file has been rotated")

// Detector tracks file identity so that rotation can be detected
// between successive open/stat calls.
type Detector struct {
	path  string
	inode uint64
	size  int64
}

// New creates a Detector for the file at path. It stats the file
// immediately to record the baseline inode and size.
func New(path string) (*Detector, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &Detector{
		path:  path,
		inode: inode(info),
		size:  info.Size(),
	}, nil
}

// Check stats the file again and returns ErrRotated if the inode has
// changed or the file is smaller than it was at baseline. A nil error
// means the file appears to be the same log.
func (d *Detector) Check() error {
	info, err := os.Stat(d.path)
	if err != nil {
		return err
	}
	if inode(info) != d.inode {
		return ErrRotated
	}
	if info.Size() < d.size {
		return ErrRotated
	}
	return nil
}

// Reset updates the baseline to the current file state. Call this
// after successfully re-opening a rotated file.
func (d *Detector) Reset() error {
	info, err := os.Stat(d.path)
	if err != nil {
		return err
	}
	d.inode = inode(info)
	d.size = info.Size()
	return nil
}

// IsRotated reports whether err is ErrRotated.
func IsRotated(err error) bool {
	return errors.Is(err, ErrRotated)
}
