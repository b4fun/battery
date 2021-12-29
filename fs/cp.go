package fs

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

// CopyAllToDir copies all entries from src filesystem to destintation dir.
func CopyAllToDir(src fs.FS, dest string) error {
	return copyToWritableFS(src, writableOSFS(dest))
}

func copyToWritableFS(src fs.FS, dest WritableFS) error {
	return fs.WalkDir(src, ".", func(path string, srcEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, err := srcEntry.Info()
		if err != nil {
			return fmt.Errorf("stat src: %w", err)
		}

		if srcEntry.IsDir() {
			return dest.MkdirAll(path, info.Mode())
		}

		srcFile, err := src.Open(path)
		if err != nil {
			return fmt.Errorf("open src: %w", err)
		}

		destFile, err := dest.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, info.Mode())
		if err != nil {
			return fmt.Errorf("open dest: %w", err)
		}

		if _, err := io.Copy(destFile, srcFile); err != nil {
			return fmt.Errorf("copy: %w", err)
		}

		if err := srcFile.Close(); err != nil {
			return fmt.Errorf("close src: %w", err)
		}
		if err := destFile.Close(); err != nil {
			return fmt.Errorf("close dest: %w", err)
		}

		return nil
	})
}
