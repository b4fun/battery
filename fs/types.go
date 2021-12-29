package fs

import "os"

// WriteableFile provides the write interface for a file.
type WriteableFile interface {
	// Write writes len(b) bytes to the File.
	// It returns the number of bytes written and an error, if any.
	// Write returns a non-nil error when n != len(b).
	Write([]byte) (int, error)

	// Close closes the file.
	Close() error
}

// WritableFS provides a writeable FS interface for interacting with the file system.
// See: https://github.com/golang/go/issues/45757
type WritableFS interface {
	// OpenFile opens the named file with specified flag.
	// TODO: define the behavior
	OpenFile(name string, flag int, perm os.FileMode) (WriteableFile, error)

	// MkdirAll creates a directory named path, along with any necessary parents, and returns nil, or else returns an error.
	// TODO: define the behavior
	MkdirAll(path string, perm os.FileMode) error
}
