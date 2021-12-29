package fs

import (
	"os"
	"path/filepath"
)

type writableOSFS string

var _ WritableFS = writableOSFS("")

func (root writableOSFS) OpenFile(name string, flag int, perm os.FileMode) (WriteableFile, error) {
	return os.OpenFile(root.fullPath(name), flag, perm|0200) // 0200 is for granting write access to the owner
}

func (root writableOSFS) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(root.fullPath(path), perm|0200) // 0200 is for granting write access to the owner
}

func (root writableOSFS) fullPath(path string) string {
	if root == "" {
		return path
	}

	return filepath.Join(string(root), path)
}
